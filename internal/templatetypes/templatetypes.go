package templatetypes

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"strings"
)

func indent(content string) (string, error) {
	out := ""
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		out += fmt.Sprintf("\t%s\n", scanner.Text())
	}
	return out, scanner.Err()
}

type Avconfig struct {
	XMLName          xml.Name         `xml:"avconfig"`
	Origin           string           `xml:"origin,attr"`
	Provider         string           `xml:"provider,attr"`
	ConnectionString string           `xml:"connectionString,attr"`
	ChannelTemplates ChannelTemplates `xml:"channeltemplates"`
}

func (a Avconfig) String() string {
	return fmt.Sprintf("\nAvconfig\n\tOrigin: %s\n\tProvider: %s\n"+
		"\tConnectionString: %s\n\tChannelTemplates:\n%s",
		a.Origin, a.Provider, a.ConnectionString, a.ChannelTemplates.String())
}

type ChannelTemplates struct {
	XMLName  xml.Name   `xml:"channeltemplates"`
	Revision int        `xml:"revision,attr"`
	Channels []Channels `xml:"channels"`
}

func (c ChannelTemplates) String() string {
	return fmt.Sprintf("|-Revision: %d\n", c.Revision)
}

type Channels struct {
	XMLName      xml.Name  `xml:"channels"`
	Name         string    `xml:"name,attr"`
	Default      bool      `xml:"default,attr"`
	Created      string    `xml:"created,attr"`
	Shared       bool      `xml:"shared,attr"`
	BasedOn      string    `xml:"based_on,attr"`
	HideFromUser bool      `xml:"hide_from_user,attr"`
	Channel      []Channel `xml:"channel"`
}

func (c Channels) String() string {
	return fmt.Sprintf("|--Name: %s\n|--Default: %t\n|--Created: %s\n|--Shared: %t\n|--BasedOn: %s\n|--HideFromUser: %t\n\n", c.Name, c.Default, c.Created, c.Shared, c.BasedOn, c.HideFromUser)
}

type Channel struct {
	XMLName                        xml.Name                  `xml:"channel"`
	Name                           string                    `xml:"name,attr"`
	Type                           int                       `xml:"type,attr"`
	TemplateType                   string                    `xml:"templatetype,attr"`
	Description                    string                    `xml:"description,attr"`
	Id                             string                    `xml:"id,attr"`
	Created                        string                    `xml:"created,attr"` // Todo: Cast to DateTime
	Autotake                       bool                      `xml:"autotake,attr"`
	AutotakeOffset                 int                       `xml:"autotakeo_ffset,attr"`
	SendToNcs                      bool                      `xml:"send_to_ncs,attr"`
	Recallnr                       int                       `xml:"recallnr,attr"`
	DefaultVariant                 bool                      `xml:"defaultvariant,attr"`
	Isstatevariant                 bool                      `xml:"isstatevariant,attr"`
	UseCommandOnTake               bool                      `xml:"use_command_on_take,attr"`
	UseCommandOnTakeOut            bool                      `xml:"use_command_on_take_out,attr"`
	CommandOnTakeOut               string                    `xml:"command_on_take_out,attr"`
	CommandOnTakeOutValue          string                    `xml:"command_on_take_out_value,attr"`
	CommandOnTakeOutValueParameter string                    `xml:"command_on_take_out_value_parameter,attr"`
	Preload                        bool                      `xml:"preload,attr"`
	Isfixedur                      bool                      `xml:"isfixedur,attr"`
	Fixeddur                       int                       `xml:"fixeddur,attr"`
	HideFromUser                   bool                      `xml:"hide_from_user,attr"`
	UseSequence                    bool                      `xml:"use_sequence,attr"`
	LoopSequence                   bool                      `xml:"loop_sequence,attr"`
	Sequence                       string                    `xml:"sequence,attr"`
	DefaultTcIn                    string                    `xml:"default_tc_in,attr"`
	DefaultTcOut                   string                    `xml:"default_tc_out,attr"`
	Pretake                        bool                      `xml:"pretake,attr"`
	Changed                        string                    `xml:"changed,attr"` // Todo: Cast to DateTime type
	CommandOnTake                  string                    `xml:"command_on_take,attr"`
	CommandOnTakeValue             string                    `xml:"command_on_take_value,attr"`
	CommandOnTakeValueParameter    string                    `xml:"command_on_take_value_parameter,attr"`
	SwitcherSetup                  SwitcherSetup             `xml:"switcher_setup"`
	AudioSetup                     AudioSetup                `xml:"audio_setup"`
	TakeCommands                   TakeCommands              `xml:"takecommands"`
	RoboticCameraControlSetup      RoboticCameraControlSetup `xml:"roboticcameracontrol_setup"`
	RouterSetup                    RouterSetup               `xml:"router_setup"`
}

//func (c Channel) String() string {
//
//}

type SwitcherSetup struct {
	XMLName              xml.Name             `xml:"switcher_setup"`
	Transitions          Transitions          `xml:"transitions"`
	SwitcherInputChannel SwitcherInputChannel `xml:"switcherinputchannel"`
	Auxiliaries          []Auxiliary          `xml:"auxilaries>auxilary"`
}

type Transitions struct {
	XMLName    xml.Name     `xml:"transitions"`
	Value      string       `xml:"value,attr"`
	Rate       int          `xml:"rate,attr"`
	Enable     bool         `xml:"enable,attr"`
	MixDelay   int          `xml:"mixdelay,attr"`
	Transition []Transition `xml:"transition"`
}

type Transition struct {
	xml.Name       `xml:"transition"`
	TransitionName string `xml:"name,attr"`
	Field          Field  `xml:"field"`
}

type Field struct {
	XMLName   xml.Name `xml:"field"`
	Name      string   `xml:"name,attr"`
	Fieldtype string   `xml:"fieldtype,attr"`
	Range     string   `xml:"range,attr"` // Todo: Cast into []int?
	Value     string   `xml:"value,attr"`
}

type Auxiliary struct {
	XMLName   xml.Name `xml:"auxilary"`
	Name      string   `xml:"name"`
	Bus       string   `xml:"bus"`
	Xpoint    string   `xml:"xpoint"`
	Audiolink string   `xml:"audiolink"`
	Index     string   `xml:"index"`
	SetDelay  string   `xml:"set_delay"`
	TransDur  string   `xml:"trans_dur"`
}

type SwitcherInputChannel struct {
	XMLName             xml.Name `xml:"switcherinputchannel"`
	Name                string   `xml:"name,attr"`
	Xpoint              string   `xml:"xpoint,attr"`
	Mixdelay            int      `xml:"mixdelay,attr"`
	XpointPp            string   `xml:"xpoint_pp,attr"`
	Mixeffect           string   `xml:"mixeffect,attr"`
	SetMeInPp           bool     `xml:"set_me_in_pp,attr"`
	CrosspointHold      bool     `xml:"crosspoint_hold,attr"`
	DskOff              bool     `xml:"dsk_off,attr"`
	Dsk2On              bool     `xml:"dsk_2_on,attr"`
	Dsk2OffDelay        int      `xml:"dsk_2_off_delay,attr"`
	Dsk3On              bool     `xml:"dsk_3_on,attr"`
	Dsk3OffDelay        int      `xml:"dsk_3_off_delay,attr"`
	Dsk4On              bool     `xml:"dsk_4_on,attr"`
	Dsk4OffDelay        int      `xml:"dsk_4_off_delay,attr"`
	MixDelay            int      `xml:"mix_delay,attr"`
	MixDelayOut         int      `xml:"mix_delay_out,attr"`
	Audiolink           string   `xml:"audiolink,attr"`
	DelayTransitiononly bool     `xml:"delay_transitiononly,attr"`
	Index               int      `xml:"index,attr"`
}

type AudioSetup struct {
	AudioChannels AudioChannels `xml:"audio_channels"`
}

func (a AudioSetup) String() string {
	indentedOut, err := indent(a.AudioChannels.String())
	if err != nil {
		panic(err) // Todo: handle error correctly!
	}
	return fmt.Sprintf("Audio Setup:\n%s", indentedOut)
}

type AudioChannels struct {
	XMLName          xml.Name           `xml:"audio_channels"`
	AudioChannelData []AudioChannelData `xml:"audio_channel"`
}

func (a AudioChannels) String() string {
	output := ""
	for _, v := range a.AudioChannelData {
		output += v.String() + "\n"
	}
	indentedOut, err := indent(output)
	if err != nil {
		panic(err) // Todo: handle error correctly!
	}
	return fmt.Sprintf("Audio Channels:\n%s", indentedOut)
}

type AudioChannelData struct {
	XMLName          xml.Name `xml:"audio_channel"`
	Name             string   `xml:"name,attr"`
	S3000Name        string   `xml:"s3000_name,attr"`
	Level            int      `xml:"level,attr"`
	Crossfadetime    int      `xml:"crossfadetime,attr"`
	CrossfadetimeOut int      `xml:"crossfadetime_out,attr"`
	Main             bool     `xml:"main,attr"`
	Keep             bool     `xml:"keep,attr"`
	Videolink        string   `xml:"videolink,attr"`
	Index            int      `xml:"index,attr"`
	Delay            int      `xml:"delay,attr"`
	Loudnesslevel    string   `xml:"loudnesslevel,attr"`
	Inputdevice      string   `xml:"inputdevice,attr"`
	Initiallevel     int      `xml:"initiallevel,attr"`
	Key              string   `xml:"key,attr"`
	CurrentLevel     int      `xml:"current_level,attr"`
	Action           string   `xml:"action,attr"`
	FaderNr          int      `xml:"faderNr,attr"`
	Fadertype        string   `xml:"fadertype,attr"`
	FadeUpInPreview  bool     `xml:"fade_up_in_preview,attr"`
	InitialLevel     int      `xml:"initial_level,attr"`
	InitialStandby   bool     `xml:"initial_standby,attr"`
	Isgroupfader     bool     `xml:"igroupfader,attr"`
	Isstandby        bool     `xml:"isstandby,attr"`
	Ismuted          bool     `xml:"ismuted,attr"`
	Channelmode      string   `xml:"channelmode,attr"`
	SkipCloseLevel   bool     `xml:"skip_close_level,attr"`
}

func (a AudioChannelData) String() string {
	return fmt.Sprintf("Name: %s\nS3000 Name: %s\nLevel: %d\nCrossfade Time: %d\nCrossfade Timeout: %d\n"+
		"Main: %t\nKeep: %t\nVideo Link: %s\nIndex: %d\nDelay: %d\nLoudness Level: %s\nInput Device: %s\n"+
		"Initial Level: %d\nKey: %s\nCurrent Level: %d\nAction: %s\nFader Number: %d\nFader Type: %s\n"+
		"Fade up in Preview: %t\nInitial Level: %d\nInitial Standby: %t\nIs Group Fader: %t\nIs Standby: %t\n"+
		"Is Muted: %t\nChannel Mode %s\nSkip Close Level: %t\n",
		a.Name, a.S3000Name, a.Level, a.Crossfadetime, a.CrossfadetimeOut, a.Main, a.Keep, a.Videolink, a.Index,
		a.Delay, a.Loudnesslevel, a.Inputdevice, a.InitialLevel, a.Key, a.CurrentLevel, a.Action, a.FaderNr, a.Fadertype,
		a.FadeUpInPreview, a.InitialLevel, a.InitialStandby, a.Isgroupfader, a.Isstandby, a.Ismuted, a.Channelmode, a.SkipCloseLevel)
}

type TakeCommands struct {
	XMLName     xml.Name      `xml:"takecommands"`
	TakeCommand []TakeCommand `xml:"takecommand"`
}

func (t TakeCommands) String() string {
	output := ""
	for _, v := range t.TakeCommand {
		output += v.String() + "\n"
	}
	indentedOut, err := indent(output)
	if err != nil {
		panic(err) // Todo: handle error correctly!
	}
	return indentedOut
}

type TakeCommand struct {
	XMLName   xml.Name `xml:"takecommand"`
	Command   string   `xml:"command,attr"`
	Value     string   `xml:"value,attr"`
	Parameter string   `xml:"parameter,attr"`
}

func (t TakeCommand) String() string {
	return fmt.Sprintf("Take Command:\n\tCommand: %s\n\tValue: %s\n\tParameter: %s", t.Command, t.Value, t.Parameter)
}

type RouterSetup struct {
	XMLName         xml.Name          `xml:"router_setup"`
	RouterSetupData []RouterSetupData `xml:"setups>setup"`
}

func (r RouterSetup) String() string {
	output := ""
	for _, v := range r.RouterSetupData {
		output += v.String() + "\n"
	}
	indentedOut, err := indent(output)
	if err != nil {
		panic(err) // Todo: handle error correctly!
	}
	return indentedOut
}

type RouterSetupData struct {
	XMLName         xml.Name `xml:"setup"`
	Matrix          string   `xml:"matrix,attr"`
	Level           string   `xml:"level,attr"`
	InputPreview    string   `xml:"input_preview,attr"`
	OutputPreview   string   `xml:"output_preview,attr"`
	InputProgram    string   `xml:"input_program,attr"`
	OutputProgram   string   `xml:"output_program,attr"`
	ProgramDelay    string   `xml:"program_delay,attr"`
	SalvoPreview    string   `xml:"salvo_preview,attr"`
	SalvoProgram    string   `xml:"salvo_program,attr"`
	InputAudiolink  string   `xml:"input_audiolink,attr"`
	OutputAudiolink string   `xml:"output_audiolink,attr"`
}

func (r RouterSetupData) String() string {
	return fmt.Sprintf("Router Setup Data:"+
		"\n\tMatrix: %s\n\tLevel: %s\n\tInput Preview: %s\n\tOutput Preview: %s"+
		"\n\tInput Program: %s\n\tOutput Program: %s\n\tProgram Delay: %s"+
		"\n\tSalvo Preview: %s\n\tSalvo Program: %s\n\tInput Audiolink: %s\n\tOutput Audiolink: %s",
		r.Matrix, r.Level, r.InputPreview, r.OutputPreview, r.InputProgram, r.OutputProgram,
		r.ProgramDelay, r.SalvoPreview, r.SalvoProgram, r.InputAudiolink, r.OutputAudiolink)
}

type RoboticCameraControlSetup struct {
	XMLName                       xml.Name                        `xml:"roboticcameracontrol_setup"`
	RoboticCameraControlSetupData []RoboticCameraControlSetupData `xml:"setups>setup"`
}

func (r RoboticCameraControlSetup) String() string {
	output := ""
	for _, v := range r.RoboticCameraControlSetupData {
		output += v.String() + "\n"
	}
	indentedOut, err := indent(output)
	if err != nil {
		panic(err) // Todo: handle error correctly!
	}
	return indentedOut
}

type RoboticCameraControlSetupData struct {
	XMLName         xml.Name `xml:"setup"`
	ShotName        string   `xml:"shot_name,attr"`
	CameraNumber    int      `xml:"camera_number,attr"`
	PagePreview     string   `xml:"page_preview,attr"`
	PresetPreview   string   `xml:"preset_preview,attr"`
	TimePreveiw     int      `xml:"time_preview,attr"`
	DelayPreview    int      `xml:"delay_preview,attr"`
	PageProgram     string   `xml:"page_program,attr"`
	PresetProgram   string   `xml:"preset_program,attr"`
	FadetimeProgram int      `xml:"fadetime_program,attr"`
	DelayProgram    string   `xml:"delay_program,attr"` // Todo: find way to cast "" as 0
	OnairProtect    string   `xml:"onair_protect,attr"`
}

func (r RoboticCameraControlSetupData) String() string {
	return fmt.Sprintf("RoboticCameraControlSetupData:"+
		"\n\tShot Name: %s\n\tCamera Number: %d\n\tPage Preview: %s\n\tPreset Preview: %s"+
		"\n\tTime Preview: %d\n\tDelay Preview: %d\n\tPage Program: %s\n\tPreset Program: %s"+
		"\n\tFadetime Program: %d\n\tDelay Program: %s\n\tOn Air Protect: %s",
		r.ShotName, r.CameraNumber, r.PagePreview, r.PresetPreview, r.TimePreveiw, r.DelayPreview,
		r.PageProgram, r.PresetProgram, r.FadetimeProgram, r.DelayProgram, r.OnairProtect)
}
