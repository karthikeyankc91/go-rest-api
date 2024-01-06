package parser

type Commands struct {
	ShowArp ShowArp `json:"show_arp,omitempty"`
}

type ShowArp struct {
	Output Output `json:"output,omitempty"`
}

type Output struct {
	Meta Meta `json:"meta,omitempty"`
}

type Meta struct {
	Command string `json:"command"`
}
