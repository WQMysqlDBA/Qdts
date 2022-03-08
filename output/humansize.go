package output

type HumanOutput struct {
	ByteSize int64
}

type HumanSizeOutput interface {
	HumanByteSize() (h int64, s string)
}

func NewHumanSizeMessage(s int64) (humansize int64,Unit string){
	var msg HumanSizeOutput
	msg = NewHumanOutput(s)
	humansize , Unit = msg.HumanByteSize()
	return
}

func (HumanOutput *HumanOutput) HumanByteSize() (h int64, s string) {
	const kb = 1024
	const mb = 1024 * kb
	const gb = 1024 * mb
	var humanSize int64
	var humanUnit string
	if HumanOutput.ByteSize > gb {
		humanSize = HumanOutput.ByteSize / gb
		humanUnit = "Gib"
	} else if HumanOutput.ByteSize > mb {
		humanSize = HumanOutput.ByteSize / mb
		humanUnit = "Mib"
	} else if HumanOutput.ByteSize > kb {
		humanSize = HumanOutput.ByteSize / kb
		humanUnit = "Kib"
	} else {
		humanSize = HumanOutput.ByteSize
		humanUnit = "Bit"
	}
	return humanSize, humanUnit
}

func NewHumanOutput(bytesize int64) *HumanOutput {
	return &HumanOutput{
		ByteSize: bytesize,
	}
}
