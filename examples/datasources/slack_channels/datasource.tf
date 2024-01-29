// Get all channels.
data "slack_channels" "all" {}

// Get a list of channels for a specific team.
data "slack_channels" "team" {
  team_id = "example"
}
