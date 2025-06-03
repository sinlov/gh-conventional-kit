package urfave_cli

import (
	"fmt"
	"github.com/sinlov/gh-conventional-kit/command"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"testing"
)

func TestGlobalFlag(t *testing.T) {
	t.Logf("~> mock GlobalFlag")
	// mock GlobalFlag

	t.Logf("~> do GlobalFlag")
	// do GlobalFlag
	flags := UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())

	// verify GlobalFlag
	assert.NotEqual(t, 0, len(flags))
}

func TestUrfaveCliAppendCliFlag(t *testing.T) {
	// mock UrfaveCliAppendCliFlag
	type args struct {
		target []cli.Flag
		elem   []cli.Flag
	}
	tests := []struct {
		name            string
		args            args
		wantResult      []cli.Flag
		wantPanic       bool
		wantPanicErrMsg string
	}{
		{
			name: "exists",
			args: args{
				target: []cli.Flag{
					&cli.StringFlag{
						Name:  "config.new_arg,new_arg",
						Usage: "new arg",
					},
					&cli.BoolFlag{
						Name:  "config.debug",
						Usage: "debug",
					},
				},
				elem: []cli.Flag{
					&cli.BoolFlag{
						Name:  "config.debug",
						Usage: "debug",
					},
				},
			},
			wantPanic:       true,
			wantPanicErrMsg: fmt.Sprintf("do UrfaveCliAppendCliFlag err, flag exists name %s at %v", "config.debug", []string{"config.debug"}),
		},
		{
			name: "success",
			args: args{
				target: []cli.Flag{
					&cli.StringFlag{
						Name:  "config.new_arg,new_arg",
						Usage: "new arg",
					},
				},
				elem: []cli.Flag{
					&cli.BoolFlag{
						Name:  "config.debug",
						Usage: "debug",
					},
				},
			},
			wantResult: []cli.Flag{
				&cli.BoolFlag{
					Name:  "config.debug",
					Usage: "debug",
				},
				&cli.StringFlag{
					Name:  "config.new_arg,new_arg",
					Usage: "new arg",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if tc.wantPanic {
				if !assert.PanicsWithError(t, tc.wantPanicErrMsg, func() {
					// do UrfaveCliAppendCliFlag with panic
					_ = UrfaveCliAppendCliFlag(tc.args.elem, tc.args.target)
				}) {
					t.Fatalf("UrfaveCliAppendCliFlag should be panic with test case %s", tc.name)
				}
				return
			}

			// do UrfaveCliAppendCliFlag
			gotResult := UrfaveCliAppendCliFlag(tc.args.elem, tc.args.target)

			// verify UrfaveCliAppendCliFlag
			assert.Equal(t, tc.wantResult, gotResult)
		})
	}
}

func TestUrfaveCliAppendCliFlags(t *testing.T) {
	// mock UrfaveCliAppendCliFlags
	type args struct {
		target []cli.Flag
		elem   [][]cli.Flag
	}
	tests := []struct {
		name            string
		args            args
		wantResult      interface{}
		wantPanic       bool
		wantPanicErrMsg string
	}{
		{
			name: "exists",
			args: args{
				target: []cli.Flag{
					&cli.StringFlag{
						Name:  "config.new_arg,new_arg",
						Usage: "new arg",
					},
					&cli.BoolFlag{
						Name:  "config.debug",
						Usage: "debug",
					},
				},
				elem: [][]cli.Flag{
					{
						&cli.BoolFlag{
							Name:  "config.debug",
							Usage: "debug",
						},
					},
				},
			},
			wantPanic:       true,
			wantPanicErrMsg: fmt.Sprintf("do UrfaveCliAppendCliFlag err, flag exists name %s at %v", "config.debug", []string{"config.debug"}),
		},

		{
			name: "success",
			args: args{
				target: []cli.Flag{
					&cli.StringFlag{
						Name:  "config.new_arg,new_arg",
						Usage: "new arg",
					},
				},
				elem: [][]cli.Flag{
					{
						&cli.BoolFlag{
							Name:  "config.debug",
							Usage: "debug",
						},
					},
					{
						&cli.UintFlag{
							Name:  "config.timeout_second",
							Usage: "do request timeout setting second.",
						},
					},
				},
			},
			wantResult: []cli.Flag{
				&cli.StringFlag{
					Name:  "config.new_arg,new_arg",
					Usage: "new arg",
				},
				&cli.BoolFlag{
					Name:  "config.debug",
					Usage: "debug",
				},
				&cli.UintFlag{
					Name:  "config.timeout_second",
					Usage: "do request timeout setting second.",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			if tc.wantPanic {
				if !assert.PanicsWithError(t, tc.wantPanicErrMsg, func() {
					// do UrfaveCliAppendCliFlags with panic
					_ = UrfaveCliAppendCliFlags(tc.args.target, tc.args.elem...)
				}) {
					t.Fatalf("UrfaveCliAppendCliFlags should be panic with test case %s", tc.name)
				}
				return
			}

			// do UrfaveCliAppendCliFlags
			gotResult := UrfaveCliAppendCliFlags(tc.args.target, tc.args.elem...)

			// verify UrfaveCliAppendCliFlags
			assert.Equal(t, tc.wantResult, gotResult)
		})
	}
}
