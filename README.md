# cwl.go

`cwl.go` is just a parser of CWL file and input files based on [CWL](https://github.com/common-workflow-language/common-workflow-language), for example [1st-tool.yaml](https://github.com/common-workflow-language/common-workflow-language/blob/master/v1.0/examples/1st-tool.cwl) and [echo-job.yml](https://github.com/common-workflow-language/common-workflow-language/blob/master/v1.0/examples/echo-job.yml).

# Example

Given

```yaml
cwlVersion: v1.0
class: CommandLineTool
baseCommand: echo
inputs:
  message:
    type: string
    inputBinding:
      position: 1
outputs: []
```

do

```go
f, _ := os.Open("given-cwl.yaml")
root := cwl.NewCWL()
err := cwl.Decode(f, root)
```

then

```go
// root
&cwl.Root{
  Version:     "v1.0",
  Class:       "CommandLineTool",
  BaseCommand: "echo",
  Inputs:      map[string]RequiredInput{
    "message": RequiredInput{
      Name:    "message",
      Type:    reflect.String,
      Binding: InputBinding{
        Position: 1,
      },
    }
  },
  Outputs: []string{},
}
```
