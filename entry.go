package cwl

// Entry represents fs entry, it means [File|Directory|Dirent]
type Entry struct {
	Class    string
	Location string
	Path     string
	Basename string
	File
	Directory
	Dirent
}

// File represents file entry.
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#File
type File struct {
	Dirname string
	Size    int64
	Format  string
}

// Directory represents direcotry entry.
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#Directory
type Directory struct {
	Listing []Entry
}

// Dirent represents ?
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#Dirent
type Dirent struct {
	Entry     string
	EntryName string
	Writable  bool
}

// NewList constructs a list of Entry from interface
func (_ Entry) NewList(i interface{}) []Entry {
	dest := []Entry{}
	switch x := i.(type) {
	case string:
		dest = append(dest, Entry{}.New(x))
	case []interface{}:
		for _, v := range x {
			dest = append(dest, Entry{}.New(v))
		}
	}
	return dest
}

// New constructs an Entry from interface
func (_ Entry) New(i interface{}) Entry {
	dest := Entry{}
	switch x := i.(type) {
	case string:
		dest.Location = x
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "entryname":
				dest.EntryName = v.(string)
			case "entry":
				dest.Entry = v.(string)
			}
		}
	}
	return dest
}
