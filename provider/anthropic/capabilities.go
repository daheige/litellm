package anthropic

import "github.com/voocel/litellm"

func (p *Provider) Capabilities(model string) litellm.Capabilities {
	caps := litellm.Capabilities{
		Provider: p.Name(),
		Model:    model,
		Thinking: litellm.ThinkingCapabilities{
			Supported:     litellm.SupportYes,
			Disable:       litellm.SupportUnknown,
			Efforts:       []string{"low", "medium", "high"},
			BudgetTokens:  litellm.SupportNo,
			IncludeOutput: litellm.SupportYes,
			Notes:         []string{"adaptive thinking baseline; disable, xhigh, and max support are model-specific"},
		},
		Reasoning: litellm.ReasoningCapabilities{
			Blocks:          litellm.SupportYes,
			StreamingDeltas: litellm.SupportYes,
			ReasoningTokens: litellm.SupportNo,
		},
		Tools: litellm.ToolCapabilities{
			Calls:               litellm.SupportYes,
			StrictSchema:        litellm.SupportYes,
			Choice:              litellm.SupportPartial,
			MultimodalResults:   litellm.SupportYes,
			RoundTripSignatures: litellm.SupportYes,
		},
		Structured: litellm.StructuredCapabilities{
			JSONObject: litellm.SupportNo,
			JSONSchema: litellm.SupportUnknown,
			Strict:     litellm.SupportYes,
		},
		Media: litellm.MediaCapabilities{
			ImageURL:   litellm.SupportYes,
			ImageBytes: litellm.SupportYes,
			FileURI:    litellm.SupportNo,
		},
		Cache: litellm.CacheCapabilities{
			Block:      litellm.SupportYes,
			Retention:  litellm.SupportYes,
			UsageRead:  litellm.SupportYes,
			UsageWrite: litellm.SupportYes,
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
	return caps
}
