- `staticcheck`` - https://github.com/dominikh/go-tools#installation

-> go install honnef.co/go/tools/cmd/staticcheck@2025.1.1

-> staticcheck main.go


-> cat tools.go | grep _ | awk -F'"' '{print $2}' | xargs -tI % go install %



(https://marcofranssen.nl/manage-go-tools-via-go-modules)