package parser

import (
	"gopkg.in/yaml.v2"
	"os"
)

func Lex(configFile string) (map[interface{}]interface{}, error) {
	reader, err := os.Open(configFile)
	defer reader.Close()
	if err != nil {
		return nil, err
	}
	cfg := make(map[interface{}]interface{})
	err = yaml.NewDecoder(reader).Decode(&cfg)
	return cfg, err
}

func Parse(cfg map[interface{}]interface{}) error {
	if err := AstParse(jvsAstRoot, cfg); err != nil {
		return err
	}
	if err := jvsAstRoot.Link(); err != nil {
		return err
	}
	return nil
}

func LoadBuildInOptions(configFile string) error {
	cfg, err := Lex(configFile)
	if err != nil {
		panic(err)
	}
	if err := cfgToastItemRequired(cfg, "options", func(item interface{}) error {
		for name, option := range item.(map[interface{}]interface{}) {
			jvsAstRoot.Options[name.(string)] = newAstOption(name.(string))
			if err := AstParse(jvsAstRoot.Options[name.(string)], option.(map[interface{}]interface{})); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func init() {
	if err := LoadBuildInOptions("buildInOptions/global_options.yaml"); err != nil {
		panic("Error in loading buildInOptions/global_options.yaml :" + err.Error())
	}
}