package bedrock

import (
	"strings"

	"github.com/voocel/litellm"
)

func (p *Provider) Capabilities(model string) litellm.Capabilities {
	claude := strings.Contains(strings.ToLower(model), "claude")
	thinking := litellm.ThinkingCapabilities{
		Supported: litellm.SupportNo,
		Disable:   litellm.SupportNo,
	}
	if claude {
		thinking = litellm.ThinkingCapabilities{
			Supported:    litellm.SupportYes,
			Disable:      litellm.SupportUnknown,
			Efforts:      []string{"low", "medium", "high"},
			BudgetTokens: litellm.SupportNo,
			Notes:        []string{"adaptive thinking baseline; disable, xhigh, and max support are model-specific"},
		}
	}
	return litellm.Capabilities{
		Provider: p.Name(),
		Model:    model,
		Thinking: thinking,
		Reasoning: litellm.ReasoningCapabilities{
			Blocks:          litellm.SupportYes,
			StreamingDeltas: litellm.SupportYes,
			ReasoningTokens: litellm.SupportNo,
		},
		Tools: litellm.ToolCapabilities{
			Calls:               litellm.SupportYes,
			StrictSchema:        litellm.SupportPartial,
			Choice:              litellm.SupportYes,
			MultimodalResults:   litellm.SupportYes,
			RoundTripSignatures: litellm.SupportYes,
		},
		Structured: litellm.StructuredCapabilities{
			JSONObject: litellm.SupportNo,
			JSONSchema: litellm.SupportPartial,
			Strict:     litellm.SupportPartial,
		},
		Media: litellm.MediaCapabilities{
			ImageURL:   litellm.SupportNo,
			ImageBytes: litellm.SupportYes,
			FileURI:    litellm.SupportNo,
		},
		Cache: litellm.CacheCapabilities{
			Block:         litellm.SupportYes,
			RequestPolicy: litellm.SupportYes,
			Retention:     litellm.SupportYes,
			UsageRead:     litellm.SupportYes,
			UsageWrite:    litellm.SupportYes,
		},
		Streaming: litellm.StreamingCapabilities{
			Supported:       litellm.SupportYes,
			Usage:           litellm.SupportYes,
			ReasoningDeltas: litellm.SupportYes,
			ToolCallDeltas:  litellm.SupportYes,
			IdleTimeout:     litellm.SupportYes,
		},
		Usage: litellm.UsageCapabilities{
			InputTokens:      litellm.SupportYes,
			OutputTokens:     litellm.SupportYes,
			TotalTokens:      litellm.SupportYes,
			ReasoningTokens:  litellm.SupportNo,
			CacheReadTokens:  litellm.SupportYes,
			CacheWriteTokens: litellm.SupportYes,
		},
	}
}
