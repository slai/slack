package slack

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestAttachment_UnmarshalMarshalJSON_WithBlocks(t *testing.T) {

	originalAttachmentJson := `{
    "id": 1,
    "blocks": [
      {
        "type": "section",
        "block_id": "xxxx",
        "text": {
          "type": "mrkdwn",
          "text": "Pick something:",
          "verbatim": true
        },
        "accessory": {
          "type": "static_select",
          "action_id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
          "placeholder": {
            "type": "plain_text",
            "text": "Select one item",
            "emoji": true
          },
          "options": [
            {
              "text": {
                "type": "plain_text",
                "text": "ghi",
                "emoji": true
              },
              "value": "ghi"
            }
          ]
        }
      }
    ],
    "color": "#13A554",
    "fallback": "[no preview available]"
  }`

	attachment := new(Attachment)
	err := json.Unmarshal([]byte(originalAttachmentJson), attachment)
	if err != nil {
		t.Fatalf("expected no error unmarshaling attachment with blocks, got: %v", err)
	}

	actualAttachmentJson, err := json.Marshal(attachment)
	if err != nil {
		t.Fatal(err)
	}

	var (
		actual   interface{}
		expected interface{}
	)
	if err = json.Unmarshal([]byte(originalAttachmentJson), &expected); err != nil {
		t.Fatal(err)
	}
	if err = json.Unmarshal(actualAttachmentJson, &actual); err != nil {
		t.Fatal(err)
	}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Fatal("actual does not match expected\n", strings.Join(diff, "\n"))
	}
}

func TestAttachment_UnmarshalMarshalJSON_WithoutBlocks(t *testing.T) {
	originalAttachmentJson := `{
    "id": 1,
    "actions": [
      {
        "type": "button",
        "text": "Open",
        "name": "open"
      },
      {
        "type": "button",
        "text": "Close",
        "name": "close"
      }
    ],
    "text": "Available actions:"
  }`

	attachment := new(Attachment)
	err := json.Unmarshal([]byte(originalAttachmentJson), attachment)
	if err != nil {
		t.Fatalf("expected no error unmarshaling attachment without blocks, got: %v", err)
	}

	actualAttachmentJson, err := json.Marshal(attachment)
	if err != nil {
		t.Fatal(err)
	}

	var (
		actual   interface{}
		expected interface{}
	)
	if err = json.Unmarshal([]byte(originalAttachmentJson), &expected); err != nil {
		t.Fatal(err)
	}
	if err = json.Unmarshal(actualAttachmentJson, &actual); err != nil {
		t.Fatal(err)
	}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Fatal("actual does not match expected\n", strings.Join(diff, "\n"))
	}
}

func TestAttachment_UnmarshalMarshalJSON_WithoutBlocks_TextEmptyString(t *testing.T) {
	originalAttachmentJson := `{
    "id": 1,
    "actions": [
      {
        "type": "button",
        "text": "Open",
        "name": "open"
      },
      {
        "type": "button",
        "text": "Close",
        "name": "close"
      }
    ],
    "text": ""
  }`

	attachment := new(Attachment)
	err := json.Unmarshal([]byte(originalAttachmentJson), attachment)
	if err != nil {
		t.Fatalf("expected no error unmarshaling attachment without blocks, got: %v", err)
	}

	actualAttachmentJson, err := json.Marshal(attachment)
	if err != nil {
		t.Fatal(err)
	}

	var (
		actual   interface{}
		expected interface{}
	)
	if err = json.Unmarshal([]byte(originalAttachmentJson), &expected); err != nil {
		t.Fatal(err)
	}
	if err = json.Unmarshal(actualAttachmentJson, &actual); err != nil {
		t.Fatal(err)
	}
	if diff := deep.Equal(actual, expected); diff != nil {
		t.Fatal("actual does not match expected\n", strings.Join(diff, "\n"))
	}
}
