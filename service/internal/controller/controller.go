package controller

func NewWebRestCtl(in digIn) digOut {
	self := &packet{
		in: in,
		digOut: digOut{
			ExampleCtrl:            newExample(in),
		},
	}
	
	return self.digOut
}


type packet struct {
	in digIn
	
	digOut
}

type digIn struct {
	dig.In
	
	SysLogger   logger.ILogger `name:"sysLogger"`
	Request     handler.IRequestParse
	SetResponse response `optional:"true"`
	
	ExampleUseCase            exampleUseCaseIn
}

type digOut struct {
	dig.Out
	
	ExampleCtrl            IExampleCtrl

}


type exampleUseCaseIn struct {
	dig.In
	
	Get example.IGetExample
}
