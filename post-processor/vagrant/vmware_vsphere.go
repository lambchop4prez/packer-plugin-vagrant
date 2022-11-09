package vagrant

import (
	"bytes"
	"text/template"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type VMWareVsphereProvider struct{}

type vsphereVagrantfileTemplate struct {
	Host            string ""
	ComputeResource string ""
	ResourcePool    string ""
	TemplateName    string ""
}

func (p *VMWareVsphereProvider) KeepInputArtifact() bool {
	return true
}

func (p *VMWareVsphereProvider) Process(ui packersdk.Ui, artifact packersdk.Artifact, dir string) (vagrantfile string, metadata map[string]interface{}, err error) {
	metadata = map[string]interface{}{"provider": "vsphere"}

	tplData := &vsphereVagrantfileTemplate{}
	// TODO: find out how to get values from builder

	var contents bytes.Buffer
	t := template.Must(template.New("vf").Parse(defaultVsphereVagrantfile))
	err = t.Execute(&contents, tplData)
	vagrantfile = contents.String()
}

var defaultVsphereVagrantfile = `
Vagrant.configure("2") do |config|
	config.vm.provider :vsphere do |vsphere|
		vsphere.host = {{ .Host }}
		vsphere.compute_resource_name = {{ .ComputeResource }}
		vsphere.resource_pool_name = {{ .ResourcePool }}
		vsphere.template_name = {{ .TemplateName }}
	end
end
`
