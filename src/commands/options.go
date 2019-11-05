package commands

type Options struct {
	Options []string
}

func (o Options) hasOption(shortName string, longName string) bool {
  shortName = "-" + shortName;
  longName = "--" + longName;

  for _, currentOption := range o.Options {
  	if currentOption == shortName || currentOption == longName {
  		return true
  	}
  }

  return false
}

