package lib

import "crypto/sha1"
import "fmt"
import "strings"
import "time"

// Commit represents an entire repository for a snapshot of time.
type Commit struct {
	Tree   string    // Hash that points to the root tree of a commit.
	Parent string    // Hash of parent commit.
	Msg    string    // Commit message.
	Author string    // Author of commit (TODO: Cryptographically signed).
	Date   time.Time // Timestamp of commit.
}

// GetCommit fetches a commit from the k/v store by it's hash value.
func GetCommit(k string) (*Commit, error) {
	contents, err := getObject(k)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(
		strings.TrimRight(contents, "\n"),
		"\n",
	)

	tree := strings.Split(lines[1], "\t")[1]
	parent := strings.Split(lines[2], "\t")[1]
	author := strings.Split(lines[3], "\t")[1]
	date, err := time.Parse(
		"2006-01-02 15:04:05.000000000 -0700 MST",
		strings.Split(lines[4], "\t")[1],
	)
	msg := strings.TrimLeft(strings.Join(lines[6:], "\n"), "\t")
	if err != nil {
		return nil, err
	}

	c := &Commit{
		Tree:   tree,
		Parent: parent,
		Author: author,
		Date:   date,
		Msg:    msg,
	}
	return c, nil
}

// Put persists a commit into the k/v store.
func (c *Commit) Put() error {
	return putObject(c)
}

// String returns the string representation of a commit.
func (c *Commit) String() string {
	return fmt.Sprintf(
		"Commit:\t%s\nTree:\t%s\nParent:\t%s\nAuthor:\t%s\nDate:\t%s\n\n\t%s\n",
		c.Hash(),
		c.Tree,
		c.Parent,
		c.Author,
		c.Date,
		c.Msg,
	)
}

// Hash returns the SHA1 hash representation of a commit.
func (c *Commit) Hash() string {
	h := sha1.New()
	h.Write([]byte(c.Tree))
	h.Write([]byte(c.Parent))
	h.Write([]byte(c.Author))
	h.Write([]byte(c.Date.String()))
	h.Write([]byte(c.Msg))

	return fmt.Sprintf("%x", h.Sum(nil))
}
