package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_text2banner(t *testing.T) {

	got := text2banner("abc-defghijklmnopqrstuvxyzw")

	want := `
  ##   #####   ####     #####  ###### ######  ####  #    #  #       # #    # #      #    # #    #  ####  #####   ####  #####   ####  ##### #    # #    # #    # #    # ###### #    #
 #  #  #    # #    #    #    # #      #      #      #    #  #       # #   #  #      ##  ## ##   # #    # #    # #    # #    # #        #   #    # #    #  #  #   #  #      #  #    #
#    # #####  #         #    # #####  #####  #      ######  #       # ####   #      # ## # # #  # #    # #    # #    # #    #  ####    #   #    # #    #   ##     ##      #   #    #
###### #    # #         #    # #      #      #  ### #    #  #  #    # #  #   #      #    # #  # # #    # #####  #    # #####       #   #   #    # #    #   ##     ##     #    # ## #
#    # #    # #    #    #    # #      #      #    # #    #  #  #    # #   #  #      #    # #   ## #    # #      #  # # #   #  #    #   #   #    #  #  #   #  #    ##    #     ##  ##
#    # #####   ####     #####  ###### #       ####  #    #  #   ####  #    # ###### #    # #    #  ####  #       ####  #    #  ####    #    ####    ##   #    #   ##   ###### #    #
`

	if want != "\n"+got {
		t.Errorf("\ngot:\n%v\nwant:\n%v", got, want)
	}
}

func Test_main(t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"banner", "-t=M"}
	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old

	want := `
#    #
##  ##
# ## #
#    #
#    #
#    #
`
	if want != "\n"+string(out) {
		t.Errorf("\ngot:\n%v\nwant:\n%v", out, want)
	}
}
