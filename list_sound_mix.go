package main

import (
	"io"

	"github.com/BurntSushi/goim/imdb"
)

func listSoundMixes(db *imdb.DB, r io.ReadCloser) {
	table := startSimpleLoad(db, "sound_mix",
		"atom_id", "outlet", "mix", "attrs")
	defer table.done()

	listAttrRows(r, table.atoms, func(id imdb.Atom, line, entity, row []byte) {
		var attrs []byte

		rowFields := splitListLine(row)
		if len(rowFields) == 0 {
			return // herp derp...
		}
		if len(rowFields) > 1 {
			attrs = rowFields[1]
		}
		ent := entityType("media", entity)
		table.add(line, id, ent.String(), unicode(rowFields[0]), unicode(attrs))
	})
}