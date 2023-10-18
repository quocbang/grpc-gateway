package repositories

type AffectedRows int64

type CommonUpdateReply struct {
	AffectedRows AffectedRows
}
