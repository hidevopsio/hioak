package gitlab

import (
	"github.com/xanzy/go-gitlab"
		"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/jinzhu/copier"
	"github.com/hidevopsio/hioak/starter/scm"
)

type GroupMember struct {
	scm.GroupMember
	client ClientInterface
}

func NewGroupMember(c ClientInterface) scm.GroupMemberInterface {
	return &GroupMember{
		client: c,
	}
}

func (gm *GroupMember) GetGroupMember(token, baseUrl string, gid, uid int) (*scm.GroupMember,  error)  {
	log.Debug("group.ListGroups()")
	scmGroupMember := &scm.GroupMember{}
	gm.client.SetBaseURL(baseUrl + ApiVersion)
	log.Debug("before c.group.ListGroups(so)")
	opt := &gitlab.ListGroupMembersOptions{}
	groupMembers, _, err := gm.client.ListGroupMembers(gid, opt)
	log.Debug("after c.group member.groupMembers(so)")
	if err != nil {
		return nil, err
	}
	for _, groupMember := range groupMembers {
		if groupMember.ID == uid {
			copier.Copy(scmGroupMember, groupMember)
		}
	}
	return scmGroupMember, nil
}


func (gm *GroupMember) ListGroupMembers(token, baseUrl string, gid, uid int) (int,  error)  {
	log.Debug("group.ListGroups()")
	gm.client.SetBaseURL(baseUrl + ApiVersion)
	log.Debug("before c.group.ListGroups(so)")
	opt := &gitlab.ListGroupMembersOptions{}
	groupMembers, _, err := gm.client.ListGroupMembers(gid, opt)
	if err != nil {
		return 0, err
	}
	log.Debug("after gm.GroupMember.ListGroupMembers(so)")
	for _, groupMember := range groupMembers{
		if groupMember.ID == uid {
			for id, permissions := range scm.Permissions  {
				if groupMember.AccessLevel == id {
					return permissions.AccessLevelValue, nil
				}
			}
		}
	}
	return 0, err

}