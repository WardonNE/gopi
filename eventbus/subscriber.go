package eventbus

// Subscriber is subscriber interface
//
// example:
//
//  type AuthSubscriber struct {
//  }
//
//  func (s *AuthSubscriber) OnLogin(event IEvent) bool {
//      // do something
//      return true
//  }
//
//  func (s *AuthSubscriber) OnLogout(event IEvent) bool {
//      // do something
//      return true
//  }
//
//  func (s *AuthSubscriber) Subscribe() map[string][]ListenerClause {
//      return map[string][]ListenerClause{
//          "onLogin": []ListenerClause{
//              s.OnLogin,
//          },
//          "onLogout": []ListenerClause{
//              s.OnLogout,
//          },
//      }
//  }
//
type Subscriber interface {
	// Subscribe returns listener map
	Subscribe() map[string][]ListenerClause
}
