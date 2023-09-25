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
	// Subscribe accepts an instance of IEventBus and returns listener map
	Subscribe(eb EventBusInterface) map[string][]ListenerClause
}
