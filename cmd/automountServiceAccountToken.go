package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func checkAutomountServiceAccountToken(result *Result) {

	// Check for use of deprecated service account name
	if result.dsa != "" {
		result.err = 1
	}

	if result.token != nil {
		// automountServiceAccountToken = true, and serviceAccountName is blank (default: default)
		if *result.token && result.sa == "" {
			result.err = 2
		}
	} else {
		// automountServiceAccountToken = nil (default: true), and serviceAccountName is blank (default: default)
		if result.sa == "" {
			result.err = 3
		}
	}
}

func printResultASAT(results []Result) {

	for _, result := range results {
		if result.dsa != "" {
			log.WithFields(log.Fields{
				"type":               result.kubeType,
				"namespace":          result.namespace,
				"name":               result.name,
				"serviceAccount":     result.dsa,
				"serviceAccountName": result.sa,
			}).Warn("deprecated serviceAccount detected (sub for serviceAccountName)")
		}

		if result.err == 2 {
			log.WithFields(log.Fields{
				"type":      result.kubeType,
				"namespace": result.namespace,
				"name":      result.name,
			}).Error("automountServiceAccountToken = true with no serviceAccountName")
		} else if result.err == 3 {
			log.WithFields(log.Fields{
				"type":      result.kubeType,
				"namespace": result.namespace,
				"name":      result.name,
			}).Error("automountServiceAccountToken nil (mounted by default) with no serviceAccountName")
		}

	}
}

func auditAutomountServiceAccountToken(items Items) (results []Result) {
	for _, item := range items.Iter() {
		result := ServiceAccountIter(item)
		checkAutomountServiceAccountToken(result)

		if result.err > 0 {
			results = append(results, *result)
		}
	}

	printResultASAT(results)
	defer wg.Done()
	return
}

// satCmd represents the sat command
var satCmd = &cobra.Command{
	Use:   "sat",
	Short: "Audit automountServiceAccountToken = true pods against an empty (default) service account",
	Long: `This command determines which pods are running with
autoMountServiceAcccountToken = true and default service account names.
	
An ERROR log is generated when a container matches one of the fol:
  automountServiceAccountToken = true and serviceAccountName is blank (default: default)
  automountServiceAccountToken = nil  and serviceAccountName is blank (default: default)

A WARN log is generated when a pod is found using Pod.Spec.DeprecatedServiceAccount
Fix this by updating serviceAccount to serviceAccountName in your .yamls

Example usage:
kubeaudit rbac sat`,
	Run: func(cmd *cobra.Command, args []string) {
		kube, err := kubeClient(rootConfig.kubeConfig)
		if err != nil {
			log.Error(err)
		}

		if rootConfig.json {
			log.SetFormatter(&log.JSONFormatter{})
		}

		// fetch deployments, statefulsets, daemonsets
		// and pods which do not belong to another abstraction
		deployments := getDeployments(kube)
		statefulSets := getStatefulSets(kube)
		daemonSets := getDaemonSets(kube)
		pods := getPods(kube)
		replicationControllers := getReplicationControllers(kube)

		wg.Add(5)
		go auditAutomountServiceAccountToken(kubeAuditStatefulSets{list: statefulSets})
		go auditAutomountServiceAccountToken(kubeAuditDaemonSets{list: daemonSets})
		go auditAutomountServiceAccountToken(kubeAuditPods{list: pods})
		go auditAutomountServiceAccountToken(kubeAuditReplicationControllers{list: replicationControllers})
		go auditAutomountServiceAccountToken(kubeAuditDeployments{list: deployments})
		wg.Wait()
	},
}

func init() {
	rbacCmd.AddCommand(satCmd)
}
