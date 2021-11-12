# mutiny

A library for the [Revolt](https://revolt.chat/) bot API, heavily inspired by [arikawa](https://github.com/diamondburned/arikawa) and [discordgo](https://github.com/bwmarrin/discordgo).

Note that this library **will not** support user accounts, only bot accounts (which ironically is pretty much the opposite of the libraries we're inspired by)

## Acknowledgements

Some of the code here (especially the gateway code) is based on or copied from [arikawa](https://github.com/diamondburned/arikawa) or [discordgo](https://github.com/bwmarrin/discordgo).

## Example

A basic ping-pong example:

```go
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/starshine-sys/mutiny/gateway"
	"github.com/starshine-sys/mutiny/session"
)

func main() {
	s, err := session.New(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("Error creating session: %v", err)
	}

	s.Handler.AddHandler(func(m *gateway.MessageEvent) {
		log.Printf("Received a message! Content: %v", m.String())

		if m.String() == "/ping" {
			_, err := s.SendMessage(m.ChannelID, fmt.Sprintf("Pong! %v", s.Ping().Round(time.Millisecond)))
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	})

	err = s.Open()
	if err != nil {
		log.Fatalf("Error opening session: %v", err)
	}

	select {}
}
```

## License

License information can be found in the `LICENSE` file in this repository.

Any code from other projects (mentioned in the specific files) is licensed under those projects' licenses.
