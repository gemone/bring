package bring

import (
	"math"
	"strconv"

	"github.com/gemone/bring/protocol"
)

const (
	defaultInt  int  = 256
	defaultByte byte = 255
)

// Handler func for  Guacamole instructions
type handlerFunc = func(client *Client, args []string) error

// Handlers for all instruction opcodes receivable by this Guacamole client.
var handlers = map[string]handlerFunc{
	"blob": func(c *Client, args []string) error {
		idx := parseInt(args[0])
		return c.streams.append(idx, args[1])
	},

	"copy": func(c *Client, args []string) error {
		srcL := parseInt(args[0])
		srcX := parseInt(args[1])
		srcY := parseInt(args[2])
		srcWidth := parseInt(args[3])
		srcHeight := parseInt(args[4])
		mask := parseByte(args[5])
		dstL := parseInt(args[6])
		dstX := parseInt(args[7])
		dstY := parseInt(args[8])
		c.display.copy(srcL, srcX, srcY, srcWidth, srcHeight,
			dstL, dstX, dstY, mask)
		return nil
	},

	"cfill": func(c *Client, args []string) error {
		layerIdx := parseInt(args[1])

		mask := parseByte(args[0])
		r := parseByte(args[2])
		g := parseByte(args[3])
		b := parseByte(args[4])
		a := parseByte(args[5])
		c.display.fill(layerIdx, r, g, b, a, mask)
		return nil
	},

	"cursor": func(c *Client, args []string) error {
		cursorHotspotX := parseInt(args[0])
		cursorHotspotY := parseInt(args[1])
		srcL := parseInt(args[2])
		srcX := parseInt(args[3])
		srcY := parseInt(args[4])
		srcWidth := parseInt(args[5])
		srcHeight := parseInt(args[6])
		c.display.setCursor(cursorHotspotX, cursorHotspotY,
			srcL, srcX, srcY, srcWidth, srcHeight)
		return nil
	},

	"disconnect": func(c *Client, args []string) error {
		c.session.Terminate()
		return nil
	},

	"dispose": func(c *Client, args []string) error {
		layerIdx := parseInt(args[0])
		c.display.dispose(layerIdx)
		return nil
	},

	"end": func(c *Client, args []string) error {
		idx := parseInt(args[0])
		c.streams.end(idx)
		c.streams.delete(idx)
		return nil
	},

	"error": func(c *Client, args []string) error {
		c.logger.Warnf("Received error from server: (%s) - %s", args[1], args[0])
		return nil
	},

	"img": func(c *Client, args []string) error {
		s := c.streams.get(parseInt(args[0]))
		op := parseByte(args[1])
		layerIdx := parseInt(args[2])
		//mimetype := args[3] // Not used
		x := parseInt(args[4])
		y := parseInt(args[5])
		s.onEnd = func(s *stream) {
			c.display.draw(layerIdx, x, y, op, s)
		}
		return nil
	},

	"log": func(c *Client, args []string) error {
		c.logger.Infof("Log from server:  %s", args[0])
		return nil
	},

	"rect": func(c *Client, args []string) error {
		layerIdx := parseInt(args[0])
		x := parseInt(args[1])
		y := parseInt(args[2])
		w := parseInt(args[3])
		h := parseInt(args[4])
		c.display.rect(layerIdx, x, y, w, h)
		return nil
	},

	"size": func(c *Client, args []string) error {
		layerIdx := parseInt(args[0])
		w := parseInt(args[1])
		h := parseInt(args[2])
		c.display.resize(layerIdx, w, h)
		return nil
	},

	"sync": func(c *Client, args []string) error {
		c.display.flush()
		if err := c.session.Send(protocol.NewInstruction("sync", args...)); err != nil {
			c.logger.Errorf("Failed to send 'sync' back to server: %s", err)
			return err
		}
		if c.onSync != nil {
			img, ts := c.display.getCanvas()
			c.onSync(img, ts)
		}
		return nil
	},
}

func parseInt(s string) int {
	// fix Incorrect conversion between integer types
	// https://github.com/gemone/bring/security/code-scanning/7
	// The number used by the protocol is not large, and it is converted to 32 bits
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return defaultInt
	}

	return int(n)
}

func parseByte(s string) byte {
	// byte uint8
	n, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return defaultByte
	}

	if n > 0 && n <= math.MaxUint8 {
		return byte(n)
	}

	return defaultByte
}
