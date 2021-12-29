module gosynapse-example

go 1.16

require (
    gitlab.com/lama-corp/infra/packages/gosynapse 1.0.0
)

replace gitlab.com/lama-corp/infra/packages/gosynapse => ../
exclude gitlab.com/lama-corp/infra/packages/gosynapse 1.0.0
