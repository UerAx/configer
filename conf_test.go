/*
 * @Author: ww
 * @Date: 2022-07-03 17:35:13
 * @Description:
 * @FilePath: /goconf/conf.go
 */

package goconf

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/uerax/goconf/category"
)

func TestNewCfgFile(t *testing.T) {
	tests := []struct {
		name string
		want *CfgFile
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCfgFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCfgFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCfgFile_New(t *testing.T) {
	tests := []struct {
		name string
		tr   *CfgFile
		want *CfgFile
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CfgFile.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCfgFile_ReadAll(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		tr      *CfgFile
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.ReadAll(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("CfgFile.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCfgFile_ReadConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		tr      *CfgFile
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.ReadConfig(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("CfgFile.ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCfgFile_GetValue(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name    string
		tr      *CfgFile
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.GetValue(tt.args.in...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CfgFile.GetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CfgFile.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCfgFile_Unmarshal4Name(t *testing.T) {
	type args struct {
		name string
		obj  interface{}
	}
	tests := []struct {
		name    string
		tr      *CfgFile
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.Unmarshal4Name(tt.args.name, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("CfgFile.Unmarshal4Name() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCfgFile_Reload(t *testing.T) {
	i, _ := category.ReadToml("./testdata/conf.toml", struct{}{})
	fmt.Println(i)
}
