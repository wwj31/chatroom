package dirty

import (
	"fmt"
	"testing"
)

func TestDirty(t *testing.T) {
	fmt.Println(CheckDirty("fjeaabdfagotssiow f u c k hefewwio"))

	fmt.Println(CheckDirty("fjfeeiow fuck hewifewfwo"))

	fmt.Println(CheckDirty("fjefeierewow f uc k hewfefefagio"))

	fmt.Println(CheckDirty("fjeggriow fu c k hewio"))

	fmt.Println(CheckDirty("dlckfjeiow fuc k hewio"))

	fmt.Println(CheckDirty("fjegriow fuc k hewioejudlcklngejaculated"))

	fmt.Println(CheckDirty("ejaculating fjegriow fuc k ejaculated"))

	fmt.Println(CheckDirty("dinkf u c k e r f u c k e rfaggixxxng"))

}
