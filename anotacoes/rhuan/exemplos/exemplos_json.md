## exmplos de estruturas

### create

- exemplo da estrutura em go:

```golang
table := Table{
	Name: "rhuan patriky",
	Columns: []string{
		"column_1",
		"column_2",
		"column_3",
		"column_4",
	},
	Rows: []Row{
		{
			Row: []string{
				"row_1",
				"row_2",
				"row_3",
				"row_4",
			},
		},
		{
			Row: []string{
				"row_5",
				"row_6",
				"row_7",
				"row_8",
			},
		},
    },
}
```

- exemplo da estrutura em json:

```json
{
    "name": "table",
    "columns": [
        "column_1",
        "column_2",
        "column_3"
    ],
    "rows": [
        {
            "row": [
                "17",
                "'row'",
                "'2022-12-06 00:00:00'"
            ]
        },
        {
            "row": [
                "18",
                "'row'",
                "'2022-12-06 00:00:00'"
            ]
        }
    ]
}
```

### update

- exemplo da estrutura em go:

```golang
table := Table{
	Name: "rhuan patriky",
	Columns: []string{
		"column_1",
		"column_2",
		"column_3",
		"column_4",
	},
	Rows: []Row{
		{
			Row: []string{
				"row_1",
				"row_2",
				"row_3",
				"row_4",
			},
		},
		{
			Row: []string{
				"row_5",
				"row_6",
				"row_7",
				"row_8",
			},
		},
	},
}
```

- exemplo da estrutura em json:

```json
{
    "name": "table",
    "columns": [
        "column_1",
        "column_2",
        "column_3"
    ],
    "rows": [
        {
            "row": [
                "17",
                "'row'",
                "'2022-12-06 00:00:00'"
            ]
        },
        {
            "row": [
                "18",
                "'row'",
                "'2022-12-06 00:00:00'"
            ]
        }
    ],
	"optional": "99"
}
```
