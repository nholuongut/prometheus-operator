// Copyright 2018 The Nho Luong DevOps
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sYAML "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/yaml"

	monitoring "github.com/nholuongut/prometheus-operator/pkg/apis/monitoring"
	monitoringv1 "github.com/nholuongut/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/nholuongut/prometheus-operator/pkg/versionutil"
)

func main() {
	fs := flag.CommandLine
	versionutil.RegisterFlags(fs)

	var ruleConfigMapName = flag.String("rule-config-map", "", "path to rule ConfigMap")
	var ruleCRDSDestination = flag.String("rule-crds-destination", "", "destination new crds should be created in")

	// No need to check for errors because Parse would exit on error.
	_ = fs.Parse(os.Args[1:])

	if versionutil.ShouldPrintVersion() {
		versionutil.Print(os.Stdout, "po-rule-migration")
		os.Exit(0)
	}

	if *ruleConfigMapName == "" {
		log.Print("please specify 'rule-config-map' flag")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *ruleCRDSDestination == "" {
		log.Print("please specify 'rule-crds-destination' flag")
		flag.PrintDefaults()
		os.Exit(1)
	}

	destPath, err := filepath.Abs(*ruleCRDSDestination)
	if err != nil {
		log.Fatalf("failed to get absolute path of '%v': %v", ruleCRDSDestination, err.Error())
	}
	ruleCRDSDestination = &destPath

	file, err := os.Open(*ruleConfigMapName)
	if err != nil {
		log.Fatalf("failed to read file '%v': %v", ruleConfigMapName, err.Error())
	}

	configMap := v1.ConfigMap{}

	err = k8sYAML.NewYAMLOrJSONDecoder(file, 100).Decode(&configMap)
	if err != nil {
		log.Fatalf("failed to decode manifest: %v", err.Error())
	}

	ruleFiles, err := CMToRule(&configMap)
	if err != nil {
		log.Fatalf("failed to transform ConfigMap to rule file crds: %v", err.Error())
	}

	for _, ruleFile := range ruleFiles {
		encodedRuleFile, err := yaml.Marshal(ruleFile)
		if err != nil {
			log.Fatalf("failed to encode ruleFile '%v': %v", ruleFile.Name, err.Error())
		}

		err = os.WriteFile(path.Join(*ruleCRDSDestination, ruleFile.Name), encodedRuleFile, 0644)
		if err != nil {
			log.Fatalf("failed to write yaml manifest for rule file '%v': %v", ruleFile.Name, err.Error())
		}
	}
}

// CMToRule takes a rule ConfigMap and transforms it to possibly multiple
// rule file crds. It is used in `cmd/po-rule-cm-to-rule-file-crds`. Thereby it
// needs to be public.
func CMToRule(cm *v1.ConfigMap) ([]monitoringv1.PrometheusRule, error) {
	rules := []monitoringv1.PrometheusRule{}

	for name, content := range cm.Data {
		ruleSpec := monitoringv1.PrometheusRuleSpec{}

		if err := k8sYAML.NewYAMLOrJSONDecoder(bytes.NewBufferString(content), 1000).Decode(&ruleSpec); err != nil {
			return []monitoringv1.PrometheusRule{}, fmt.Errorf(
				"unmarshal rules file %v in  configmap '%v' in namespace '%v': %w",
				name, cm.Name, cm.Namespace, err,
			)
		}

		rule := monitoringv1.PrometheusRule{
			TypeMeta: metav1.TypeMeta{
				Kind:       monitoringv1.PrometheusRuleKind,
				APIVersion: monitoring.GroupName + "/" + monitoringv1.Version,
			},

			ObjectMeta: metav1.ObjectMeta{
				Name:      cm.Name + "-" + name,
				Namespace: cm.Namespace,
				Labels:    cm.Labels,
			},
			Spec: ruleSpec,
		}

		rules = append(rules, rule)
	}

	return rules, nil
}
