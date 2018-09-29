package characters

// Storage defines a storage medium. It could be anything that implements this interface.
type Storage interface {
	Save(Character) error
	Load(string) (Character, error)
	Update(Character) error
}
