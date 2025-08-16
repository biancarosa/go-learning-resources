- `staticcheck`` - https://github.com/dominikh/go-tools#installation

-> go install honnef.co/go/tools/cmd/staticcheck@2025.1.1

-> staticcheck main.go


# Dependencias do seu sistema

--- Logrus

# Dependencias do seu ambiente de dev / CI

--- Vet (built-in)
--- Staticcheck @ 2025.1.1

---- git hooks? antes de commitar, por favor rode vet & staticcheck
---- makefile? comando pra rodar esses 2 tools
---- CI - qual versão que o Staticcheck usa no seu CI vs na sua maquina?

--- Desafio: consistencia entre versões de maquinas de devs, CI, etc