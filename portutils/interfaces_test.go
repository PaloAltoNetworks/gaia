package portutils

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_PortsList(t *testing.T) {

	Convey("Given a PortsList", t, func() {
		pl := PortsList{}

		Convey("I should be able to add a value", func() {
			pl = append(pl, 80, 90)
			So(len(pl), ShouldEqual, 2)
			So(pl[0], ShouldEqual, 80)
			So(pl[1], ShouldEqual, 90)
		})
	})
}

func Test_PortsRange(t *testing.T) {

	Convey("Given a list of PortsRange", t, func() {
		prs := &PortsRangeList{}

		Convey("I should be able to add a value", func() {
			*prs = append(*prs, &PortsRange{FromPort: 80, ToPort: 443})
			So(len(*prs), ShouldEqual, 1)
			So((*prs)[0].FromPort, ShouldEqual, 80)
			So((*prs)[0].ToPort, ShouldEqual, 443)

		})
	})
}

func Test_PortsRange_HasOverlapWithPort(t *testing.T) {

	type args struct {
		pr   *PortsRange
		port int64
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "overlap",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				port: 81,
			},
			want: true,
		},
		{
			name: "no overlap left",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				port: 79,
			},
			want: false,
		},
		{
			name: "no overlap right",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				port: 444,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.pr.HasOverlapWithPort(tt.args.port)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasOverlapWithPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PortsRange_HasOverlapWithPortsList(t *testing.T) {

	type args struct {
		pr *PortsRange
		pl *PortsList
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "overlap",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				pl: &PortsList{
					70,
					85,
					500,
				},
			},
			want: true,
		},
		{
			name: "no overlap",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				pl: &PortsList{
					70,
					500,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.pr.HasOverlapWithPortsList(tt.args.pl)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasOverlapWithPortsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PortsRange_HasOverlapWithPortsRange(t *testing.T) {

	type args struct {
		pr  *PortsRange
		opr *PortsRange
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "overlap left",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				opr: &PortsRange{
					FromPort: 70,
					ToPort:   81,
				},
			},
			want: true,
		},
		{
			name: "overlap left",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				opr: &PortsRange{
					FromPort: 442,
					ToPort:   500,
				},
			},
			want: true,
		},
		{
			name: "no overlap left",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				opr: &PortsRange{
					70,
					79,
				},
			},
			want: false,
		},
		{
			name: "no overlap right",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				opr: &PortsRange{
					444,
					500,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.pr.HasOverlapWithPortsRange(tt.args.opr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasOverlapWithPortsRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PortsRange_ContainsPortsRange(t *testing.T) {

	type args struct {
		pr  *PortsRange
		opr *PortsRange
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "exact match",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				opr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
			},
			want: true,
		},
		{
			name: "contains",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				opr: &PortsRange{
					FromPort: 81,
					ToPort:   442,
				},
			},
			want: true,
		},
		{
			name: "overlap not contain",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				opr: &PortsRange{
					FromPort: 81,
					ToPort:   444,
				},
			},
			want: false,
		},
		{
			name: "no contain",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				opr: &PortsRange{
					FromPort: 600,
					ToPort:   700,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.pr.ContainsPortsRange(tt.args.opr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContainsPortsRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PortsRange_HasOverlapWithPortsRanges(t *testing.T) {

	type args struct {
		pr  *PortsRange
		prs *PortsRangeList
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "overlap left",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				prs: &PortsRangeList{
					&PortsRange{
						FromPort: 50,
						ToPort:   79,
					},
					&PortsRange{
						FromPort: 90,
						ToPort:   100,
					},
				},
			},
			want: true,
		},
		{
			name: "overlap right",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				prs: &PortsRangeList{
					&PortsRange{
						FromPort: 50,
						ToPort:   79,
					},
					&PortsRange{
						FromPort: 440,
						ToPort:   500,
					},
				},
			},
			want: true,
		},
		{
			name: "no overlap",
			args: args{
				pr: &PortsRange{
					FromPort: 80,
					ToPort:   443,
				},
				prs: &PortsRangeList{
					&PortsRange{
						FromPort: 50,
						ToPort:   79,
					},
					&PortsRange{
						FromPort: 500,
						ToPort:   1000,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.pr.HasOverlapWithPortsRanges(tt.args.prs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasOverlapWithPortsRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PortsList_HasOverlapWithPort(t *testing.T) {

	type args struct {
		pl   *PortsList
		port int64
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "overlap",
			args: args{
				pl: &PortsList{
					80,
					443,
				},
				port: 80,
			},
			want: true,
		},
		{
			name: "no overlap",
			args: args{
				pl: &PortsList{
					80,
					443,
				},
				port: 8080,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.pl.HasOverlapWithPort(tt.args.port)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasOverlapWithPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PortsList_HasOverlapWithPortsList(t *testing.T) {

	type args struct {
		pl  *PortsList
		opl *PortsList
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "overlap",
			args: args{
				pl: &PortsList{
					80,
					443,
				},
				opl: &PortsList{
					80,
					500,
				},
			},
			want: true,
		},
		{
			name: "no overlap",
			args: args{
				pl: &PortsList{
					80,
					443,
				},
				opl: &PortsList{
					50,
					500,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.pl.HasOverlapWithPortsList(tt.args.opl)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasOverlapWithPortsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PortsList_HasOverlapWithPortsRanges(t *testing.T) {

	type args struct {
		pl  *PortsList
		prs *PortsRangeList
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "overlap left",
			args: args{
				pl: &PortsList{
					80,
					443,
				},
				prs: &PortsRangeList{
					&PortsRange{
						FromPort: 80,
						ToPort:   100,
					},
				},
			},
			want: true,
		},
		{
			name: "overlap right",
			args: args{
				pl: &PortsList{
					80,
					443,
				},
				prs: &PortsRangeList{
					&PortsRange{
						FromPort: 440,
						ToPort:   500,
					},
				},
			},
			want: true,
		},
		{
			name: "no overlap",
			args: args{
				pl: &PortsList{
					80,
					443,
				},
				prs: &PortsRangeList{
					&PortsRange{
						FromPort: 450,
						ToPort:   500,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.pl.HasOverlapWithPortsRanges(tt.args.prs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasOverlapWithPortsRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
