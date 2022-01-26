package pkg

import (
	"github.com/GuessWhoSamFoo/experiment/manifests/imports/cartorun"
	"github.com/GuessWhoSamFoo/experiment/manifests/imports/k8s"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s"
)

var Commit string

type Example struct {
	cdk8s.ChartProps
}

func NewExample(scope constructs.Construct, id string, props *Example) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	// TODO: Pass kplus objects to chart
	// kplus.ServiceAccount_FromServiceAccountName(jsii.String("service-account-with-role-to-create-taskruns"))

	k8s.NewKubeServiceAccount(chart, jsii.String("taskrun-sa"), &k8s.KubeServiceAccountProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("service-account-with-role-to-create-taskruns"),
		},
	})

	k8s.NewKubeRoleBinding(chart, jsii.String("test-role-binding"), &k8s.KubeRoleBindingProps{
		RoleRef: &k8s.RoleRef{
			ApiGroup: jsii.String("rbac.authorization.k8s.io"),
			Kind:     jsii.String("Role"),
			Name:     jsii.String("testing-role"),
		},
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("testing-role-binding"),
		},
		Subjects: &[]*k8s.Subject{
			{
				Kind: jsii.String("ServiceAccount"),
				Name: jsii.String("service-account-with-role-to-create-taskruns"),
			},
		},
	})

	k8s.NewKubeRole(chart, jsii.String("test-role"), &k8s.KubeRoleProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("testing-role"),
		},
		Rules: &[]*k8s.PolicyRule{
			{
				ApiGroups: &[]*string{jsii.String("tekton.dev")},
				Resources: &[]*string{jsii.String("tasks"), jsii.String("taskruns")},
				Verbs: &[]*string{
					jsii.String("list"),
					jsii.String("create"),
					jsii.String("update"),
					jsii.String("delete"),
					jsii.String("patch"),
					jsii.String("watch"),
					jsii.String("get"),
				},
			},
		},
	})

	k8s.NewKubeRoleBinding(chart, jsii.String("carto-role-binding"), &k8s.KubeRoleBindingProps{
		RoleRef: &k8s.RoleRef{
			ApiGroup: jsii.String("rbac.authorization.k8s.io"),
			Kind:     jsii.String("Role"),
			Name:     jsii.String("carto-role"),
		},
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("carto-role-binding"),
		},
		Subjects: &[]*k8s.Subject{
			{
				Kind: jsii.String("ServiceAccount"),
				Name: jsii.String("service-account-with-role-to-create-taskruns"),
			},
		},
	})

	k8s.NewKubeRole(chart, jsii.String("carto-role"), &k8s.KubeRoleProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String("carto-role"),
		},
		Rules: &[]*k8s.PolicyRule{
			{
				ApiGroups: &[]*string{jsii.String("carto.run")},
				Resources: &[]*string{jsii.String("runnables")},
				Verbs: &[]*string{
					jsii.String("list"),
					jsii.String("create"),
					jsii.String("update"),
					jsii.String("delete"),
					jsii.String("patch"),
					jsii.String("watch"),
					jsii.String("get"),
				},
			},
		},
	})

	// TODO: Fix patching
	//t := tektondev.NewTask(chart, jsii.String("task"), &tektondev.TaskProps{
	//	Metadata: &cdk8s.ApiObjectMetadata{
	//		Labels: &map[string]*string{
	//			"apps.tanzu.vmware.com/task": jsii.String("test"),
	//		},
	//		Name: jsii.String("test"),
	//	},
	//})
	//t.AddJsonPatch(cdk8s.JsonPatch_Add(jsii.String("/spec"), "123"))

	//t.AddJsonPatch(cdk8s.JsonPatch_Add(jsii.String("path"), map[string]interface{}{
	//	"params": []map[string]interface{}{
	//		{
	//			"name": "blob-url",
	//		},
	//		{
	//			"name": "blob-revision",
	//		},
	//	},
	//	"steps": []map[string]interface{}{
	//		{
	//			"name":  "test",
	//			"image": "golang",
	//			"command": []string{
	//				"bash",
	//				"-cxe",
	//				fmt.Sprintf("set -o pipefail\n          cd `mktemp -d`\n          git clone $(params.blob-url) && cd \"`basename $(params.blob-url) .git`\"\n          git checkout $(params.blob-revision)\n          go test -v ./..."),
	//			},
	//		},
	//	},
	//}))

	cartorun.NewClusterRunTemplate(chart, jsii.String("cluster-run"), &cartorun.ClusterRunTemplateProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("tekton-taskrun"),
		},
		Spec: &cartorun.ClusterRunTemplateSpec{
			Template: &map[string]interface{}{
				"apiVersion": "tekton.dev/v1beta1",
				"kind":       "TaskRun",
				"metadata": &map[string]interface{}{
					"generateName": "test-",
				},
				"spec": &map[string]interface{}{
					"taskRef": &map[string]interface{}{
						"name": "test",
					},
					"params": &[]map[string]interface{}{
						{
							"name":  "blob-url",
							"value": "$(runnable.spec.inputs.blob-url)$",
						},
						{
							"name":  "blob-revision",
							"value": "$(runnable.spec.inputs.blob-revision)$",
						},
					},
				},
			},
			Outputs: &map[string]*string{
				"url":      jsii.String("spec.params[?(@.name==\"blob-url\")].value"),
				"revision": jsii.String("spec.params[?(@.name==\"blob-revision\")].value"),
			},
		},
	})

	cartorun.NewRunnable(chart, jsii.String("runnable"), &cartorun.RunnableProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("test"),
		},
		Spec: &cartorun.RunnableSpec{
			RunTemplateRef: &cartorun.RunnableSpecRunTemplateRef{
				Name: jsii.String("tekton-taskrun"),
			},
			Inputs: &map[string]interface{}{
				"blob-url":      jsii.String("https://github.com/kontinue/hello-world"),
				"blob-revision": jsii.String(Commit),
			},
			ServiceAccountName: jsii.String("service-account-with-role-to-create-taskruns"),
		},
	})

	return chart
}
