// Code generated from the Seolyticize tool registry. DO NOT EDIT.
// Regenerate with `pnpm --filter @seolyticize/sdk generate:languages`.

package seolyticize

import (
	"context"
	"encoding/json"
)

// ToolCount is how many tools this SDK was generated with.
const ToolCount = 199

// readOnlyTools are retried on server errors; every other call is not, since
// a write that failed with a 500 may already have happened.
var readOnlyTools = map[string]bool{
	"analyse_server_log":             true,
	"batch_domain_metrics":           true,
	"build_disavow_file":             true,
	"estimate_ai_visibility_cost":    true,
	"estimate_cost":                  true,
	"explain_keyword_difficulty":     true,
	"export_ai_visibility_data":      true,
	"export_search_console_data":     true,
	"get_agent_readiness":            true,
	"get_ai_visibility":              true,
	"get_ai_visibility_answer":       true,
	"get_ai_visibility_results":      true,
	"get_ai_visibility_runs":         true,
	"get_ai_visibility_setup":        true,
	"get_ai_visibility_trend":        true,
	"get_audit_summary":              true,
	"get_backlink_audit":             true,
	"get_billing":                    true,
	"get_bing_performance":           true,
	"get_branding":                   true,
	"get_cannibalization":            true,
	"get_click_curve":                true,
	"get_client":                     true,
	"get_client_brief":               true,
	"get_client_costs":               true,
	"get_client_integrations":        true,
	"get_connection_link":            true,
	"get_content_decay":              true,
	"get_content_gap":                true,
	"get_costs":                      true,
	"get_crawl_status":               true,
	"get_deploy_hook":                true,
	"get_domain_facts":               true,
	"get_domain_technologies":        true,
	"get_entity_pageviews":           true,
	"get_ga4_audience":               true,
	"get_ga4_ecommerce":              true,
	"get_ga4_key_events":             true,
	"get_ga4_landing_pages":          true,
	"get_ga4_measurement_health":     true,
	"get_ga4_overview":               true,
	"get_ga4_page_performance":       true,
	"get_ga4_site_search":            true,
	"get_ga4_traffic_acquisition":    true,
	"get_google_service_account":     true,
	"get_index_coverage":             true,
	"get_indexnow_key":               true,
	"get_keyword_difficulty":         true,
	"get_keyword_gap":                true,
	"get_keyword_list":               true,
	"get_link_checks":                true,
	"get_link_intersect":             true,
	"get_local_visibility":           true,
	"get_low_ctr_queries":            true,
	"get_next_steps":                 true,
	"get_opportunities":              true,
	"get_page_issues":                true,
	"get_panel_reach":                true,
	"get_position_distribution":      true,
	"get_rank_history":               true,
	"get_rank_summary":               true,
	"get_run_status":                 true,
	"get_search_console_performance": true,
	"get_search_demand":              true,
	"get_search_opportunities":       true,
	"get_serp_features":              true,
	"get_share_of_voice":             true,
	"get_site_status":                true,
	"get_top_searches":               true,
	"get_url_swaps":                  true,
	"get_web_analytics_breakdown":    true,
	"get_web_analytics_setup":        true,
	"get_web_analytics_summary":      true,
	"get_written_report":             true,
	"get_yandex_performance":         true,
	"list_alert_rules":               true,
	"list_clients":                   true,
	"list_crawl_issues":              true,
	"list_crawls":                    true,
	"list_issue_pages":               true,
	"list_keyword_lists":             true,
	"list_link_prospects":            true,
	"list_local_locations":           true,
	"list_markets":                   true,
	"list_projects":                  true,
	"list_rank_trackers":             true,
	"list_recent_alerts":             true,
	"list_report_schedules":          true,
	"list_report_templates":          true,
	"list_reports":                   true,
	"list_sitemaps":                  true,
	"list_tracked_keywords":          true,
	"list_written_reports":           true,
	"page_vitals":                    true,
	"striking_distance":              true,
	"whoami":                         true,
}

// AddCompetitorsInput is the input to add_competitors.
type AddCompetitorsInput struct {
	ClientID  string   `json:"clientId"`
	TrackerID string   `json:"trackerId"`
	Domains   []string `json:"domains"` // Domains such as example.com. A URL is reduced to its domain.
}

// AddCompetitors calls add_competitors: Add competitors to a tracker. Adds
// competitor domains to a rank tracker, so share of voice and the competitor
// comparisons measure them on the same keywords. Answers 'compare us with
// these sites'. Returns the tracker's full competitor list. Domains already
// there are left as they are.
func (c *Client) AddCompetitors(ctx context.Context, in AddCompetitorsInput) (json.RawMessage, error) {
	return c.Call(ctx, "add_competitors", in)
}

// AddKeywordsToListInput is the input to add_keywords_to_list.
type AddKeywordsToListInput struct {
	ListID   string           `json:"listId"`
	Keywords []map[string]any `json:"keywords"`
}

// AddKeywordsToList calls add_keywords_to_list: Add keywords to a list. Adds
// keywords to an existing list. Keywords already on it are left alone and
// counted as already present. Answers 'save these to that list'. Location code
// defaults to 2840 (United States), language to en, engine to google.
func (c *Client) AddKeywordsToList(ctx context.Context, in AddKeywordsToListInput) (json.RawMessage, error) {
	return c.Call(ctx, "add_keywords_to_list", in)
}

// AddLinkProspectsInput is the input to add_link_prospects.
type AddLinkProspectsInput struct {
	ClientID  string           `json:"clientId"`
	Prospects []map[string]any `json:"prospects"`
}

// AddLinkProspects calls add_link_prospects: Add sites to a client's outreach
// list. Adds linking-site prospects to the client's outreach list at the
// identified stage — typically the domains get_link_intersect finds linking to
// competitors but not the client, with that as the reason. A site already on
// the list is left at its current stage, never reset. Answers 'add these sites
// to our outreach'. Free.
func (c *Client) AddLinkProspects(ctx context.Context, in AddLinkProspectsInput) (json.RawMessage, error) {
	return c.Call(ctx, "add_link_prospects", in)
}

// AddListToTrackerInput is the input to add_list_to_tracker.
type AddListToTrackerInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId"`
	ListID    string `json:"listId"`
}

// AddListToTracker calls add_list_to_tracker: Track a keyword list. Adds every
// keyword on a list to one of the client's rank trackers. Answers 'start
// tracking everything on that list'. Keywords for another country, language or
// engine than the tracker reads are skipped and listed with the reason, never
// tracked in the wrong market. Each added keyword is a paid lookup on every
// future run, so quote the tracker with estimate_cost afterwards if cost
// matters; refused when the plan's keyword allowance would be exceeded.
func (c *Client) AddListToTracker(ctx context.Context, in AddListToTrackerInput) (json.RawMessage, error) {
	return c.Call(ctx, "add_list_to_tracker", in)
}

// AddLocalLocationInput is the input to add_local_location.
type AddLocalLocationInput struct {
	ClientID string  `json:"clientId"`
	Name     string  `json:"name"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
}

// AddLocalLocation calls add_local_location: Add a business location. Adds a
// business location to a client for local SEO, at a latitude and longitude —
// the centre of every grid scan around it. Answers 'start tracking this branch
// in the map pack'. Free.
func (c *Client) AddLocalLocation(ctx context.Context, in AddLocalLocationInput) (json.RawMessage, error) {
	return c.Call(ctx, "add_local_location", in)
}

// AddTeamMemberInput is the input to add_team_member.
type AddTeamMemberInput struct {
	Email     string   `json:"email"`
	Role      string   `json:"role"`
	ClientIDs []string `json:"clientIds,omitempty"`
}

// AddTeamMember calls add_team_member: Add someone to the team. Gives an
// existing Seolyticize account access to the agency with a role — admin,
// manager, analyst or client_viewer — optionally limited to some clients.
// Answers 'give Sam access as an analyst'. The person must already have signed
// up; this never creates an account for them. Refused for someone already on
// the team, and for a client outside the caller's own reach. Owners can only
// be added by an owner, in the app.
func (c *Client) AddTeamMember(ctx context.Context, in AddTeamMemberInput) (json.RawMessage, error) {
	return c.Call(ctx, "add_team_member", in)
}

// AddTrackedKeywordsInput is the input to add_tracked_keywords.
type AddTrackedKeywordsInput struct {
	ClientID  string   `json:"clientId"`
	TrackerID string   `json:"trackerId"`
	Terms     []string `json:"terms"`
}

// AddTrackedKeywords calls add_tracked_keywords: Track keywords. Adds terms to
// a tracker, ignoring ones already tracked. Answers 'start tracking these
// keywords'. Each added term is a paid lookup on every future run; refused
// when the plan's keyword allowance would be exceeded. Safe to repeat.
func (c *Client) AddTrackedKeywords(ctx context.Context, in AddTrackedKeywordsInput) (json.RawMessage, error) {
	return c.Call(ctx, "add_tracked_keywords", in)
}

// AnalyseServerLogInput is the input to analyse_server_log.
type AnalyseServerLogInput struct {
	ClientID string `json:"clientId"`
	Contents string `json:"contents"`
}

// AnalyseServerLog calls analyse_server_log: Analyse a server access log.
// Reads a web server access log in combined format (nginx's and Apache's
// default) passed as text, and reports what search and AI crawlers spent their
// visits on: analysis.byBot (hits per crawler), topPaths, errorPaths (path,
// status, hits), wastedOnErrors and wastedOnRedirects (shares of crawler
// requests that returned no content), aiCrawlerShare, totalHits and
// totalBytes. Also linesRead, linesUnparsed (lines in a format it did not
// recognise — many means the wrong format, not a quiet site), humanHits and
// truncated. Up to 64 MB per call; split a bigger log and call once per part.
// Nothing is stored. Answers 'what is Googlebot actually crawling on our
// site'. Free.
func (c *Client) AnalyseServerLog(ctx context.Context, in AnalyseServerLogInput) (json.RawMessage, error) {
	return c.Call(ctx, "analyse_server_log", in)
}

// ArchiveClientInput is the input to archive_client.
type ArchiveClientInput struct {
	ClientID string `json:"clientId"`
}

// ArchiveClient calls archive_client: Archive a client. Retires a client: it
// stops being tracked, crawled and reported on, stops counting against the
// plan, and disappears from lists. History is kept. Answers 'we no longer work
// with this client'. Safe to repeat.
func (c *Client) ArchiveClient(ctx context.Context, in ArchiveClientInput) (json.RawMessage, error) {
	return c.Call(ctx, "archive_client", in)
}

// ArchiveReportInput is the input to archive_report.
type ArchiveReportInput struct {
	ClientID string `json:"clientId"`
	ReportID string `json:"reportId"`
}

// ArchiveReport calls archive_report: Retire a saved report. Retires a
// client's saved report so it is no longer sent, by schedule or by hand.
// Answers 'we no longer send this report'. Archived rather than deleted:
// reports already sent to the client stay downloadable. Safe to repeat.
func (c *Client) ArchiveReport(ctx context.Context, in ArchiveReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "archive_report", in)
}

// BatchDomainMetricsInput is the input to batch_domain_metrics.
type BatchDomainMetricsInput struct {
	Domains []string `json:"domains"`
}

// BatchDomainMetrics calls batch_domain_metrics: Look up domains in bulk.
// Authority, keywords, traffic and referring domains for any list of domains,
// managed or not. Answers 'how strong is this site' for competitors and
// prospects.
func (c *Client) BatchDomainMetrics(ctx context.Context, in BatchDomainMetricsInput) (json.RawMessage, error) {
	return c.Call(ctx, "batch_domain_metrics", in)
}

// BuildDisavowFileInput is the input to build_disavow_file.
type BuildDisavowFileInput struct {
	Entries []map[string]any `json:"entries"`
}

// BuildDisavowFile calls build_disavow_file: Build a disavow file. Builds a
// Google disavow file from domains and URLs: a domain entry disavows every
// link from that domain, a url entry names one page. Answers 'make the disavow
// file for these links'. Free; nothing is submitted — Google takes the file by
// upload in Search Console, which a person does. Use get_backlink_audit to
// find the links worth disavowing, and have a person review every line first:
// disavowing a good link removes value the client earned. Malformed entries
// are left out rather than written.
func (c *Client) BuildDisavowFile(ctx context.Context, in BuildDisavowFileInput) (json.RawMessage, error) {
	return c.Call(ctx, "build_disavow_file", in)
}

// ConnectGoogleAnalyticsServiceAccountInput is the input to connect_google_analytics_service_account.
type ConnectGoogleAnalyticsServiceAccountInput struct {
	ClientID   string `json:"clientId"`
	PropertyID string `json:"propertyId"`
}

// ConnectGoogleAnalyticsServiceAccount calls
// connect_google_analytics_service_account: Connect GA4 by service account.
// Connects a client's GA4 property as the agency's Google service account,
// once the property's owner has added that address with the Viewer role.
// Answers 'they added our address in GA4, connect it'. propertyId is GA4's
// numeric property id (Admin → Property details), not a G- measurement id.
// Access is checked with Google first, and refused with the exact steps when
// it is not granted yet. Free.
func (c *Client) ConnectGoogleAnalyticsServiceAccount(ctx context.Context, in ConnectGoogleAnalyticsServiceAccountInput) (json.RawMessage, error) {
	return c.Call(ctx, "connect_google_analytics_service_account", in)
}

// ConnectSearchConsoleServiceAccountInput is the input to connect_search_console_service_account.
type ConnectSearchConsoleServiceAccountInput struct {
	ClientID    string `json:"clientId"`
	PropertyURL string `json:"propertyUrl,omitempty"`
}

// ConnectSearchConsoleServiceAccount calls
// connect_search_console_service_account: Connect Search Console by service
// account. Connects a client's Search Console property as the agency's Google
// service account (get_google_service_account), once the property's owner has
// added that address as a user. Answers 'they added our address, connect
// Search Console'. Access is checked with Google first: with no propertyUrl,
// the client's site is tried as a Domain property (sc-domain:) and then as
// https:// and https://www. URL-prefix properties, and the first the account
// can read is connected. If none can be read yet, it is refused with the exact
// steps for the owner. Starts the sixteen-month backfill. Free.
func (c *Client) ConnectSearchConsoleServiceAccount(ctx context.Context, in ConnectSearchConsoleServiceAccountInput) (json.RawMessage, error) {
	return c.Call(ctx, "connect_search_console_service_account", in)
}

// ConnectYandexWithTokenInput is the input to connect_yandex_with_token.
type ConnectYandexWithTokenInput struct {
	ClientID string `json:"clientId"`
	Token    string `json:"token"`
}

// ConnectYandexWithToken calls connect_yandex_with_token: Connect Yandex
// Webmaster with a token. Connects a client's Yandex Webmaster with an OAuth
// token a person made in their own Yandex app (get_connection_link with
// yandex_webmaster returns the steps), instead of signing in through
// Seolyticize. Answers 'here is the Yandex token, connect it'. Yandex is asked
// whose token it is and the client's site is matched against that account's
// verified hosts; the response says which host matched, or that none did.
// Works even where this deployment has no Yandex sign-in set up. The token is
// stored sealed and never returned. Free.
func (c *Client) ConnectYandexWithToken(ctx context.Context, in ConnectYandexWithTokenInput) (json.RawMessage, error) {
	return c.Call(ctx, "connect_yandex_with_token", in)
}

// CreateAlertRuleInput is the input to create_alert_rule.
type CreateAlertRuleInput struct {
	ClientID        string   `json:"clientId"`
	Kind            string   `json:"kind"`
	Threshold       int64    `json:"threshold"`
	CooldownHours   *int64   `json:"cooldownHours,omitempty"`
	WebhookURL      string   `json:"webhookUrl,omitempty"`
	Channel         string   `json:"channel,omitempty"`
	EmailRecipients []string `json:"emailRecipients,omitempty"`
	TelegramChatID  string   `json:"telegramChatId,omitempty"`
}

// CreateAlertRule calls create_alert_rule: Create an alert rule. Raises an
// alert when a client's rankings, traffic, issues or backlinks cross a
// threshold, or the site goes down, and delivers it by channel: slack
// (webhookUrl, Slack's format; the default), webhook (webhookUrl, plain JSON),
// email (emailRecipients, up to ten) or telegram (telegramChatId, with the
// agency's bot from Settings). Webhooks must be https and public. Answers
// 'tell me when rankings drop'.
func (c *Client) CreateAlertRule(ctx context.Context, in CreateAlertRuleInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_alert_rule", in)
}

// CreateAPIKeyInput is the input to create_api_key.
type CreateAPIKeyInput struct {
	Name              string   `json:"name"`
	ClientID          string   `json:"clientId,omitempty"`
	Role              string   `json:"role,omitempty"`
	ExpiresInDays     *int64   `json:"expiresInDays,omitempty"`
	MaxCostUsdPerCall *float64 `json:"maxCostUsdPerCall,omitempty"` // The most any one call with this key may spend, in USD. Calls asking for more are capped to it.
}

// CreateAPIKey calls create_api_key: Create an API key. Mints an API key,
// agency-wide or pinned to one client, with a role: client_viewer (read-only),
// analyst, manager, site_admin or admin. Answers 'give this integration its
// own key'. For an agent that runs one site, pin it to that client with
// site_admin: full control of that client — connecting consoles, crawls,
// onboarding — and none of the agency's settings, team, keys or billing; the
// key cannot even see other clients. A client-pinned key is at most
// site_admin, and site_admin needs a clientId. expiresInDays makes it stop
// working on its own; maxCostUsdPerCall caps what any one call with it may
// spend. The key is in the response once and can never be shown again. A key
// cannot be given more access than the caller has.
func (c *Client) CreateAPIKey(ctx context.Context, in CreateAPIKeyInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_api_key", in)
}

// CreateClientInput is the input to create_client.
type CreateClientInput struct {
	Name             string   `json:"name"`
	Domain           string   `json:"domain"`
	Keywords         []string `json:"keywords,omitempty"`
	MonthlyBudgetUsd *float64 `json:"monthlyBudgetUsd,omitempty"`
}

// CreateClient calls create_client: Create a client. Onboards a client in one
// step: the client, a project, its site and a rank tracker, optionally with
// keywords to track. Answers 'set up a new client'. Refused when the name is
// taken or the plan's client or keyword allowance is used up; every tracked
// keyword is a paid lookup on every run.
func (c *Client) CreateClient(ctx context.Context, in CreateClientInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_client", in)
}

// CreateDeployHookInput is the input to create_deploy_hook.
type CreateDeployHookInput struct {
	ClientID   string `json:"clientId"`
	SitemapURL string `json:"sitemapUrl,omitempty"` // Read this sitemap instead of robots.txt's Sitemap: line or /sitemap.xml.
}

// CreateDeployHook calls create_deploy_hook: Create a deploy hook. Answers
// 'tell search engines whenever we deploy'. Makes the client's deploy hook: a
// secret URL the site's pipeline calls after every deploy (one curl, or a
// Netlify/Vercel/Cloudflare Pages/Coolify notification). Each call reads the
// sitemap, compares it with the last one, and submits the new and changed
// pages to IndexNow; the first call records the baseline only. Returns the URL
// once, with how to call it from each platform. Making it again replaces the
// URL. Free.
func (c *Client) CreateDeployHook(ctx context.Context, in CreateDeployHookInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_deploy_hook", in)
}

// CreateInviteCodeInput is the input to create_invite_code.
type CreateInviteCodeInput struct {
	Label                string   `json:"label"`
	MaxUses              *int64   `json:"maxUses,omitempty"`
	ExpiresInDays        *int64   `json:"expiresInDays,omitempty"`
	KeyExpiresInDays     *int64   `json:"keyExpiresInDays,omitempty"`
	KeyMaxCostUsdPerCall *float64 `json:"keyMaxCostUsdPerCall,omitempty"`
}

// CreateInviteCode calls create_invite_code: Create an invite code for a
// site's agent. Answers 'let an agent connect its own site without our admin
// key'. Makes a code to hand to the agent that runs one site. The agent sends
// it, with its domain, to POST /api/v1/claim (no key needed) and gets back a
// key pinned to the one new client for that site, with the site_admin role:
// full control of that client, nothing else in the agency. The code is in this
// response once. maxUses (default 1) is how many sites it can claim;
// expiresInDays ends the code, keyExpiresInDays the keys it makes, and
// keyMaxCostUsdPerCall caps what any one call with those keys may spend. Free.
func (c *Client) CreateInviteCode(ctx context.Context, in CreateInviteCodeInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_invite_code", in)
}

// CreateKeywordListInput is the input to create_keyword_list.
type CreateKeywordListInput struct {
	Name     string           `json:"name"`
	ClientID string           `json:"clientId,omitempty"`
	Keywords []map[string]any `json:"keywords,omitempty"`
}

// CreateKeywordList calls create_keyword_list: Create a keyword list. Creates
// a named keyword list, for one client or for the agency, optionally with its
// first keywords. Answers 'keep these keywords somewhere'. A keyword is its
// words plus a location code, language and engine; the same words for another
// country are a different keyword. Returns created: false with the reason when
// the name is already taken, rather than failing.
func (c *Client) CreateKeywordList(ctx context.Context, in CreateKeywordListInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_keyword_list", in)
}

// CreatePortalLinkInput is the input to create_portal_link.
type CreatePortalLinkInput struct {
	ClientID string `json:"clientId"`
	Label    string `json:"label"` // Who the link is for, so it can be told apart and revoked later
}

// CreatePortalLink calls create_portal_link: Create a client portal link.
// Makes a white-label link that shows one client their own dashboard under the
// agency's name, with no sign-in. Answers 'give the client a link to their
// results'. The full address is in this response once and can never be shown
// again; send it to the client and keep it out of logs. Anyone holding it can
// see that client's results until it is revoked.
func (c *Client) CreatePortalLink(ctx context.Context, in CreatePortalLinkInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_portal_link", in)
}

// CreateProjectInput is the input to create_project.
type CreateProjectInput struct {
	ClientID     string `json:"clientId"`
	Name         string `json:"name"`
	Domain       string `json:"domain,omitempty"`
	LocationCode *int64 `json:"locationCode,omitempty"`
	LanguageCode string `json:"languageCode,omitempty"`
}

// CreateProject calls create_project: Create a project. Adds a project to a
// client, optionally with a site domain to crawl. Answers 'add a second site
// or market for this client'. The domain is written like example.com.
func (c *Client) CreateProject(ctx context.Context, in CreateProjectInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_project", in)
}

// CreateRankTrackerInput is the input to create_rank_tracker.
type CreateRankTrackerInput struct {
	ClientID     string `json:"clientId"`
	Name         string `json:"name"`
	ProjectID    string `json:"projectId,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Engine       string `json:"engine,omitempty"`
	Device       string `json:"device,omitempty"`
	Frequency    string `json:"frequency,omitempty"`
	LocationCode *int64 `json:"locationCode,omitempty"`
	LanguageCode string `json:"languageCode,omitempty"`
}

// CreateRankTracker calls create_rank_tracker: Create a rank tracker. Starts
// tracking a client's positions on one engine, device and locale. Answers
// 'track this client on Bing too' or 'track mobile separately'. Defaults to
// the client's first project and site domain. Engine and locale cannot be
// changed later — they are part of every keyword's identity — so a different
// market is a new tracker.
func (c *Client) CreateRankTracker(ctx context.Context, in CreateRankTrackerInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_rank_tracker", in)
}

// CreateReportInput is the input to create_report.
type CreateReportInput struct {
	ClientID string   `json:"clientId"`
	Title    string   `json:"title"`
	Sections []string `json:"sections"`
}

// CreateReport calls create_report: Create a report. Saves a report for a
// client from a list of sections; each is wired to the client's own tracker
// and site. Answers 'build a monthly report for this client'. Sections the
// client has nothing for yet are left out. Send it with send_report; schedules
// are set in the app.
func (c *Client) CreateReport(ctx context.Context, in CreateReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_report", in)
}

// DeleteAlertRuleInput is the input to delete_alert_rule.
type DeleteAlertRuleInput struct {
	ClientID string `json:"clientId"`
	RuleID   string `json:"ruleId"`
}

// DeleteAlertRule calls delete_alert_rule: Delete an alert rule. Deletes a
// rule and every alert it ever fired. Answers 'remove this alert for good'. To
// stop it but keep its history, pause it with update_alert_rule instead. Safe
// to repeat.
func (c *Client) DeleteAlertRule(ctx context.Context, in DeleteAlertRuleInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_alert_rule", in)
}

// DeleteKeywordListInput is the input to delete_keyword_list.
type DeleteKeywordListInput struct {
	ListID string `json:"listId"`
}

// DeleteKeywordList calls delete_keyword_list: Delete a keyword list. Deletes
// a keyword list and the keywords on it, for good. Answers 'get rid of that
// list'. Keywords a tracker took from the list stay tracked.
func (c *Client) DeleteKeywordList(ctx context.Context, in DeleteKeywordListInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_keyword_list", in)
}

// DeleteRankTrackerInput is the input to delete_rank_tracker.
type DeleteRankTrackerInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId"`
}

// DeleteRankTracker calls delete_rank_tracker: Delete a rank tracker. Stops a
// tracker. Positions it already captured stay in reports; no further runs are
// made or charged. Answers 'stop tracking this'. Safe to repeat.
func (c *Client) DeleteRankTracker(ctx context.Context, in DeleteRankTrackerInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_rank_tracker", in)
}

// DeleteReportScheduleInput is the input to delete_report_schedule.
type DeleteReportScheduleInput struct {
	ClientID   string `json:"clientId"`
	ScheduleID string `json:"scheduleId"`
}

// DeleteReportSchedule calls delete_report_schedule: Delete a report schedule.
// Removes a schedule for good: its recipients stop receiving that report.
// Answers 'stop sending this report'. The report itself and its past sends are
// kept. To stop it for a while instead, pause it with update_report_schedule.
func (c *Client) DeleteReportSchedule(ctx context.Context, in DeleteReportScheduleInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_report_schedule", in)
}

// DeleteReportTemplateInput is the input to delete_report_template.
type DeleteReportTemplateInput struct {
	TemplateID string `json:"templateId"`
	ClientID   string `json:"clientId,omitempty"`
}

// DeleteReportTemplate calls delete_report_template: Delete a report template.
// Permanently deletes one report template. Answers 'we no longer use that
// format'. Reports written to it are kept. Call list_report_templates to find
// its id. Free.
func (c *Client) DeleteReportTemplate(ctx context.Context, in DeleteReportTemplateInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_report_template", in)
}

// DeleteSitemapInput is the input to delete_sitemap.
type DeleteSitemapInput struct {
	ClientID   string `json:"clientId"`
	SitemapURL string `json:"sitemapUrl"`
}

// DeleteSitemap calls delete_sitemap: Remove a sitemap from Google. Removes a
// sitemap from the client's Google Search Console property. Answers 'stop
// Google reading this old sitemap'. Same refusals as submit_sitemap: only
// addresses on a site the agency manages, and only once sitemap management has
// been allowed.
func (c *Client) DeleteSitemap(ctx context.Context, in DeleteSitemapInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_sitemap", in)
}

// DeleteWrittenReportInput is the input to delete_written_report.
type DeleteWrittenReportInput struct {
	ClientID string `json:"clientId"`
	ReportID string `json:"reportId"`
}

// DeleteWrittenReport calls delete_written_report: Delete a written report.
// Permanently deletes one written report; its share link stops working at
// once. Answers 'remove that report'. Does not touch data reports, templates
// or anything else. Free.
func (c *Client) DeleteWrittenReport(ctx context.Context, in DeleteWrittenReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_written_report", in)
}

// DisconnectAppInput is the input to disconnect_app.
type DisconnectAppInput struct {
	GrantID string `json:"grantId"`
}

// DisconnectApp calls disconnect_app: Disconnect an app. Stops a signed-in app
// working, at once: its access and refresh tokens both end. Answers 'remove
// that assistant's access'. Takes the grantId from list_connected_apps. The
// app can only come back if somebody signs in and approves it again. Safe to
// repeat.
func (c *Client) DisconnectApp(ctx context.Context, in DisconnectAppInput) (json.RawMessage, error) {
	return c.Call(ctx, "disconnect_app", in)
}

// DraftClientBriefInput is the input to draft_client_brief.
type DraftClientBriefInput struct {
	ClientID      string `json:"clientId"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// DraftClientBrief calls draft_client_brief: Draft the client brief from their
// website. Reads the client's home page and up to four of its about, services,
// product or pricing pages, and drafts the brief's overview, positioning,
// audience and voice, plus the key pages with what each is for. Answers 'set
// up this client's brief'. Nothing is saved: show the draft to the user,
// correct it, then save what is right with update_client_brief —
// alreadyWritten lists fields the brief already has, which should not be
// overwritten without asking. Says 'Not stated on the site.' where the pages
// are silent, and never names competitors. A paid Claude call charged to the
// client's budget (costMicros reports what it cost).
func (c *Client) DraftClientBrief(ctx context.Context, in DraftClientBriefInput) (json.RawMessage, error) {
	return c.Call(ctx, "draft_client_brief", in)
}

// EstimateAIVisibilityCostInput is the input to estimate_ai_visibility_cost.
type EstimateAIVisibilityCostInput struct {
	ClientID string `json:"clientId"`
}

// EstimateAIVisibilityCost calls estimate_ai_visibility_cost: What an AI
// visibility check costs. Estimates one check and a month at the current
// cadence: questions × assistants answers per run, and the most the Claude
// answers can cost (Seolyticize meters Claude; other assistants bill the
// agency's own accounts at their own prices, so they are counted, not priced).
// Answers 'how much will tracking this cost'. Free.
func (c *Client) EstimateAIVisibilityCost(ctx context.Context, in EstimateAIVisibilityCostInput) (json.RawMessage, error) {
	return c.Call(ctx, "estimate_ai_visibility_cost", in)
}

// EstimateCostInput is the input to estimate_cost.
type EstimateCostInput struct {
	Tool      string         `json:"tool"`
	Arguments map[string]any `json:"arguments,omitempty"`
}

// EstimateCost calls estimate_cost: Estimate what a call will cost. Quotes the
// most a tool call can cost before it is made, without buying anything.
// Answers 'how much will this cost'. Pass the tool's name and the exact
// arguments you would call it with; the result lists each purchase the call
// can make with its ceiling, plus monthlyMicros for scheduled work. Show the
// estimate to the user before a paid call, then pass estimatedMaxMicros as
// that call's maxCostMicros so it cannot spend more. Cached data costs
// nothing, so real charges are often lower. Also quotes "rank_tracker": a rank
// tracker's captures (arguments: clientId, trackerId) — per run and per month.
// Free.
func (c *Client) EstimateCost(ctx context.Context, in EstimateCostInput) (json.RawMessage, error) {
	return c.Call(ctx, "estimate_cost", in)
}

// EstimateDomainTrafficInput is the input to estimate_domain_traffic.
type EstimateDomainTrafficInput struct {
	ClientID      string   `json:"clientId"` // The client whose account and monthly cap pay for the lookup
	Targets       []string `json:"targets"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// EstimateDomainTraffic calls estimate_domain_traffic: Estimate search traffic
// for domains. A vendor's current estimate of monthly organic search traffic
// and ranking keywords for up to 100 domains, bought now. Answers 'how much
// search traffic do these competitors get'. Modelled from rankings, not
// measured visits. batch_domain_metrics reads what is already stored for free;
// use this when those figures are missing or stale. A paid lookup on the
// client's account: quote it with estimate_cost.
func (c *Client) EstimateDomainTraffic(ctx context.Context, in EstimateDomainTrafficInput) (json.RawMessage, error) {
	return c.Call(ctx, "estimate_domain_traffic", in)
}

// ExplainKeywordDifficultyInput is the input to explain_keyword_difficulty.
type ExplainKeywordDifficultyInput struct {
	Term         string `json:"term"`
	LocationCode *int64 `json:"locationCode,omitempty"`
	LanguageCode string `json:"languageCode,omitempty"`
	KeywordID    string `json:"keywordId,omitempty"`
}

// ExplainKeywordDifficulty calls explain_keyword_difficulty: Why a keyword is
// as hard as it is. Which first-page results set a keyword's difficulty score
// and which make it harder than the score suggests: each result's referring
// domains and role (sets the median, above it, below it, or unscored for lack
// of link data), the strongest and weakest result, the domain authority beside
// each, and a one-paragraph summary. Answers 'why is this keyword a 47' so the
// number can be checked rather than trusted. Read from the archive and our
// link graph; buys nothing, and a keyword never captured is reported unscored.
func (c *Client) ExplainKeywordDifficulty(ctx context.Context, in ExplainKeywordDifficultyInput) (json.RawMessage, error) {
	return c.Call(ctx, "explain_keyword_difficulty", in)
}

// ExploreAIPromptInput is the input to explore_ai_prompt.
type ExploreAIPromptInput struct {
	ClientID      string   `json:"clientId,omitempty"`
	Prompt        string   `json:"prompt"`
	Engines       []string `json:"engines,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// ExploreAIPrompt calls explore_ai_prompt: Ask assistants one question now.
// Asks one question live on the chosen assistants (claude, and any the agency
// has its own API key for) and returns each answer's text, citations, and —
// when clientId is given — which tracked brands it named. Nothing is stored;
// use it to see what assistants say before tracking a question, or to look up
// how they describe a brand. Answers 'what does ChatGPT say about this'.
// Claude is metered (meteredMicros); others bill the agency's own accounts.
func (c *Client) ExploreAIPrompt(ctx context.Context, in ExploreAIPromptInput) (json.RawMessage, error) {
	return c.Call(ctx, "explore_ai_prompt", in)
}

// ExportAIVisibilityDataInput is the input to export_ai_visibility_data.
type ExportAIVisibilityDataInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
	Include  string `json:"include,omitempty"`
}

// ExportAIVisibilityData calls export_ai_visibility_data: Export AI visibility
// data. CSV of the last N days (default 90): include "mentions" for one row
// per answer and brand (named, position, cited), or "answers" for the full
// answer text and citations (up to 2,000 answers). Returns rows and csv.
// Answers 'give me the raw data for a deck'. Free.
func (c *Client) ExportAIVisibilityData(ctx context.Context, in ExportAIVisibilityDataInput) (json.RawMessage, error) {
	return c.Call(ctx, "export_ai_visibility_data", in)
}

// ExportSearchConsoleDataInput is the input to export_search_console_data.
type ExportSearchConsoleDataInput struct {
	ClientID   string `json:"clientId"`
	SearchType string `json:"searchType,omitempty"`
	From       string `json:"from"`
	To         string `json:"to"`
	Country    string `json:"country,omitempty"`
	Device     string `json:"device,omitempty"`
	Dimension  string `json:"dimension,omitempty"`
}

// ExportSearchConsoleData calls export_search_console_data: Export Search
// Console queries or pages as CSV. Google Search Console queries or pages
// (dimension) for a search type, date range, and optionally one country and
// device, as CSV with clicks, impressions, click-through rate and average
// position — up to 5,000 rows. Returns filename, rows and csv. Cells a
// spreadsheet would run as formulas are neutralised. Answers 'give me the
// query data in a spreadsheet'. Free.
func (c *Client) ExportSearchConsoleData(ctx context.Context, in ExportSearchConsoleDataInput) (json.RawMessage, error) {
	return c.Call(ctx, "export_search_console_data", in)
}

// FindSerpCompetitorsInput is the input to find_serp_competitors.
type FindSerpCompetitorsInput struct {
	ClientID      string   `json:"clientId"`
	Keywords      []string `json:"keywords"`
	LocationCode  *int64   `json:"locationCode,omitempty"`
	LanguageCode  string   `json:"languageCode,omitempty"`
	Limit         *int64   `json:"limit,omitempty"`
	Exclude       []string `json:"exclude,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// FindSerpCompetitors calls find_serp_competitors: Who competes across a
// keyword set. Finds the domains that rank across a set of keywords (up to
// 200) in one market: visibility (share of the keywords they hold in the top
// ten), keyword count, average position and estimated traffic. Answers 'who
// are our real search competitors for these terms'. List the client's own
// domain in exclude. One purchase per call.
func (c *Client) FindSerpCompetitors(ctx context.Context, in FindSerpCompetitorsInput) (json.RawMessage, error) {
	return c.Call(ctx, "find_serp_competitors", in)
}

// GetAgentReadinessInput is the input to get_agent_readiness.
type GetAgentReadinessInput struct {
	ClientID string `json:"clientId"`
}

// GetAgentReadiness calls get_agent_readiness: Get agent readiness. Whether
// the client's site can be read and cited by AI assistants — crawler access,
// rendering, structured data. Answers 'why are assistants not citing us'.
func (c *Client) GetAgentReadiness(ctx context.Context, in GetAgentReadinessInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_agent_readiness", in)
}

// GetAIVisibilityInput is the input to get_ai_visibility.
type GetAIVisibilityInput struct {
	ClientID string `json:"clientId"`
}

// GetAIVisibility calls get_ai_visibility: Get AI assistant visibility. How
// often the client's brand appears in AI assistant answers versus competitors,
// and what kind of pages those answers were built from. Answers 'are we
// visible in ChatGPT, Copilot, Gemini and Claude' and 'what do we have to be
// on to get cited'.
func (c *Client) GetAIVisibility(ctx context.Context, in GetAIVisibilityInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_visibility", in)
}

// GetAIVisibilityAnswerInput is the input to get_ai_visibility_answer.
type GetAIVisibilityAnswerInput struct {
	ClientID   string `json:"clientId"`
	PromptID   string `json:"promptId"`
	Model      string `json:"model"`
	CapturedAt string `json:"capturedAt"`
}

// GetAIVisibilityAnswer calls get_ai_visibility_answer: Read one AI answer in
// full. The full text one assistant gave to one question in one check, with
// the pages it cited and where each tracked brand appears in it. found is
// false when that answer was never captured or has aged out (answers are kept
// 13 months). Answers 'what exactly did ChatGPT say'. Free.
func (c *Client) GetAIVisibilityAnswer(ctx context.Context, in GetAIVisibilityAnswerInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_visibility_answer", in)
}

// GetAIVisibilityResultsInput is the input to get_ai_visibility_results.
type GetAIVisibilityResultsInput struct {
	ClientID   string `json:"clientId"`
	CapturedAt string `json:"capturedAt,omitempty"`
}

// GetAIVisibilityResults calls get_ai_visibility_results: AI visibility
// results, answer by answer. One check (the latest unless capturedAt names one
// from get_ai_visibility_runs), as a row per question and assistant: whether
// the client was named, where in the answer (character offset; earlier is more
// prominent), whether its site was cited, and which other brands were named,
// in order. Answers 'which questions are we missing from'. Free.
func (c *Client) GetAIVisibilityResults(ctx context.Context, in GetAIVisibilityResultsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_visibility_results", in)
}

// GetAIVisibilityRunsInput is the input to get_ai_visibility_runs.
type GetAIVisibilityRunsInput struct {
	ClientID string `json:"clientId"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetAIVisibilityRuns calls get_ai_visibility_runs: AI visibility checks so
// far. Lists the client's checks, newest first: when, how many answers, which
// assistants, and in how many the client's own brand was named (mentionRate;
// null when no brand is marked as the client's). Answers 'when did we last
// check, and how did it go'. Free.
func (c *Client) GetAIVisibilityRuns(ctx context.Context, in GetAIVisibilityRunsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_visibility_runs", in)
}

// GetAIVisibilitySetupInput is the input to get_ai_visibility_setup.
type GetAIVisibilitySetupInput struct {
	ClientID string `json:"clientId"`
}

// GetAIVisibilitySetup calls get_ai_visibility_setup: AI visibility setup.
// Reads what is tracked for a client: the questions (with their market), the
// brands looked for (which one is the client's own), the schedule (cadence,
// which assistants, last run), and which assistants this agency can ask at all
// and how (its own key or its proxy pool). Answers 'what are we measuring in
// AI assistants'. Free.
func (c *Client) GetAIVisibilitySetup(ctx context.Context, in GetAIVisibilitySetupInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_visibility_setup", in)
}

// GetAIVisibilityTrendInput is the input to get_ai_visibility_trend.
type GetAIVisibilityTrendInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
}

// GetAIVisibilityTrend calls get_ai_visibility_trend: Is AI visibility
// improving. The client's own mention rate over the last N days (default 28)
// against the N days before, overall and per assistant, counted by answer,
// with how many answers each side rests on. Answers 'are we being recommended
// more than last month'. Free.
func (c *Client) GetAIVisibilityTrend(ctx context.Context, in GetAIVisibilityTrendInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_visibility_trend", in)
}

// GetAuditSummaryInput is the input to get_audit_summary.
type GetAuditSummaryInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"`
}

// GetAuditSummary calls get_audit_summary: Get site audit summary. Technical
// issues one crawl found, grouped by rule: each with its name, severity,
// category, how many pages it affects and howToFix. Answers 'what is broken on
// the site'. crawlId defaults to the latest completed crawl. Then
// list_issue_pages shows which pages an issue is on.
func (c *Client) GetAuditSummary(ctx context.Context, in GetAuditSummaryInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_audit_summary", in)
}

// GetBacklinkAuditInput is the input to get_backlink_audit.
type GetBacklinkAuditInput struct {
	ClientID   string   `json:"clientId"`
	MoneyTerms []string `json:"moneyTerms,omitempty"`
}

// GetBacklinkAudit calls get_backlink_audit: Audit a client's backlinks. The
// client's backlinks, each scored from signals an agency can check — source
// authority, how many sites the source links out to, and the anchor — with a
// verdict of healthy, review or toxic and the reasons. Also the top referring
// domains, the top anchors, and the links our own crawler found (page-level,
// ours alone). Brand terms come from the client's name and domain, so a
// branded anchor is not read as paid; pass moneyTerms (commercial phrases such
// as 'cheap flights') to flag exact-match anchors. Answers 'are any of our
// links a problem'. 'review' means look at it, not that it is bad. domain is
// null when the client has no site yet. Free.
func (c *Client) GetBacklinkAudit(ctx context.Context, in GetBacklinkAuditInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_backlink_audit", in)
}

// GetBacklinkSpamScoresInput is the input to get_backlink_spam_scores.
type GetBacklinkSpamScoresInput struct {
	ClientID      string   `json:"clientId"`
	Domains       []string `json:"domains"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// GetBacklinkSpamScores calls get_backlink_spam_scores: The vendor's spam
// scores for linking sites. The data vendor's own spam score (0–100, null when
// it has none) for each domain you pass — usually the sources from
// get_backlink_audit. A second opinion beside ours: a link both call bad is
// one to act on. Paid, one purchase per call however many domains (up to
// 1000): quote it with estimate_cost and get the user's go-ahead first.
// Answers 'what does the vendor think of these links'.
func (c *Client) GetBacklinkSpamScores(ctx context.Context, in GetBacklinkSpamScoresInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_backlink_spam_scores", in)
}

// GetBillingInput is the input to get_billing.
type GetBillingInput struct {
}

// GetBilling calls get_billing: Read the agency's plan. The agency's plan, its
// limits, how many clients it uses, and the subscription's status and renewal
// date. Answers 'how many more clients can we add'. `configured` is false
// where this deployment sells nothing — every agency is on the trial plan
// then, and that is not an error.
func (c *Client) GetBilling(ctx context.Context, in GetBillingInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_billing", in)
}

// GetBillingLinkInput is the input to get_billing_link.
type GetBillingLinkInput struct {
	Plan string `json:"plan,omitempty"`
}

// GetBillingLink calls get_billing_link: Get a link to pay or manage the plan.
// A Stripe address for a person to open: checkout for a plan (studio, agency
// or network), or — with no plan, for an agency already paying — the portal
// where cards, invoices and cancellation live. Answers 'upgrade our plan'. An
// agent cannot pay; hand the address to a person. Refused where this
// deployment has no billing set up.
func (c *Client) GetBillingLink(ctx context.Context, in GetBillingLinkInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_billing_link", in)
}

// GetBingPerformanceInput is the input to get_bing_performance.
type GetBingPerformanceInput struct {
	ClientID string `json:"clientId"`
}

// GetBingPerformance calls get_bing_performance: What Bing Webmaster Tools
// reports. The client's site as Bing reports it, on the agency's Webmaster
// Tools key: daily clicks and impressions, page and query stats, crawl stats
// (pages crawled, in index, errors by status), crawl issues per URL, sitemaps
// and feeds, and the URL submission quota. Answers 'how is this site doing in
// Bing'. These are Bing's numbers about Bing's searchers and must never be
// added to Google's. Each part is separately present or failed; connected is
// false when the agency has no Bing key.
func (c *Client) GetBingPerformance(ctx context.Context, in GetBingPerformanceInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_bing_performance", in)
}

// GetBrandMentionsInput is the input to get_brand_mentions.
type GetBrandMentionsInput struct {
	ClientID      string `json:"clientId"` // The client whose account and monthly cap pay for the lookup
	Keyword       string `json:"keyword"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// GetBrandMentions calls get_brand_mentions: Find recent mentions of a brand
// or phrase. Who has written about a brand or phrase lately — the page, its
// title and snippet, when it was published, and how it read (positive,
// negative or neutral) — one result per publishing domain, up to 50. Answers
// 'who is talking about us' and finds unlinked mentions worth asking for a
// link. A paid lookup on the client's account: quote it with estimate_cost.
func (c *Client) GetBrandMentions(ctx context.Context, in GetBrandMentionsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_brand_mentions", in)
}

// GetBrandingInput is the input to get_branding.
type GetBrandingInput struct {
}

// GetBranding calls get_branding: Read the agency's branding. The agency's
// name, logo address and primary colour: what clients see on reports and
// portal links. Answers 'what will this report look like to the client'. Empty
// strings where nothing is set.
func (c *Client) GetBranding(ctx context.Context, in GetBrandingInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_branding", in)
}

// GetCannibalizationInput is the input to get_cannibalization.
type GetCannibalizationInput struct {
	ClientID string `json:"clientId"`
}

// GetCannibalization calls get_cannibalization: Get cannibalization. Keywords
// where several of the client's own pages compete in the same result. Answers
// 'where are we competing with ourselves'.
func (c *Client) GetCannibalization(ctx context.Context, in GetCannibalizationInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_cannibalization", in)
}

// GetClickCurveInput is the input to get_click_curve.
type GetClickCurveInput struct {
	Engine string `json:"engine,omitempty"`
	Market string `json:"market,omitempty"`
	Query  string `json:"query,omitempty"`
}

// GetClickCurve calls get_click_curve: How clicks fall away down a results
// page. The share of clicks each position took, as the panel observed it — for
// one query, or across a whole market. Answers 'what is a position actually
// worth here' with a measurement rather than a published average. Positions
// too few people clicked are withheld, so the shares deliberately do not add
// up to one: the rest of the page is unmeasured, not empty.
func (c *Client) GetClickCurve(ctx context.Context, in GetClickCurveInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_click_curve", in)
}

// GetClientInput is the input to get_client.
type GetClientInput struct {
	ClientID string `json:"clientId"`
}

// GetClient calls get_client: Get a client. One client's settings with its
// projects, sites and rank trackers. Answers 'what is set up for this client'
// — use it to find tracker, project and site ids before changing anything.
func (c *Client) GetClient(ctx context.Context, in GetClientInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_client", in)
}

// GetClientBriefInput is the input to get_client_brief.
type GetClientBriefInput struct {
	ClientID string `json:"clientId"`
}

// GetClientBrief calls get_client_brief: Read the client brief. Reads what the
// agency knows about a client: overview, current goal, positioning, audience,
// writing voice, competitors, key pages, custom sections, and the research log
// (newest first). Answers 'what do we know about this client'. Read it before
// researching, writing or reporting for a client, so the work matches the
// business and builds on earlier findings. Fields nobody has written are null,
// not empty. Free.
func (c *Client) GetClientBrief(ctx context.Context, in GetClientBriefInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_client_brief", in)
}

// GetClientCostsInput is the input to get_client_costs.
type GetClientCostsInput struct {
	ClientID  string `json:"clientId"`
	SinceDays *int64 `json:"sinceDays,omitempty"`
}

// GetClientCosts calls get_client_costs: Vendor spend for one client. What one
// client's data cost over recent days, by operation, with request counts.
// Answers 'what are we paying for on this account'. Amounts are US dollars.
func (c *Client) GetClientCosts(ctx context.Context, in GetClientCostsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_client_costs", in)
}

// GetClientIntegrationsInput is the input to get_client_integrations.
type GetClientIntegrationsInput struct {
	ClientID string `json:"clientId"`
}

// GetClientIntegrations calls get_client_integrations: Which consoles are
// connected for a client. Whether a client's Google Search Console and Google
// Analytics properties are connected and syncing, with the last sync error,
// and whether IndexNow is set up. Answers 'is Search Console connected for
// this client'. Never returns tokens.
func (c *Client) GetClientIntegrations(ctx context.Context, in GetClientIntegrationsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_client_integrations", in)
}

// GetConnectionLinkInput is the input to get_connection_link.
type GetConnectionLinkInput struct {
	ClientID      string `json:"clientId"`
	Service       string `json:"service"`
	PropertyURL   string `json:"propertyUrl,omitempty"`
	GA4PropertyID string `json:"ga4PropertyId,omitempty"`
	LocationID    string `json:"locationId,omitempty"`
}

// GetConnectionLink calls get_connection_link: Get a link for a person to
// connect a service. The link a PERSON must open to connect one of a client's
// accounts — Google Search Console (search_console), Search Console sitemap
// management (search_console_sitemaps), Google Analytics 4 (google_analytics),
// Yandex Webmaster (yandex_webmaster) or a Google Business Profile listing
// (google_business_profile) — with what they will be asked to approve and what
// is connected now. Answers 'connect Search Console for this client' and 'did
// the connection work'. An agent cannot complete these itself: the account
// owner signs in to that service and grants access. Give the url and the steps
// to your human, then call this again to check state. The person must be
// signed in to Seolyticize with permission to manage the client.
// google_analytics needs ga4PropertyId (the numeric property id, not a G-
// measurement id); google_business_profile needs locationId; search_console
// takes an optional propertyUrl (default: the client's site, https://domain/;
// a Domain property is sc-domain:example.com). Free.
func (c *Client) GetConnectionLink(ctx context.Context, in GetConnectionLinkInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_connection_link", in)
}

// GetContentDecayInput is the input to get_content_decay.
type GetContentDecayInput struct {
	ClientID string `json:"clientId"`
}

// GetContentDecay calls get_content_decay: Get decaying content. Pages losing
// the traffic they used to earn, over consecutive periods. Answers 'what
// should we refresh'.
func (c *Client) GetContentDecay(ctx context.Context, in GetContentDecayInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_content_decay", in)
}

// GetContentGapInput is the input to get_content_gap.
type GetContentGapInput struct {
	ClientID          string   `json:"clientId"`
	CompetitorDomains []string `json:"competitorDomains"`
}

// GetContentGap calls get_content_gap: Get content gap. Competitor pages that
// own a cluster of keywords the client does not cover. Answers 'what page
// should we build next' — use this rather than keyword gap when planning
// content.
func (c *Client) GetContentGap(ctx context.Context, in GetContentGapInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_content_gap", in)
}

// GetCostsInput is the input to get_costs.
type GetCostsInput struct {
	SinceDays *int64 `json:"sinceDays,omitempty"`
}

// GetCosts calls get_costs: Vendor spend by client. What the agency spent on
// data vendors per client over recent days, with request counts. Answers
// 'which clients cost us the most'. Amounts are US dollars.
func (c *Client) GetCosts(ctx context.Context, in GetCostsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_costs", in)
}

// GetCrawlStatusInput is the input to get_crawl_status.
type GetCrawlStatusInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId"`
}

// GetCrawlStatus calls get_crawl_status: Get crawl status. Whether one crawl
// is running, complete or failed, with page and issue counts and the failure
// reason. Answers 'has the audit finished'.
func (c *Client) GetCrawlStatus(ctx context.Context, in GetCrawlStatusInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_crawl_status", in)
}

// GetDeployHookInput is the input to get_deploy_hook.
type GetDeployHookInput struct {
	ClientID string `json:"clientId"`
}

// GetDeployHook calls get_deploy_hook: Get the deploy hook's status. Answers
// 'did the last deploy notify search engines'. Whether the client has a deploy
// hook, how many sitemap pages it knows, and what its last call did. Never the
// URL itself. Free.
func (c *Client) GetDeployHook(ctx context.Context, in GetDeployHookInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_deploy_hook", in)
}

// GetDomainFactsInput is the input to get_domain_facts.
type GetDomainFactsInput struct {
	Domain        string `json:"domain"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// GetDomainFacts calls get_domain_facts: What the open web says about a
// domain. Registration date and registrar from the registry, position on the
// Tranco top-sites list, when the Internet Archive first captured it, and the
// subdomains its certificates have announced. Answers 'how old and how big is
// this, and what else is on it' for any domain, free and without touching a
// vendor. Each figure is absent far more often than present, and an absence is
// never a zero: no registry record, not in the top million, never archived.
func (c *Client) GetDomainFacts(ctx context.Context, in GetDomainFactsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_domain_facts", in)
}

// GetDomainTechnologiesInput is the input to get_domain_technologies.
type GetDomainTechnologiesInput struct {
	Domain string `json:"domain"`
}

// GetDomainTechnologies calls get_domain_technologies: See what a domain is
// built with. The technologies our own crawler detected on any domain — CMS,
// analytics, commerce, CDN and so on — with the version most pages showed and
// how many pages showed it, as of the day it was last crawled. Answers 'what
// is this site built on'. Free. Null when our crawler has never fetched the
// domain, which says nothing about what it runs.
func (c *Client) GetDomainTechnologies(ctx context.Context, in GetDomainTechnologiesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_domain_technologies", in)
}

// GetEntityPageviewsInput is the input to get_entity_pageviews.
type GetEntityPageviewsInput struct {
	Entity      string `json:"entity"`
	Language    string `json:"language,omitempty"`
	Granularity string `json:"granularity,omitempty"`
	From        string `json:"from"`
	To          string `json:"to"`
}

// GetEntityPageviews calls get_entity_pageviews: How many people looked a
// brand or topic up. Wikipedia article views for a brand or topic over a
// period, daily or monthly. Answers 'is interest in this name growing' with a
// complete count published by Wikimedia rather than a vendor's model or a
// panel sample. Returns nothing at all when Wikipedia has no article, which is
// the case for most small brands — and says which article it counted, because
// the closest article is not always the brand's own.
func (c *Client) GetEntityPageviews(ctx context.Context, in GetEntityPageviewsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_entity_pageviews", in)
}

// GetGA4AudienceInput is the input to get_ga4_audience.
type GetGA4AudienceInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Compare  *bool  `json:"compare,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	By       string `json:"by,omitempty"`
	Channel  string `json:"channel,omitempty"`
}

// GetGA4Audience calls get_ga4_audience: Who arrives (GA4). Users, sessions,
// engaged sessions and key events by device, country, or new versus returning
// (by), organic search by default. Answers 'is mobile or desktop converting,
// which countries'. Rows carry key, current (metrics by name) and previous
// (null when compare is false, metrics null where that period had no row).
// Also period, previousPeriod, currency, thresholded and folded. Refuses with
// a sentence when the client has not connected a GA4 property. Free.
func (c *Client) GetGA4Audience(ctx context.Context, in GetGA4AudienceInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ga4_audience", in)
}

// GetGA4EcommerceInput is the input to get_ga4_ecommerce.
type GetGA4EcommerceInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Compare  *bool  `json:"compare,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	By       string `json:"by,omitempty"`
}

// GetGA4Ecommerce calls get_ga4_ecommerce: Shop performance (GA4). Item views,
// add-to-carts, purchases and item revenue by item, or transactions and
// purchase revenue by organic landing page (by "landingPage"). activity is
// "detected" or "none" — none means the property sends no ecommerce events for
// the period, not that nothing sold. Answers 'what is selling, and from which
// pages'. Rows carry key, current (metrics by name) and previous (null when
// compare is false, metrics null where that period had no row). Also period,
// previousPeriod, currency, thresholded and folded. Refuses with a sentence
// when the client has not connected a GA4 property. Free.
func (c *Client) GetGA4Ecommerce(ctx context.Context, in GetGA4EcommerceInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ga4_ecommerce", in)
}

// GetGA4KeyEventsInput is the input to get_ga4_key_events.
type GetGA4KeyEventsInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Compare  *bool  `json:"compare,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	By       string `json:"by,omitempty"`
}

// GetGA4KeyEvents calls get_ga4_key_events: Key events (GA4). Counts of the
// property's key events (what used to be conversions) by event name, or by the
// organic landing page that led to them (by "landingPage"). Answers 'what are
// visitors converting on'. Rows carry key, current (metrics by name) and
// previous (null when compare is false, metrics null where that period had no
// row). Also period, previousPeriod, currency, thresholded and folded. Refuses
// with a sentence when the client has not connected a GA4 property. Free.
func (c *Client) GetGA4KeyEvents(ctx context.Context, in GetGA4KeyEventsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ga4_key_events", in)
}

// GetGA4LandingPagesInput is the input to get_ga4_landing_pages.
type GetGA4LandingPagesInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Compare  *bool  `json:"compare,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	Channel  string `json:"channel,omitempty"`
}

// GetGA4LandingPages calls get_ga4_landing_pages: What landing pages did
// (GA4). Sessions, engaged sessions, engagement rate, key events, transactions
// and revenue by landing page, organic search by default. Answers 'which pages
// search sends people to, and what they did there'. Rows carry key, current
// (metrics by name) and previous (null when compare is false, metrics null
// where that period had no row). Also period, previousPeriod, currency,
// thresholded and folded. Refuses with a sentence when the client has not
// connected a GA4 property. Free.
func (c *Client) GetGA4LandingPages(ctx context.Context, in GetGA4LandingPagesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ga4_landing_pages", in)
}

// GetGA4MeasurementHealthInput is the input to get_ga4_measurement_health.
type GetGA4MeasurementHealthInput struct {
	ClientID string `json:"clientId"`
}

// GetGA4MeasurementHealth calls get_ga4_measurement_health: Is the property
// measuring properly (GA4). Checks the connected property's setup: web data
// streams and measurement ids, enhanced measurement per stream (scrolls,
// outbound clicks, site search, forms, downloads), marked key events, custom
// dimensions and metrics, and the last day with sessions. findings lists
// problems, warnings and notes in plain words; a part that could not be read
// is reported as unread, not as fine. Answers 'can we trust this client's
// Analytics'. Refuses with a sentence when the client has not connected a GA4
// property. Free.
func (c *Client) GetGA4MeasurementHealth(ctx context.Context, in GetGA4MeasurementHealthInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ga4_measurement_health", in)
}

// GetGA4OverviewInput is the input to get_ga4_overview.
type GetGA4OverviewInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Compare  *bool  `json:"compare,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	Channel  string `json:"channel,omitempty"`
}

// GetGA4Overview calls get_ga4_overview: Is organic traffic growing (GA4).
// Top-line sessions, users, engaged sessions, engagement rate, key events,
// transactions and revenue for organic search (channel "all" for every
// channel), against the previous period, plus a daily sessions and key-events
// line. Answers 'is organic traffic improving'. total holds the figures. Rows
// carry key, current (metrics by name) and previous (null when compare is
// false, metrics null where that period had no row). Also period,
// previousPeriod, currency, thresholded and folded. Refuses with a sentence
// when the client has not connected a GA4 property. Free.
func (c *Client) GetGA4Overview(ctx context.Context, in GetGA4OverviewInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ga4_overview", in)
}

// GetGA4PagePerformanceInput is the input to get_ga4_page_performance.
type GetGA4PagePerformanceInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Compare  *bool  `json:"compare,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	Channel  string `json:"channel,omitempty"`
}

// GetGA4PagePerformance calls get_ga4_page_performance: Page views and
// engagement (GA4). Views, users, total engagement time in seconds and key
// events by page path, organic search by default (channel "all" for every
// channel). Answers 'which pages hold attention'. Rows carry key, current
// (metrics by name) and previous (null when compare is false, metrics null
// where that period had no row). Also period, previousPeriod, currency,
// thresholded and folded. Refuses with a sentence when the client has not
// connected a GA4 property. Free.
func (c *Client) GetGA4PagePerformance(ctx context.Context, in GetGA4PagePerformanceInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ga4_page_performance", in)
}

// GetGA4SiteSearchInput is the input to get_ga4_site_search.
type GetGA4SiteSearchInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Compare  *bool  `json:"compare,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetGA4SiteSearch calls get_ga4_site_search: What visitors search for on the
// site (GA4). Terms people typed into the site's own search, with searches and
// users. measured is false when the property records no site search (enhanced
// measurement off, or no search box). Answers 'what are people looking for
// that the site does not surface'. Rows carry key, current (metrics by name)
// and previous (null when compare is false, metrics null where that period had
// no row). Also period, previousPeriod, currency, thresholded and folded.
// Refuses with a sentence when the client has not connected a GA4 property.
// Free.
func (c *Client) GetGA4SiteSearch(ctx context.Context, in GetGA4SiteSearchInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ga4_site_search", in)
}

// GetGA4TrafficAcquisitionInput is the input to get_ga4_traffic_acquisition.
type GetGA4TrafficAcquisitionInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Compare  *bool  `json:"compare,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	GroupBy  string `json:"groupBy,omitempty"`
}

// GetGA4TrafficAcquisition calls get_ga4_traffic_acquisition: Where traffic
// comes from (GA4). Sessions, users, engaged sessions, key events,
// transactions and revenue by channel group, source/medium or campaign
// (groupBy), against the previous period. Answers 'which channels grew or
// fell'. Rows carry key, current (metrics by name) and previous (null when
// compare is false, metrics null where that period had no row). Also period,
// previousPeriod, currency, thresholded and folded. Refuses with a sentence
// when the client has not connected a GA4 property. Free.
func (c *Client) GetGA4TrafficAcquisition(ctx context.Context, in GetGA4TrafficAcquisitionInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ga4_traffic_acquisition", in)
}

// GetGoogleServiceAccountInput is the input to get_google_service_account.
type GetGoogleServiceAccountInput struct {
}

// GetGoogleServiceAccount calls get_google_service_account: Get the agency's
// Google service account. The email address of the agency's own Google service
// account, made the first time it is asked for. Answers 'what do I add in
// Search Console so we can connect without signing in'. A client's property
// owner adds this address as a user — Search Console: Settings → Users and
// permissions, Restricted; GA4: Admin → Property access management, Viewer —
// and then connect_search_console_service_account or
// connect_google_analytics_service_account connects it with no Google sign-in.
// The address is the agency's alone; give it to your human to pass on. Free.
func (c *Client) GetGoogleServiceAccount(ctx context.Context, in GetGoogleServiceAccountInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_google_service_account", in)
}

// GetIndexCoverageInput is the input to get_index_coverage.
type GetIndexCoverageInput struct {
	ClientID string `json:"clientId"`
}

// GetIndexCoverage calls get_index_coverage: Which of the client's pages
// Google has indexed. The newest URL Inspection verdict for every page of the
// client's site that has been inspected, grouped by verdict and by Google's
// coverage state, with pages whose Google-chosen canonical differs from their
// own. Answers 'how much of this site is in Google, and what is keeping the
// rest out'. Covers inspected pages only — Search Console's own indexing
// report has no API — and says how many that is and how many are still queued.
func (c *Client) GetIndexCoverage(ctx context.Context, in GetIndexCoverageInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_index_coverage", in)
}

// GetIndexnowKeyInput is the input to get_indexnow_key.
type GetIndexnowKeyInput struct {
	ClientID string `json:"clientId"`
}

// GetIndexnowKey calls get_indexnow_key: Get the IndexNow key. The client's
// IndexNow key, the file the site must serve for it, and whether that file has
// been verified. Lists what submit_indexnow needs in place. A null key means
// none has been issued yet: use issue_indexnow_key.
func (c *Client) GetIndexnowKey(ctx context.Context, in GetIndexnowKeyInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_indexnow_key", in)
}

// GetKeywordDifficultyInput is the input to get_keyword_difficulty.
type GetKeywordDifficultyInput struct {
	ClientID   string   `json:"clientId"`
	KeywordIDs []string `json:"keywordIds"`
}

// GetKeywordDifficulty calls get_keyword_difficulty: Get keyword difficulty.
// How hard the first page looks for given keywords, scored from how many
// domains link to the pages already ranking there. Answers 'is this term worth
// going after'. Null where too few results have link data to judge.
func (c *Client) GetKeywordDifficulty(ctx context.Context, in GetKeywordDifficultyInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_keyword_difficulty", in)
}

// GetKeywordGapInput is the input to get_keyword_gap.
type GetKeywordGapInput struct {
	ClientID          string   `json:"clientId"`
	CompetitorDomains []string `json:"competitorDomains"`
}

// GetKeywordGap calls get_keyword_gap: Get keyword gap. Keywords competitors
// rank for that the client does not, sized by the clicks available. Answers
// 'which terms are we missing'.
func (c *Client) GetKeywordGap(ctx context.Context, in GetKeywordGapInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_keyword_gap", in)
}

// GetKeywordListInput is the input to get_keyword_list.
type GetKeywordListInput struct {
	ListID string `json:"listId"`
}

// GetKeywordList calls get_keyword_list: Read a keyword list. Every keyword in
// one list (by listId from list_keyword_lists): the term, its market, language
// and engine, its keywordId, and its note. Answers 'what is on this list'.
// Free.
func (c *Client) GetKeywordList(ctx context.Context, in GetKeywordListInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_keyword_list", in)
}

// GetKeywordTrendsInput is the input to get_keyword_trends.
type GetKeywordTrendsInput struct {
	ClientID      string   `json:"clientId"` // The client whose account and monthly cap pay for the lookup
	Terms         []string `json:"terms"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// GetKeywordTrends calls get_keyword_trends: Get Google Trends interest for
// terms. Interest over time from Google Trends for up to five terms in the
// United States, each on a 0–100 scale relative to its own peak — so it shows
// when people care, not how many. Answers 'is this seasonal' and 'is this
// growing'. Different from get_search_demand, which counts people in our own
// panel. A paid lookup on the client's account: quote it with estimate_cost.
func (c *Client) GetKeywordTrends(ctx context.Context, in GetKeywordTrendsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_keyword_trends", in)
}

// GetLinkChecksInput is the input to get_link_checks.
type GetLinkChecksInput struct {
	ClientID string `json:"clientId"`
}

// GetLinkChecks calls get_link_checks: Do the linking pages still link?. What
// our own crawler last saw on each page said to link to the client: live (the
// link is there), lost (the page no longer links) or unknown (the page could
// not be read), with counts. Read from the last check, never fetched now.
// Answers 'which links have we lost'. Run recheck_links to check again. Free.
func (c *Client) GetLinkChecks(ctx context.Context, in GetLinkChecksInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_link_checks", in)
}

// GetLinkIntersectInput is the input to get_link_intersect.
type GetLinkIntersectInput struct {
	ClientID          string   `json:"clientId"`
	CompetitorDomains []string `json:"competitorDomains"`
}

// GetLinkIntersect calls get_link_intersect: Get link prospects. Sites linking
// to the client's competitors but not to the client. Answers 'who should we
// approach for links'.
func (c *Client) GetLinkIntersect(ctx context.Context, in GetLinkIntersectInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_link_intersect", in)
}

// GetLiveSerpsInput is the input to get_live_serps.
type GetLiveSerpsInput struct {
	ClientID      string   `json:"clientId"`
	Keywords      []string `json:"keywords"`
	LocationCode  *int64   `json:"locationCode,omitempty"`
	LanguageCode  string   `json:"languageCode,omitempty"`
	Engine        string   `json:"engine,omitempty"`
	Device        string   `json:"device,omitempty"`
	Top           *int64   `json:"top,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// GetLiveSerps calls get_live_serps: Read several results pages now. Reads the
// live results page for up to 10 keywords in one market and engine (Google by
// default), and returns each page's top results (position, type, domain, url,
// title) and the search features on it. Answers 'who ranks for these right
// now'. Each keyword is one purchase, made in turn; one that fails is reported
// and the rest still return; maxCostMicros stops before the purchase that
// would pass it. Quote with estimate_cost.
func (c *Client) GetLiveSerps(ctx context.Context, in GetLiveSerpsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_live_serps", in)
}

// GetLocalVisibilityInput is the input to get_local_visibility.
type GetLocalVisibilityInput struct {
	ClientID   string `json:"clientId"`
	LocationID string `json:"locationId"`
	Term       string `json:"term"`
}

// GetLocalVisibility calls get_local_visibility: Map-pack visibility and
// reviews for a location. The latest grid scan for a search term at one
// location — the map-pack position at every point, and the share of points
// ranking in the top 3 and top 10 — plus an analysis of the location's reviews
// (rating, volume, recency, response rate). Answers 'where in town do we show
// up for this search'. scanned is false when that term has never been scanned
// there: no grid, rather than a grid of blanks that would read as invisible
// everywhere. Free.
func (c *Client) GetLocalVisibility(ctx context.Context, in GetLocalVisibilityInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_local_visibility", in)
}

// GetLowCtrQueriesInput is the input to get_low_ctr_queries.
type GetLowCtrQueriesInput struct {
	ClientID       string `json:"clientId"`
	Days           *int64 `json:"days,omitempty"`
	MinImpressions *int64 `json:"minImpressions,omitempty"`
	Limit          *int64 `json:"limit,omitempty"`
}

// GetLowCtrQueries calls get_low_ctr_queries: Queries losing clicks to their
// snippet. Search Console queries on page one (positions 1–10) with enough
// impressions where far fewer people click than that position normally earns
// (under half the click curve's rate): actual and expected click-through rate
// and the clicks missed, worst first. Answers 'which titles and descriptions
// should we rewrite'. Check get_serp_features first — an AI overview or map
// pack above the result also lowers clicks. Free.
func (c *Client) GetLowCtrQueries(ctx context.Context, in GetLowCtrQueriesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_low_ctr_queries", in)
}

// GetNextStepsInput is the input to get_next_steps.
type GetNextStepsInput struct {
	ClientID string `json:"clientId"`
}

// GetNextSteps calls get_next_steps: What to do next for a site. Answers 'what
// is left to do for this site, and what call does it'. The onboarding
// checklist with live status — done, todo, in progress, failed or blocked and
// why — each with the exact tool and arguments to call next, then, once the
// site is set up, the weekly loop's calls with their arguments. Reads only;
// changes nothing. Free.
func (c *Client) GetNextSteps(ctx context.Context, in GetNextStepsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_next_steps", in)
}

// GetOpportunitiesInput is the input to get_opportunities.
type GetOpportunitiesInput struct {
	ClientID string `json:"clientId"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetOpportunities calls get_opportunities: Everything worth working on, in
// one call. Gathers the four kinds of opportunity from stored data:
// strikingDistance (pages around positions 4–20 by extra clicks), lowCtr
// (page-one queries losing clicks to their snippet), contentDecay (pages
// losing clicks), and cannibalization (pages splitting one query). Each part
// is a section with data or an error — an unread part is not an empty one.
// Answers 'what should we work on for this client'. Use the individual tools
// for detail. Free.
func (c *Client) GetOpportunities(ctx context.Context, in GetOpportunitiesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_opportunities", in)
}

// GetPageIssuesInput is the input to get_page_issues.
type GetPageIssuesInput struct {
	ClientID string `json:"clientId"`
	URL      string `json:"url"`
	CrawlID  string `json:"crawlId,omitempty"`
}

// GetPageIssues calls get_page_issues: Get every issue on one page. Lists
// every issue a crawl found on one page, with severity, message and howToFix.
// Answers 'what is wrong with this page' and, after a fix and a new crawl, 'is
// it fixed'. crawlId defaults to the latest completed crawl. Free.
func (c *Client) GetPageIssues(ctx context.Context, in GetPageIssuesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_page_issues", in)
}

// GetPanelReachInput is the input to get_panel_reach.
type GetPanelReachInput struct {
	Hosts []string `json:"hosts"`
	From  string   `json:"from"`
	To    string   `json:"to"`
}

// GetPanelReach calls get_panel_reach: Get panel reach. How many panel members
// visited given hosts over a period, and how often. Answers 'what did real
// people actually do', as far as the panel saw. Returns the panelist count
// alongside the visits so a thin panel is visibly thin rather than quietly
// wrong.
func (c *Client) GetPanelReach(ctx context.Context, in GetPanelReachInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_panel_reach", in)
}

// GetPositionDistributionInput is the input to get_position_distribution.
type GetPositionDistributionInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId"`
	Days      *int64 `json:"days,omitempty"`
	Device    string `json:"device,omitempty"`
}

// GetPositionDistribution calls get_position_distribution: Get position
// distribution. How a tracker's keywords spread across positions 1-3, 4-10,
// 11-20, 21-50 and 51-100 each day, with average position and visibility, per
// device. Answers 'how many keywords are on page one, and is that growing'.
// Keywords read and not found are counted as not ranking, apart from every
// band; keywords not read that day are counted as not captured. The average is
// over ranking keywords only.
func (c *Client) GetPositionDistribution(ctx context.Context, in GetPositionDistributionInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_position_distribution", in)
}

// GetRankHistoryInput is the input to get_rank_history.
type GetRankHistoryInput struct {
	ClientID   string   `json:"clientId"`
	TrackerID  string   `json:"trackerId"`
	KeywordIDs []string `json:"keywordIds,omitempty"`
	Days       *int64   `json:"days,omitempty"`
	Device     string   `json:"device,omitempty"`
}

// GetRankHistory calls get_rank_history: Get rank history. Day-by-day
// positions for a tracker's keywords, each day with the page that ranked and
// the search features on the page — and which of those the client held.
// Answers 'how has this keyword moved, and with which page'. Mobile and
// desktop are separate series. A null position means the page was read and the
// client was not on it; a missing day means nobody read the page that day.
// Neither is a zero.
func (c *Client) GetRankHistory(ctx context.Context, in GetRankHistoryInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_rank_history", in)
}

// GetRankSummaryInput is the input to get_rank_summary.
type GetRankSummaryInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId"`
}

// GetRankSummary calls get_rank_summary: Get rank summary. How many tracked
// keywords improved, declined or held for a client, with the average position.
// Answers 'are rankings getting better'.
func (c *Client) GetRankSummary(ctx context.Context, in GetRankSummaryInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_rank_summary", in)
}

// GetRunStatusInput is the input to get_run_status.
type GetRunStatusInput struct {
	ClientID string `json:"clientId"`
	Job      string `json:"job"`
	RunID    string `json:"runId"`
}

// GetRunStatus calls get_run_status: Get the status of a run. How a run
// started with refresh_rankings or refresh_domain_data is going: running, done
// or failed, and — when it failed or captured only some keywords — the reason,
// written for an agency to act on (a missing vendor key, an unreachable site).
// Answers 'did that capture work'. A run id that is not this client's, or that
// has aged out, reads as unknown.
func (c *Client) GetRunStatus(ctx context.Context, in GetRunStatusInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_run_status", in)
}

// GetSearchConsolePerformanceInput is the input to get_search_console_performance.
type GetSearchConsolePerformanceInput struct {
	ClientID   string   `json:"clientId"`
	SearchType string   `json:"searchType,omitempty"`
	From       string   `json:"from"`
	To         string   `json:"to"`
	Country    string   `json:"country,omitempty"`
	Device     string   `json:"device,omitempty"`
	Dimensions []string `json:"dimensions,omitempty"`
	Page       string   `json:"page,omitempty"`
	Query      string   `json:"query,omitempty"`
}

// GetSearchConsolePerformance calls get_search_console_performance: What
// Google Search Console measured. Clicks, impressions and position from the
// client's own Google Search Console, for one search type (web, image, video,
// news, discover, googleNews) over a date range, optionally filtered to one
// country (ISO 3166-1 alpha-3, e.g. gbr) and one device. Ask for any of the
// dimensions date, query, page, country and device; asking for both page and
// query returns the page-by-query drill-down, which has no country or device
// split. Answers 'what did this site actually earn in Google'. Discover and
// Google News have no queries and no position. Days Google is still counting
// are marked provisional. Returns a refusal with a reason when Search Console
// is not connected or a filter does not apply.
func (c *Client) GetSearchConsolePerformance(ctx context.Context, in GetSearchConsolePerformanceInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_search_console_performance", in)
}

// GetSearchDemandInput is the input to get_search_demand.
type GetSearchDemandInput struct {
	Query  string `json:"query"`
	Market string `json:"market"`
	Since  string `json:"since"`
}

// GetSearchDemand calls get_search_demand: What the panel saw people search
// for. How many people in the panel searched a term, and how often, in one
// market. Answers 'do real people actually search this' — a reading from
// people rather than a vendor's model. Returns nothing at all when too few
// people searched it, because a term seen by two panelists is those two people
// rather than a measurement.
func (c *Client) GetSearchDemand(ctx context.Context, in GetSearchDemandInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_search_demand", in)
}

// GetSearchOpportunitiesInput is the input to get_search_opportunities.
type GetSearchOpportunitiesInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Compare  *bool  `json:"compare,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetSearchOpportunities calls get_search_opportunities: Pages worth an edit
// first (Search Console × GA4). Joins Search Console pages ranking around
// positions 4–20 with what their organic visitors do in GA4 over the same
// window, and ranks them by projected revenue, then projected key events, then
// extra clicks. Each row: page, extraClicks (from the click curve at third
// place), its top queries, the page's organic sessions, key event rate and
// revenue per session (analytics null where GA4 has no row for it — unknown,
// not worthless), and the projections. searchConsole is false when the client
// has no site or console data. Answers 'which page should we improve first'.
// Refuses with a sentence when the client has not connected a GA4 property.
// Free.
func (c *Client) GetSearchOpportunities(ctx context.Context, in GetSearchOpportunitiesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_search_opportunities", in)
}

// GetSerpFeaturesInput is the input to get_serp_features.
type GetSerpFeaturesInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId"`
	Device    string `json:"device,omitempty"`
}

// GetSerpFeatures calls get_serp_features: Get SERP features. Which search
// features — featured snippet, AI Overview, local pack, video, People Also Ask
// and others — appeared on each tracked keyword's latest page, and which the
// client holds, with totals per feature. Answers 'which snippets and overviews
// do we own'. Features are null for a capture whose source did not report
// them, which is not the same as a page with none.
func (c *Client) GetSerpFeatures(ctx context.Context, in GetSerpFeaturesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_serp_features", in)
}

// GetShareOfVoiceInput is the input to get_share_of_voice.
type GetShareOfVoiceInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId"`
}

// GetShareOfVoice calls get_share_of_voice: Get share of voice. The client's
// share of available search clicks across its tracked keywords, alongside
// competitors. Answers 'how much of the market do we hold'.
func (c *Client) GetShareOfVoice(ctx context.Context, in GetShareOfVoiceInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_share_of_voice", in)
}

// GetSiteStatusInput is the input to get_site_status.
type GetSiteStatusInput struct {
	ClientID string `json:"clientId"`
}

// GetSiteStatus calls get_site_status: Site verification and connection
// status. For a client's sites: whether the IndexNow key file has been
// published and verified, which Search Console properties are connected and
// syncing, and the latest crawl. Answers 'is this site set up and verified'.
func (c *Client) GetSiteStatus(ctx context.Context, in GetSiteStatusInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_site_status", in)
}

// GetSubdomainTrafficInput is the input to get_subdomain_traffic.
type GetSubdomainTrafficInput struct {
	ClientID      string `json:"clientId"` // The client whose account and monthly cap pay for the lookup
	Domain        string `json:"domain"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// GetSubdomainTraffic calls get_subdomain_traffic: See which subdomains carry
// a site's search traffic. A site's subdomains ranked by estimated organic
// traffic and ranking keywords, up to 50. Answers 'which parts of this site
// carry its search presence'. get_domain_facts lists subdomain names from
// certificate logs for free but says nothing of their traffic. A paid lookup
// on the client's account: quote it with estimate_cost.
func (c *Client) GetSubdomainTraffic(ctx context.Context, in GetSubdomainTrafficInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_subdomain_traffic", in)
}

// GetTopSearchesInput is the input to get_top_searches.
type GetTopSearchesInput struct {
	Market string `json:"market"`
	Since  string `json:"since"`
	Limit  *int64 `json:"limit,omitempty"`
}

// GetTopSearches calls get_top_searches: The terms the panel searched most.
// What a market actually searched for over a period, most-searched first,
// counted by distinct people rather than by events. Answers 'what is being
// looked for right now' for keyword discovery, without a vendor's list in
// between. Terms below the reporting floor never appear.
func (c *Client) GetTopSearches(ctx context.Context, in GetTopSearchesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_top_searches", in)
}

// GetURLSwapsInput is the input to get_url_swaps.
type GetURLSwapsInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId"`
	Days      *int64 `json:"days,omitempty"`
	Device    string `json:"device,omitempty"`
}

// GetURLSwaps calls get_url_swaps: Get ranking URL swaps. Tracked keywords
// whose ranking page changed over a period — which page replaced which, when,
// and at what positions — flagging pages that take turns, and any keyword
// where several of the client's pages rank at once right now. Answers 'which
// keywords have two of our pages fighting over them'. Use it alongside
// get_cannibalization, which looks at one snapshot rather than across time.
func (c *Client) GetURLSwaps(ctx context.Context, in GetURLSwapsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_url_swaps", in)
}

// GetVendorConnectionsInput is the input to get_vendor_connections.
type GetVendorConnectionsInput struct {
}

// GetVendorConnections calls get_vendor_connections: Which vendor accounts are
// connected. Which data vendors and AI assistants the agency has connected its
// own account for, and whether each last checked out. Answers 'why is this
// data missing' when a vendor was never connected. Never returns any part of a
// credential.
func (c *Client) GetVendorConnections(ctx context.Context, in GetVendorConnectionsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_vendor_connections", in)
}

// GetWebAnalyticsBreakdownInput is the input to get_web_analytics_breakdown.
type GetWebAnalyticsBreakdownInput struct {
	ClientID  string   `json:"clientId"`
	Breakdown string   `json:"breakdown"`
	Days      *float64 `json:"days,omitempty"` // The period: the last 7, 30 or 90 days, today included.
	Limit     *int64   `json:"limit,omitempty"`
}

// GetWebAnalyticsBreakdown calls get_web_analytics_breakdown: Traffic by page,
// source, campaign, country or device (first-party). Lists the top rows of one
// breakdown from Seolyticize's own web analytics: pages (pageviews),
// entry_pages and exit_pages (visits), sources and referrers (visits; key null
// means direct), campaigns (visits, with detail "source / medium"), countries
// (ISO codes; null means unknown), devices, browsers, os, or custom events
// (event count). Each row has key, visitors and count; counts says what count
// counts, and totalKeys how many rows exist before the limit. Answers 'which
// pages and sources drive visits'. Refuses with a sentence when web analytics
// has not been set up for the client. Free.
func (c *Client) GetWebAnalyticsBreakdown(ctx context.Context, in GetWebAnalyticsBreakdownInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_web_analytics_breakdown", in)
}

// GetWebAnalyticsSetupInput is the input to get_web_analytics_setup.
type GetWebAnalyticsSetupInput struct {
	ClientID string `json:"clientId"`
}

// GetWebAnalyticsSetup calls get_web_analytics_setup: Is web analytics
// installed. Whether Seolyticize's own cookieless web analytics is set up for
// a client, with the one-line tracking snippet to install, the hostnames
// allowed to report, and lastEventAt (null until the first visit is recorded).
// Answers 'is the tracking script installed' and 'what snippet do I add to the
// site'. configured is false when no snippet has been created yet. Free.
func (c *Client) GetWebAnalyticsSetup(ctx context.Context, in GetWebAnalyticsSetupInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_web_analytics_setup", in)
}

// GetWebAnalyticsSummaryInput is the input to get_web_analytics_summary.
type GetWebAnalyticsSummaryInput struct {
	ClientID string   `json:"clientId"`
	Days     *float64 `json:"days,omitempty"` // The period: the last 7, 30 or 90 days, today included.
}

// GetWebAnalyticsSummary calls get_web_analytics_summary: Site traffic summary
// (first-party). Visitors, visits, pageviews, views per visit, bounce rate and
// average visit duration from Seolyticize's own web analytics, against the
// same span before (previous is null when that span was not fully measured),
// plus a daily series, realtimeVisitors in the last 30 minutes, and notes on
// how figures are counted. Answers 'how much traffic did the site get' for a
// client without GA4. Visitors are counted per day without cookies; a figure
// of null means not measured, not zero. Refuses with a sentence when web
// analytics has not been set up for the client. Free.
func (c *Client) GetWebAnalyticsSummary(ctx context.Context, in GetWebAnalyticsSummaryInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_web_analytics_summary", in)
}

// GetWrittenReportInput is the input to get_written_report.
type GetWrittenReportInput struct {
	ClientID string `json:"clientId"`
	ReportID string `json:"reportId"`
}

// GetWrittenReport calls get_written_report: Read a written report. Reads one
// written report: its Markdown, the rendered HTML, its summary, who wrote it,
// and its public shareUrl (null while private). Answers 'show me that report'.
// Free.
func (c *Client) GetWrittenReport(ctx context.Context, in GetWrittenReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_written_report", in)
}

// GetYandexPerformanceInput is the input to get_yandex_performance.
type GetYandexPerformanceInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from"`
	To       string `json:"to"`
}

// GetYandexPerformance calls get_yandex_performance: What Yandex Webmaster
// reports. The client's site as Yandex Webmaster reports it: site quality
// index, searchable and excluded pages, problems by severity, popular search
// queries with impressions, clicks and position, indexing history by HTTP
// status, sitemaps, and the recrawl queue with its quota. Answers 'how is this
// site doing in Yandex'. configured is false when this deployment has no
// Yandex app, and connected is false until the client's site owner has granted
// access.
func (c *Client) GetYandexPerformance(ctx context.Context, in GetYandexPerformanceInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_yandex_performance", in)
}

// InspectURLInput is the input to inspect_url.
type InspectURLInput struct {
	ClientID string `json:"clientId"`
	URL      string `json:"url"`
}

// InspectURL calls inspect_url: Ask Google what it holds for one page. Runs
// Google's URL Inspection on one address of the client's own site: whether it
// is indexed, Google's coverage state, when it was last crawled, Google's
// chosen canonical against the site's own, mobile usability (deprecated by
// Google and often absent) and rich results detected. Answers 'is this page in
// Google, and if not why'. Spends one of the property's 2,000 daily
// inspections; refuses addresses that are not on a site the agency manages for
// this client, and says so when the day's allowance is used.
func (c *Client) InspectURL(ctx context.Context, in InspectURLInput) (json.RawMessage, error) {
	return c.Call(ctx, "inspect_url", in)
}

// IssueIndexnowKeyInput is the input to issue_indexnow_key.
type IssueIndexnowKeyInput struct {
	ClientID string `json:"clientId"`
}

// IssueIndexnowKey calls issue_indexnow_key: Issue an IndexNow key. Issues the
// client's IndexNow key and returns the file to put on the site: fileName at
// the site's root, containing exactly the key. Answers 'set up IndexNow'.
// Re-issuing replaces the key and invalidates the file already on the site, so
// only do it when asked to or when the key must be retired. Hand the file to
// whoever can publish to the site (sites running the Seolyticize WordPress
// plugin serve it themselves), then call verify_indexnow_key.
func (c *Client) IssueIndexnowKey(ctx context.Context, in IssueIndexnowKeyInput) (json.RawMessage, error) {
	return c.Call(ctx, "issue_indexnow_key", in)
}

// IssueSiteVerificationInput is the input to issue_site_verification.
type IssueSiteVerificationInput struct {
	ClientID string `json:"clientId"`
	Method   string `json:"method,omitempty"`
	SiteURL  string `json:"siteUrl,omitempty"` // A URL-prefix site other than https://<domain>/, such as https://www.example.com/. Not used for dns.
}

// IssueSiteVerification calls issue_site_verification: Issue a Google site
// verification. Answers 'prove we own the site to Google without anyone
// opening Search Console'. Asks Google for a verification token for the
// agency's own Google account and returns what to deploy: a file at the site
// root (file), a <head> tag (meta) or a TXT record (dns). Then call
// verify_site_ownership. onboard_site does this for you. Free.
func (c *Client) IssueSiteVerification(ctx context.Context, in IssueSiteVerificationInput) (json.RawMessage, error) {
	return c.Call(ctx, "issue_site_verification", in)
}

// KeywordExplorerInput is the input to keyword_explorer.
type KeywordExplorerInput struct {
	ClientID      string         `json:"clientId"`
	Seed          string         `json:"seed"`
	Mode          string         `json:"mode,omitempty"`
	LocationCode  *int64         `json:"locationCode,omitempty"`
	LanguageCode  string         `json:"languageCode,omitempty"`
	Filters       map[string]any `json:"filters,omitempty"`
	Sort          map[string]any `json:"sort,omitempty"`
	Page          *int64         `json:"page,omitempty"`
	PageSize      *int64         `json:"pageSize,omitempty"`
	Format        string         `json:"format,omitempty"`
	MaxCostMicros *int64         `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// KeywordExplorer calls keyword_explorer: Explore keyword ideas with filters.
// Expands a seed term into keyword ideas, suggestions, related terms or
// questions, then filters, sorts and pages them: include and exclude words,
// volume, difficulty, cost per click and word-count ranges, questions only,
// search intent, and result-page features. Answers 'which terms around this
// topic are worth going after'. The expansion is bought once per seed, mode
// and market; every later filter on the same seed is free. Rows a filter
// cannot judge because the number is unknown are left out and counted in
// unknownExcluded rather than treated as zero. Set format to csv for every
// matching row as a CSV file.
func (c *Client) KeywordExplorer(ctx context.Context, in KeywordExplorerInput) (json.RawMessage, error) {
	return c.Call(ctx, "keyword_explorer", in)
}

// KeywordTrafficPotentialInput is the input to keyword_traffic_potential.
type KeywordTrafficPotentialInput struct {
	ClientID      string `json:"clientId"`
	Term          string `json:"term"`
	LocationCode  *int64 `json:"locationCode,omitempty"`
	LanguageCode  string `json:"languageCode,omitempty"`
	Refresh       *bool  `json:"refresh,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// KeywordTrafficPotential calls keyword_traffic_potential: What ranking first
// for a keyword is worth. Traffic potential and parent topic for a keyword:
// the estimated monthly organic traffic of the page ranking first, summed
// across every keyword that page ranks for, and that page's highest-volume
// keyword. Answers 'how much traffic is this keyword really worth' and 'what
// topic should the page be about' — a keyword's own volume undersells it.
// Estimated through the measured click curve and labelled with its source.
// Null when the leader's rankings are unknown, never zero. May buy the results
// page and the leader's ranked keywords when the archive holds neither.
func (c *Client) KeywordTrafficPotential(ctx context.Context, in KeywordTrafficPotentialInput) (json.RawMessage, error) {
	return c.Call(ctx, "keyword_traffic_potential", in)
}

// ListAlertRulesInput is the input to list_alert_rules.
type ListAlertRulesInput struct {
	ClientID string `json:"clientId"`
}

// ListAlertRules calls list_alert_rules: List alert rules. What a client's
// alerts watch for, at what threshold, and whether each is paused. Lists the
// rules; list_recent_alerts lists what they fired.
func (c *Client) ListAlertRules(ctx context.Context, in ListAlertRulesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_alert_rules", in)
}

// ListAPIActivityInput is the input to list_api_activity.
type ListAPIActivityInput struct {
	SinceHours *int64 `json:"sinceHours,omitempty"`
	KeyID      string `json:"keyId,omitempty"`
	Limit      *int64 `json:"limit,omitempty"`
}

// ListAPIActivity calls list_api_activity: List API activity. Recent tool
// calls made with the agency's keys — which key, which tool, which client, how
// it ended and how long it took — newest first. Lists who did what through the
// API and MCP. Inputs are never recorded.
func (c *Client) ListAPIActivity(ctx context.Context, in ListAPIActivityInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_api_activity", in)
}

// ListAPIKeysInput is the input to list_api_keys.
type ListAPIKeysInput struct {
}

// ListAPIKeys calls list_api_keys: List API keys. The agency's live API keys
// with name, role, which client each is pinned to and when it was last used.
// Lists keys to audit or revoke. Never returns a key.
func (c *Client) ListAPIKeys(ctx context.Context, in ListAPIKeysInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_api_keys", in)
}

// ListClientsInput is the input to list_clients.
type ListClientsInput struct {
}

// ListClients calls list_clients: List clients. Lists every client the caller
// can access. Use this first when a request names a client by name rather than
// id.
func (c *Client) ListClients(ctx context.Context, in ListClientsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_clients", in)
}

// ListConnectedAppsInput is the input to list_connected_apps.
type ListConnectedAppsInput struct {
}

// ListConnectedApps calls list_connected_apps: List connected apps. Assistants
// and agent apps connected to the agency by signing in (MCP OAuth) rather than
// with an API key: which app, the role and client it was given, who approved
// it, and when it was last used. Answers 'what has access to our account'.
func (c *Client) ListConnectedApps(ctx context.Context, in ListConnectedAppsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_connected_apps", in)
}

// ListCrawlIssuesInput is the input to list_crawl_issues.
type ListCrawlIssuesInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"`
}

// ListCrawlIssues calls list_crawl_issues: List site audit issues. Technical
// issues one crawl found, by rule: its readable label, severity, category, how
// many pages it affects, and howToFix — concrete steps to fix it and check the
// fix. Lists what to fix. Use crawlId 'latest' for the newest completed crawl.
func (c *Client) ListCrawlIssues(ctx context.Context, in ListCrawlIssuesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_crawl_issues", in)
}

// ListCrawlsInput is the input to list_crawls.
type ListCrawlsInput struct {
	ClientID string `json:"clientId"`
	Limit    *int64 `json:"limit,omitempty"`
}

// ListCrawls calls list_crawls: List crawls. A client's recent crawls of every
// status, newest first. Lists what the crawler has been doing, including
// crawls that failed.
func (c *Client) ListCrawls(ctx context.Context, in ListCrawlsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_crawls", in)
}

// ListInviteCodesInput is the input to list_invite_codes.
type ListInviteCodesInput struct {
}

// ListInviteCodes calls list_invite_codes: List invite codes. Lists the
// agency's invite codes — label, uses left, expiry, whether revoked — never
// the codes themselves. Answers 'which invite codes are out there'. Free.
func (c *Client) ListInviteCodes(ctx context.Context, in ListInviteCodesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_invite_codes", in)
}

// ListIssuePagesInput is the input to list_issue_pages.
type ListIssuePagesInput struct {
	ClientID string `json:"clientId"`
	RuleID   string `json:"ruleId"`
	CrawlID  string `json:"crawlId,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	Offset   *int64 `json:"offset,omitempty"`
}

// ListIssuePages calls list_issue_pages: List the pages an issue is on. Lists
// the pages where one crawl issue was found, with the message for each page,
// so they can be fixed. Answers 'which pages have missing titles' or 'where
// are the broken links'. ruleId comes from get_audit_summary or
// list_crawl_issues; crawlId defaults to the latest completed crawl. Paged:
// limit (up to 500) and offset; total is the whole count. Free.
func (c *Client) ListIssuePages(ctx context.Context, in ListIssuePagesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_issue_pages", in)
}

// ListKeywordListsInput is the input to list_keyword_lists.
type ListKeywordListsInput struct {
	ClientID string `json:"clientId,omitempty"`
}

// ListKeywordLists calls list_keyword_lists: List keyword lists. The keyword
// lists the caller can see, with how many keywords each holds — one client's,
// or every list the caller may reach when no client is named. Answers 'where
// did we save those keywords'. The agency's own lists, which belong to no
// client, are only visible to people who can see every client.
func (c *Client) ListKeywordLists(ctx context.Context, in ListKeywordListsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_keyword_lists", in)
}

// ListLinkProspectsInput is the input to list_link_prospects.
type ListLinkProspectsInput struct {
	ClientID string `json:"clientId"`
}

// ListLinkProspects calls list_link_prospects: List a client's link prospects.
// The client's outreach list: sites worth asking for a link, each with its
// stage (identified, contacted, replied, won, rejected), why it was added, and
// when it last moved, most recently moved first. Answers 'where is our link
// outreach'. Free.
func (c *Client) ListLinkProspects(ctx context.Context, in ListLinkProspectsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_link_prospects", in)
}

// ListLocalLocationsInput is the input to list_local_locations.
type ListLocalLocationsInput struct {
	ClientID string `json:"clientId"`
}

// ListLocalLocations calls list_local_locations: List a client's business
// locations. The client's business locations for local SEO: id, name, latitude
// and longitude. Answers 'which locations do we track in the map pack'. Pass
// an id to get_local_visibility, run_local_grid_scan or sync_local_reviews.
// Free.
func (c *Client) ListLocalLocations(ctx context.Context, in ListLocalLocationsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_local_locations", in)
}

// ListMarketsInput is the input to list_markets.
type ListMarketsInput struct {
	Query string `json:"query,omitempty"`
}

// ListMarkets calls list_markets: Find a market's location code. Lists the
// countries and regions this product measures, with the location code other
// tools take (locationCode), the ISO code and Google's own country code.
// Search by name, ISO code (gb), Google code (uk) or location code (2826); an
// exact code wins over a name match. Answers 'what is the location code for
// Germany'. Only markets the product can measure honestly are listed. Free.
func (c *Client) ListMarkets(ctx context.Context, in ListMarketsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_markets", in)
}

// ListPortalLinksInput is the input to list_portal_links.
type ListPortalLinksInput struct {
	ClientID string `json:"clientId"`
}

// ListPortalLinks calls list_portal_links: List client portal links. The live
// white-label portal links for a client: label, when each was made, and when
// it was last opened. Lists who can currently see the client's results without
// an account. The addresses themselves are not shown — they exist only at the
// moment a link is made.
func (c *Client) ListPortalLinks(ctx context.Context, in ListPortalLinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_portal_links", in)
}

// ListProjectsInput is the input to list_projects.
type ListProjectsInput struct {
	ClientID string `json:"clientId"`
}

// ListProjects calls list_projects: List projects and sites. A client's
// projects with the sites under each. Lists where crawls and trackers hang
// from.
func (c *Client) ListProjects(ctx context.Context, in ListProjectsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_projects", in)
}

// ListRankTrackersInput is the input to list_rank_trackers.
type ListRankTrackersInput struct {
	ClientID string `json:"clientId"`
}

// ListRankTrackers calls list_rank_trackers: List rank trackers. A client's
// rank trackers with engine, device, cadence and locale. Lists the tracker ids
// the ranking tools take.
func (c *Client) ListRankTrackers(ctx context.Context, in ListRankTrackersInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_rank_trackers", in)
}

// ListRecentAlertsInput is the input to list_recent_alerts.
type ListRecentAlertsInput struct {
	ClientID string `json:"clientId"`
	Since    string `json:"since"`
}

// ListRecentAlerts calls list_recent_alerts: List recent alerts. Alerts raised
// for a client since a given time, newest first. Lists what has fired, for
// anything watching — a chat channel, an automation, a status board.
func (c *Client) ListRecentAlerts(ctx context.Context, in ListRecentAlertsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_recent_alerts", in)
}

// ListReportSchedulesInput is the input to list_report_schedules.
type ListReportSchedulesInput struct {
	ClientID string `json:"clientId"`
}

// ListReportSchedules calls list_report_schedules: List report schedules.
// Every schedule that emails one of a client's saved reports: which report,
// monthly or weekly, the recipients, whether it is paused and when it last
// sent. Lists who receives what and how often. send_report sends to these
// recipients, so a report with no schedule has nobody to send to.
func (c *Client) ListReportSchedules(ctx context.Context, in ListReportSchedulesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_report_schedules", in)
}

// ListReportTemplatesInput is the input to list_report_templates.
type ListReportTemplatesInput struct {
	ClientID string `json:"clientId,omitempty"`
}

// ListReportTemplates calls list_report_templates: List report templates.
// Lists the report templates that apply: the agency's own, plus the client's
// when clientId is given — each with its audience, sections in order with
// guidance, tone and sign-off. Answers 'how do we usually write this report'.
// Follow a template's sections and tone when writing a report it fits, and
// name it in save_written_report. Free.
func (c *Client) ListReportTemplates(ctx context.Context, in ListReportTemplatesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_report_templates", in)
}

// ListReportsInput is the input to list_reports.
type ListReportsInput struct {
	ClientID string `json:"clientId"`
}

// ListReports calls list_reports: List reports. A client's saved reports and
// their recent deliveries, including failed ones with the reason. Lists what
// goes out and answers 'did my report go out'.
func (c *Client) ListReports(ctx context.Context, in ListReportsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_reports", in)
}

// ListSitemapsInput is the input to list_sitemaps.
type ListSitemapsInput struct {
	ClientID string `json:"clientId"`
}

// ListSitemaps calls list_sitemaps: The sitemaps each search engine knows
// about. The client's sitemaps as Google, Bing and Yandex each report them:
// submitted URL counts, errors and warnings, when each was last read. Answers
// 'are our sitemaps being read, and are they broken'. Google no longer reports
// how many of a sitemap's URLs are indexed, so that figure is always null;
// instead, of the pages inspected, how many came from each sitemap and passed.
// Each engine's part is separately present, refused or failed.
func (c *Client) ListSitemaps(ctx context.Context, in ListSitemapsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_sitemaps", in)
}

// ListTeamMembersInput is the input to list_team_members.
type ListTeamMembersInput struct {
}

// ListTeamMembers calls list_team_members: List the team. Everyone with access
// to the agency: email, name, role, and the clients each is limited to (empty
// means agency-wide). Also lists the clients a new member could be limited to.
// Answers 'who can see this account'.
func (c *Client) ListTeamMembers(ctx context.Context, in ListTeamMembersInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_team_members", in)
}

// ListTrackedKeywordsInput is the input to list_tracked_keywords.
type ListTrackedKeywordsInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId"`
}

// ListTrackedKeywords calls list_tracked_keywords: List tracked keywords. The
// terms a tracker measures and the competitor domains it compares against.
// Lists what is being paid for on every run.
func (c *Client) ListTrackedKeywords(ctx context.Context, in ListTrackedKeywordsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_tracked_keywords", in)
}

// ListWrittenReportsInput is the input to list_written_reports.
type ListWrittenReportsInput struct {
	ClientID string `json:"clientId"`
}

// ListWrittenReports calls list_written_reports: List written reports. Lists
// the written reports saved for a client, newest change first: id, title, a
// plain-text summary, who wrote and last changed it, and whether it is shared.
// Answers 'what reports have we written for this client'. Call it before
// save_written_report to find a report to replace rather than writing a
// duplicate. Free.
func (c *Client) ListWrittenReports(ctx context.Context, in ListWrittenReportsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_written_reports", in)
}

// OnboardSiteInput is the input to onboard_site.
type OnboardSiteInput struct {
	Domain       string   `json:"domain"`
	Name         string   `json:"name,omitempty"`
	Framework    string   `json:"framework,omitempty"`
	Keywords     []string `json:"keywords,omitempty"`
	Verification string   `json:"verification,omitempty"` // file (a page at the site root, the default), meta (a tag in the <head>) or dns (a TXT record, which also covers every subdomain and protocol).
}

// OnboardSite calls onboard_site: Onboard a site in one call. Answers 'set
// this site up'. Makes the client, project, site and rank tracker — or finds
// the client already here for this domain, so calling it twice is safe —
// issues the IndexNow key and a Google site verification for the agency's own
// Google account, sets up cookieless analytics, and returns one deploy
// manifest: files to serve at exact paths, tags for the <head>, the analytics
// snippet placed for the framework given, and CSP entries. Deploy all of it,
// then call verify_site. No Google sign-in or Search Console settings are
// needed at any point. Free.
func (c *Client) OnboardSite(ctx context.Context, in OnboardSiteInput) (json.RawMessage, error) {
	return c.Call(ctx, "onboard_site", in)
}

// PageHeadersInput is the input to page_headers.
type PageHeadersInput struct {
	URL           string `json:"url"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// PageHeaders calls page_headers: What the server actually returned. Status
// code, redirect chain, content type and the headers that decide indexing for
// one address. Answers 'why is this page not in the index' when the markup
// looks fine — an X-Robots-Tag noindex is invisible to anything reading the
// page.
func (c *Client) PageHeaders(ctx context.Context, in PageHeadersInput) (json.RawMessage, error) {
	return c.Call(ctx, "page_headers", in)
}

// PageLabInput is the input to page_lab.
type PageLabInput struct {
	URL           string `json:"url"`
	Strategy      string `json:"strategy,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// PageLab calls page_lab: How a page scores in one synthetic run. Lighthouse
// performance, SEO, accessibility and best-practices scores for one address,
// from PageSpeed Insights. Answers 'what is slow or wrong on this page' with a
// reproducible lab run, including for pages too quiet for field data. This is
// not what real visitors experienced — use page_vitals for that, and say which
// is which when reporting both. There is no interaction metric here: a
// synthetic run interacts with nothing.
func (c *Client) PageLab(ctx context.Context, in PageLabInput) (json.RawMessage, error) {
	return c.Call(ctx, "page_lab", in)
}

// PageVitalsInput is the input to page_vitals.
type PageVitalsInput struct {
	URL string `json:"url"`
}

// PageVitals calls page_vitals: How a page performs for real visitors. Largest
// Contentful Paint, Interaction to Next Paint and Cumulative Layout Shift for
// one address, from the Chrome UX Report. Answers 'is this page fast enough
// for real visitors' with field data over a trailing 28 days, which is what
// Google ranks on rather than a lab score.
func (c *Client) PageVitals(ctx context.Context, in PageVitalsInput) (json.RawMessage, error) {
	return c.Call(ctx, "page_vitals", in)
}

// PinTrackedKeywordsInput is the input to pin_tracked_keywords.
type PinTrackedKeywordsInput struct {
	ClientID   string   `json:"clientId"`
	TrackerID  string   `json:"trackerId"`
	KeywordIDs []string `json:"keywordIds"`
	Pinned     *bool    `json:"pinned,omitempty"`
}

// PinTrackedKeywords calls pin_tracked_keywords: Pin tracked keywords to the
// top. Pins tracked keywords (by keywordId from list_tracked_keywords or
// get_rank_history) to the top of a tracker's table for everyone on the
// account, or unpins them with pinned false. Pinned keywords are listed first,
// in the order they were pinned. Ids not on the tracker are ignored; changed
// says how many moved. Answers 'keep these at the top'. Free.
func (c *Client) PinTrackedKeywords(ctx context.Context, in PinTrackedKeywordsInput) (json.RawMessage, error) {
	return c.Call(ctx, "pin_tracked_keywords", in)
}

// QueueURLInspectionsInput is the input to queue_url_inspections.
type QueueURLInspectionsInput struct {
	ClientID string   `json:"clientId"`
	URLs     []string `json:"urls"`
}

// QueueURLInspections calls queue_url_inspections: Queue URLs for Search
// Console inspection. Queues up to 500 of the client's URLs for Search
// Console's URL inspection, which runs in the background within Google's daily
// allowance. Answers 'check whether all these pages are indexed'. URLs that
// are not on the client's own site are refused and listed with the reason.
// Read the results with get_index_coverage as they arrive; inspect_url checks
// one URL immediately. Needs Search Console connected (see
// get_connection_link).
func (c *Client) QueueURLInspections(ctx context.Context, in QueueURLInspectionsInput) (json.RawMessage, error) {
	return c.Call(ctx, "queue_url_inspections", in)
}

// RecheckLinksInput is the input to recheck_links.
type RecheckLinksInput struct {
	ClientID string `json:"clientId"`
}

// RecheckLinks calls recheck_links: Re-check the pages that link to a client.
// Sends our crawler to the pages said to link to the client to see whether
// each still does — slowly and politely, so it takes a while; read the result
// with get_link_checks. A second request within the hour lands on the check
// already running. No vendor charge, but it visits other people's sites on the
// client's behalf, so it needs client:manage. Refused when the client has no
// rank tracker yet. Answers 'check our links again'.
func (c *Client) RecheckLinks(ctx context.Context, in RecheckLinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "recheck_links", in)
}

// RefreshDomainDataInput is the input to refresh_domain_data.
type RefreshDomainDataInput struct {
	ClientID      string `json:"clientId"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// RefreshDomainData calls refresh_domain_data: Refresh domain and backlink
// data now. Refreshes everything bought about the client's domain — overview,
// ranked keywords, competitors, top pages and the backlink profile — plus what
// up to five tracked competitors rank for and who links to them. Answers
// 'update the backlinks and keyword data'. The dearest thing the product does:
// quote it with estimate_cost first; with maxCostMicros the refresh is refused
// before it starts when its estimate is over the limit. Runs in the background
// for a few minutes: poll get_run_status with job 'domain'. A second request
// within the hour lands on the refresh already going.
func (c *Client) RefreshDomainData(ctx context.Context, in RefreshDomainDataInput) (json.RawMessage, error) {
	return c.Call(ctx, "refresh_domain_data", in)
}

// RefreshRankingsInput is the input to refresh_rankings.
type RefreshRankingsInput struct {
	ClientID      string `json:"clientId"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// RefreshRankings calls refresh_rankings: Capture rankings now. Captures the
// client's rank tracker now instead of at its next scheduled run: every
// tracked keyword on every device the tracker reads, each a paid lookup.
// Answers 'get me today's positions'. Quote it with estimate_cost first; with
// maxCostMicros the run is refused before it starts when its estimate is over
// the limit. Runs in the background: poll get_run_status with job 'rank', then
// read get_rankings or get_rank_summary. A second request within the same
// minute lands on the run already going rather than paying twice.
func (c *Client) RefreshRankings(ctx context.Context, in RefreshRankingsInput) (json.RawMessage, error) {
	return c.Call(ctx, "refresh_rankings", in)
}

// RemoveCompetitorsInput is the input to remove_competitors.
type RemoveCompetitorsInput struct {
	ClientID  string   `json:"clientId"`
	TrackerID string   `json:"trackerId"`
	Domains   []string `json:"domains"` // Domains such as example.com. A URL is reduced to its domain.
}

// RemoveCompetitors calls remove_competitors: Remove competitors from a
// tracker. Stops measuring the named competitor domains on a rank tracker.
// Answers 'we no longer compete with this site'. Returns the tracker's
// remaining competitors. Naming a domain that is not on the list is not an
// error.
func (c *Client) RemoveCompetitors(ctx context.Context, in RemoveCompetitorsInput) (json.RawMessage, error) {
	return c.Call(ctx, "remove_competitors", in)
}

// RemoveKeywordsFromListInput is the input to remove_keywords_from_list.
type RemoveKeywordsFromListInput struct {
	ListID     string   `json:"listId"`
	KeywordIDs []string `json:"keywordIds"`
}

// RemoveKeywordsFromList calls remove_keywords_from_list: Remove keywords from
// a list. Removes keywords from a list by keyword id. Tracked keywords that
// came from the list stay tracked. Answers 'take these off that list'.
func (c *Client) RemoveKeywordsFromList(ctx context.Context, in RemoveKeywordsFromListInput) (json.RawMessage, error) {
	return c.Call(ctx, "remove_keywords_from_list", in)
}

// RemoveTeamMemberInput is the input to remove_team_member.
type RemoveTeamMemberInput struct {
	MembershipID string `json:"membershipId"`
}

// RemoveTeamMember calls remove_team_member: Remove someone from the team.
// Ends a member's access to the agency, by the membershipId list_team_members
// gives. Answers 'Sam has left'. Their account itself is untouched. The last
// owner cannot be removed. Safe to repeat.
func (c *Client) RemoveTeamMember(ctx context.Context, in RemoveTeamMemberInput) (json.RawMessage, error) {
	return c.Call(ctx, "remove_team_member", in)
}

// RemoveTrackedKeywordsInput is the input to remove_tracked_keywords.
type RemoveTrackedKeywordsInput struct {
	ClientID  string   `json:"clientId"`
	TrackerID string   `json:"trackerId"`
	Terms     []string `json:"terms"`
}

// RemoveTrackedKeywords calls remove_tracked_keywords: Stop tracking keywords.
// Removes terms from a tracker. Answers 'stop tracking these keywords'. Terms
// that were not tracked are ignored, so it is safe to repeat.
func (c *Client) RemoveTrackedKeywords(ctx context.Context, in RemoveTrackedKeywordsInput) (json.RawMessage, error) {
	return c.Call(ctx, "remove_tracked_keywords", in)
}

// RemoveVendorCredentialInput is the input to remove_vendor_credential.
type RemoveVendorCredentialInput struct {
	Provider string `json:"provider"`
}

// RemoveVendorCredential calls remove_vendor_credential: Disconnect a vendor
// account. Deletes the agency's saved account for one vendor. Answers 'stop
// using our DataForSEO account'. Anything that relied on it — keyword data
// from DataForSEO, an assistant's answers — stops until another is saved. Safe
// to repeat.
func (c *Client) RemoveVendorCredential(ctx context.Context, in RemoveVendorCredentialInput) (json.RawMessage, error) {
	return c.Call(ctx, "remove_vendor_credential", in)
}

// RenameKeywordListInput is the input to rename_keyword_list.
type RenameKeywordListInput struct {
	ListID string `json:"listId"`
	Name   string `json:"name"`
}

// RenameKeywordList calls rename_keyword_list: Rename a keyword list. Gives a
// keyword list a new name. Answers 'call that list something else'. Refused
// when another list for the same client (or the agency) already has the name.
func (c *Client) RenameKeywordList(ctx context.Context, in RenameKeywordListInput) (json.RawMessage, error) {
	return c.Call(ctx, "rename_keyword_list", in)
}

// RevokeAPIKeyInput is the input to revoke_api_key.
type RevokeAPIKeyInput struct {
	KeyID string `json:"keyId"`
}

// RevokeAPIKey calls revoke_api_key: Revoke an API key. Stops a key working,
// immediately and for good. Answers 'this key leaked'. Safe to repeat;
// revoking the key making the call is allowed and takes effect on its next
// request.
func (c *Client) RevokeAPIKey(ctx context.Context, in RevokeAPIKeyInput) (json.RawMessage, error) {
	return c.Call(ctx, "revoke_api_key", in)
}

// RevokeDeployHookInput is the input to revoke_deploy_hook.
type RevokeDeployHookInput struct {
	ClientID string `json:"clientId"`
}

// RevokeDeployHook calls revoke_deploy_hook: Revoke the deploy hook. Answers
// 'that deploy hook URL leaked'. Stops the client's deploy hook URL working at
// once. Free.
func (c *Client) RevokeDeployHook(ctx context.Context, in RevokeDeployHookInput) (json.RawMessage, error) {
	return c.Call(ctx, "revoke_deploy_hook", in)
}

// RevokeInviteCodeInput is the input to revoke_invite_code.
type RevokeInviteCodeInput struct {
	InviteCodeID string `json:"inviteCodeId"`
}

// RevokeInviteCode calls revoke_invite_code: Revoke an invite code. Stops an
// invite code being claimed again. Answers 'that code was shared too widely'.
// Keys already made from it keep working until revoked with revoke_api_key.
// Free.
func (c *Client) RevokeInviteCode(ctx context.Context, in RevokeInviteCodeInput) (json.RawMessage, error) {
	return c.Call(ctx, "revoke_invite_code", in)
}

// RevokePortalLinkInput is the input to revoke_portal_link.
type RevokePortalLinkInput struct {
	ClientID string `json:"clientId"`
	LinkID   string `json:"linkId"`
}

// RevokePortalLink calls revoke_portal_link: Revoke a client portal link.
// Stops a portal link working, immediately and for good. Answers 'that link
// was forwarded' or 'the engagement ended'. A link that is already revoked, or
// is not this client's, is refused rather than reported as done.
func (c *Client) RevokePortalLink(ctx context.Context, in RevokePortalLinkInput) (json.RawMessage, error) {
	return c.Call(ctx, "revoke_portal_link", in)
}

// RunAIVisibilityCheckInput is the input to run_ai_visibility_check.
type RunAIVisibilityCheckInput struct {
	ClientID      string `json:"clientId"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// RunAIVisibilityCheck calls run_ai_visibility_check: Run an AI visibility
// check now. Asks every tracked question on the chosen assistants now, in the
// background; results arrive within minutes in get_ai_visibility_runs and
// get_ai_visibility_results. The schedule is unchanged. With maxCostMicros,
// the check is refused before it starts when its estimate is over the limit.
// Get the user's go-ahead first: it spends. Answers 'check it now'.
func (c *Client) RunAIVisibilityCheck(ctx context.Context, in RunAIVisibilityCheckInput) (json.RawMessage, error) {
	return c.Call(ctx, "run_ai_visibility_check", in)
}

// RunLocalGridScanInput is the input to run_local_grid_scan.
type RunLocalGridScanInput struct {
	ClientID      string   `json:"clientId"`
	LocationID    string   `json:"locationId"`
	Term          string   `json:"term"`
	Size          *float64 `json:"size,omitempty"`          // Grid width in points: 3, 5 or 7.
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// RunLocalGridScan calls run_local_grid_scan: Run a map-pack grid scan.
// Searches the term at every point of a size × size grid around the location
// (1.5 km apart) and records where the business appears in the map pack. Every
// point is a paid lookup, so a 7 × 7 grid is 49 lookups: quote it with
// estimate_cost and get the user's go-ahead first. Runs in the background;
// read the result with get_local_visibility a few minutes later. With
// maxCostMicros, the scan is refused before it starts when its estimate is
// over the limit. Answers 'scan our map-pack visibility now'.
func (c *Client) RunLocalGridScan(ctx context.Context, in RunLocalGridScanInput) (json.RawMessage, error) {
	return c.Call(ctx, "run_local_grid_scan", in)
}

// SaveAIVisibilitySetupInput is the input to save_ai_visibility_setup.
type SaveAIVisibilitySetupInput struct {
	ClientID        string           `json:"clientId"`
	AddPrompts      []map[string]any `json:"addPrompts,omitempty"`
	RetirePromptIDs []string         `json:"retirePromptIds,omitempty"`
	AddBrands       []map[string]any `json:"addBrands,omitempty"`
}

// SaveAIVisibilitySetup calls save_ai_visibility_setup: Change what AI
// visibility tracks. Adds questions (each optionally with a two-letter
// market), retires questions by id (past answers keep their question), and
// adds brands to look for (name, domain, aliases; isClient for the client's
// own — there should be exactly one). Questions already tracked are skipped.
// Write questions the way a buyer would ask an assistant, without naming any
// brand; suggest_ai_prompts drafts them. Answers 'track these questions'.
// Free; the checks that follow are what cost.
func (c *Client) SaveAIVisibilitySetup(ctx context.Context, in SaveAIVisibilitySetupInput) (json.RawMessage, error) {
	return c.Call(ctx, "save_ai_visibility_setup", in)
}

// SaveBaiduTokenInput is the input to save_baidu_token.
type SaveBaiduTokenInput struct {
	ClientID string `json:"clientId"`
	SiteURL  string `json:"siteUrl"`
	Token    string `json:"token"`
}

// SaveBaiduToken calls save_baidu_token: Save the client's Baidu push token.
// Saves the link-push token from Baidu Ziyuan (ziyuan.baidu.com, link
// submission page) for the client's site, so submit_baidu_urls can push pages.
// Answers 'connect Baidu for this client'. The site must be the client's own
// registered site; the token is stored encrypted and never returned. Replaces
// any token already saved.
func (c *Client) SaveBaiduToken(ctx context.Context, in SaveBaiduTokenInput) (json.RawMessage, error) {
	return c.Call(ctx, "save_baidu_token", in)
}

// SaveReportTemplateInput is the input to save_report_template.
type SaveReportTemplateInput struct {
	TemplateID  string           `json:"templateId,omitempty"`
	ClientID    string           `json:"clientId,omitempty"`
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Audience    string           `json:"audience,omitempty"`
	Sections    []map[string]any `json:"sections"`
	Tone        string           `json:"tone,omitempty"`
	SignOff     string           `json:"signOff,omitempty"`
}

// SaveReportTemplate calls save_report_template: Save a report template. Saves
// a reusable brief for a kind of report: a name, who it is for, sections in
// order (each with a line on what goes under it), tone and sign-off. Answers
// 'make this our monthly report format'. Without clientId it is agency-wide,
// which only someone who sees every client may change; with one it is that
// client's own. Without templateId a new template is made; with one, it is
// replaced. Only save a template when the user asks for one — a one-off report
// needs none. Free.
func (c *Client) SaveReportTemplate(ctx context.Context, in SaveReportTemplateInput) (json.RawMessage, error) {
	return c.Call(ctx, "save_report_template", in)
}

// SaveVendorCredentialInput is the input to save_vendor_credential.
type SaveVendorCredentialInput struct {
	Provider  string         `json:"provider"`
	Login     string         `json:"login,omitempty"`
	Secret    string         `json:"secret"`
	GoogleAds map[string]any `json:"googleAds,omitempty"`
}

// SaveVendorCredential calls save_vendor_credential: Connect a vendor account.
// Stores the agency's own account for a data vendor or AI assistant —
// DataForSEO (login + password), SerpApi, a proxy pool (serpproxy: one URL per
// line), Google Ads (six values in googleAds), an Anthropic, OpenAI, Gemini or
// Perplexity key, and the rest get_vendor_connections names — and checks it
// works straight away. Answers 'connect our DataForSEO account'. Replaces any
// key already saved for that vendor. The response says whether the check
// passed; it never contains the key.
func (c *Client) SaveVendorCredential(ctx context.Context, in SaveVendorCredentialInput) (json.RawMessage, error) {
	return c.Call(ctx, "save_vendor_credential", in)
}

// SaveWrittenReportInput is the input to save_written_report.
type SaveWrittenReportInput struct {
	ClientID   string `json:"clientId"`
	ReportID   string `json:"reportId,omitempty"`
	Title      string `json:"title"`
	Markdown   string `json:"markdown"`
	TemplateID string `json:"templateId,omitempty"`
}

// SaveWrittenReport calls save_written_report: Save a written report. Saves a
// finished report to a client, where everyone on the account can read and
// print it. Answers 'save this report for the client'. Write it in Markdown
// (headings, lists, tables, links); raw HTML is shown as text, not run.
// Without reportId a new report is created; with one, that report's title and
// text are replaced and its share link, if any, keeps working. Name templateId
// when it was written to a report template. New reports are private — use
// set_report_sharing only when the user asks for a link. Free.
func (c *Client) SaveWrittenReport(ctx context.Context, in SaveWrittenReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "save_written_report", in)
}

// ScheduleReportInput is the input to schedule_report.
type ScheduleReportInput struct {
	ClientID   string   `json:"clientId"`
	ReportID   string   `json:"reportId"`
	Cadence    string   `json:"cadence,omitempty"`
	Recipients []string `json:"recipients"` // Email addresses the report goes to. At least one.
}

// ScheduleReport calls schedule_report: Schedule a report. Puts one of a
// client's saved reports (see list_reports, create_report) on a monthly or
// weekly cadence for a list of recipients. Answers 'send this report to the
// client every month'. The scheduled send runs on its own; send_report sends
// one now, to every recipient of every schedule for that report. A report can
// have more than one schedule.
func (c *Client) ScheduleReport(ctx context.Context, in ScheduleReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "schedule_report", in)
}

// SendReportInput is the input to send_report.
type SendReportInput struct {
	ClientID       string `json:"clientId"`
	ReportID       string `json:"reportId"`
	IdempotencyKey string `json:"idempotencyKey,omitempty"`
}

// SendReport calls send_report: Send a report now. Sends a saved report to its
// scheduled recipients now, covering the month just gone. Answers 'send the
// report'. Recipients come from the report's schedule and cannot be named
// here. Pass an idempotencyKey to make a retry land on the same send instead
// of emailing the client twice.
func (c *Client) SendReport(ctx context.Context, in SendReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "send_report", in)
}

// SerpOverviewInput is the input to serp_overview.
type SerpOverviewInput struct {
	ClientID      string `json:"clientId"`
	Term          string `json:"term"`
	LocationCode  *int64 `json:"locationCode,omitempty"`
	LanguageCode  string `json:"languageCode,omitempty"`
	Refresh       *bool  `json:"refresh,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// SerpOverview calls serp_overview: Who is on the first page for a keyword.
// The current top ten organic results for a keyword, each with domain
// authority, referring domains, the estimated organic traffic of that page,
// links into the page where counted, and every feature on the results page —
// plus why the keyword's difficulty is what it is. Answers 'could we get onto
// that page, and who would we have to beat'. Read from a recent capture when
// one is held, bought otherwise. Each figure says where it came from (our link
// graph or the vendor; panel or published click curve), and anything unknown
// is null, never zero.
func (c *Client) SerpOverview(ctx context.Context, in SerpOverviewInput) (json.RawMessage, error) {
	return c.Call(ctx, "serp_overview", in)
}

// SetAIVisibilityScheduleInput is the input to set_ai_visibility_schedule.
type SetAIVisibilityScheduleInput struct {
	ClientID string   `json:"clientId"`
	Cadence  string   `json:"cadence,omitempty"`
	Engines  []string `json:"engines,omitempty"`
}

// SetAIVisibilitySchedule calls set_ai_visibility_schedule: Set how often AI
// visibility is checked. Sets a client's cadence (off, daily, weekly or
// monthly; weekly unless changed) and which assistants it is checked on
// (engines; null or empty for every one the agency has configured). Every
// scheduled check asks every question on every chosen assistant, so show
// estimate_ai_visibility_cost first. Answers 'check this daily on ChatGPT and
// Perplexity'. Free to set.
func (c *Client) SetAIVisibilitySchedule(ctx context.Context, in SetAIVisibilityScheduleInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_ai_visibility_schedule", in)
}

// SetBrandingInput is the input to set_branding.
type SetBrandingInput struct {
	Name         string `json:"name"`
	LogoURL      string `json:"logoUrl"`
	PrimaryColor string `json:"primaryColor"`
}

// SetBranding calls set_branding: Set the agency's branding. Renames the
// agency and sets the logo and colour on its reports and portal links. Answers
// 'put our logo on the reports'. The logo must be an https:// address and the
// colour six-digit hex such as #1f5ad6; empty strings clear them. Send all
// three: this replaces what is there.
func (c *Client) SetBranding(ctx context.Context, in SetBrandingInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_branding", in)
}

// SetKeywordNotesInput is the input to set_keyword_notes.
type SetKeywordNotesInput struct {
	ListID string           `json:"listId"`
	Notes  []map[string]any `json:"notes"`
}

// SetKeywordNotes calls set_keyword_notes: Note why keywords are on a list.
// Sets notes on keywords in a list — why each is there, which page it belongs
// to, what was decided — by keywordId from get_keyword_list. An empty note
// clears it; ids not on the list are ignored and changed says how many were
// set. Answers 'note that this one is for the implants page'. Free.
func (c *Client) SetKeywordNotes(ctx context.Context, in SetKeywordNotesInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_keyword_notes", in)
}

// SetLinkProspectStatusInput is the input to set_link_prospect_status.
type SetLinkProspectStatusInput struct {
	ClientID   string `json:"clientId"`
	ProspectID string `json:"prospectId"`
	Status     string `json:"status"`
}

// SetLinkProspectStatus calls set_link_prospect_status: Move a link prospect
// to another stage. Moves one prospect on the client's outreach list to
// identified, contacted, replied, won or rejected. Take the id from
// list_link_prospects. Answers 'we heard back from this site'. Free.
func (c *Client) SetLinkProspectStatus(ctx context.Context, in SetLinkProspectStatusInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_link_prospect_status", in)
}

// SetReportSharingInput is the input to set_report_sharing.
type SetReportSharingInput struct {
	ClientID string `json:"clientId"`
	ReportID string `json:"reportId"`
	Shared   bool   `json:"shared"`
}

// SetReportSharing calls set_report_sharing: Share a written report by link.
// Turns a written report's public link on or off. Answers 'give me a link I
// can send the client'. Anyone holding the link can read the report without an
// account, so only turn it on when the user explicitly asks. Turning it on for
// a report already shared returns the same link; turning it off breaks the
// link immediately, and sharing again later makes a new one. The page carries
// a preview card for Slack, email and social networks. Free.
func (c *Client) SetReportSharing(ctx context.Context, in SetReportSharingInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_report_sharing", in)
}

// SetUpWebAnalyticsInput is the input to set_up_web_analytics.
type SetUpWebAnalyticsInput struct {
	ClientID     string   `json:"clientId"`
	Hostnames    []string `json:"hostnames,omitempty"`
	ExcludePaths []string `json:"excludePaths,omitempty"` // Paths never recorded, enforced on the server whatever the snippet says: ["/app/*", "/verify", "/reset"]. "*" matches anything. Replaces the list; [] clears it.
	IncludePaths []string `json:"includePaths,omitempty"` // When set, only these paths are recorded: ["/", "/blog/*"]. [] clears it.
}

// SetUpWebAnalytics calls set_up_web_analytics: Set up web analytics. Creates
// a client's cookieless web analytics, or changes which hostnames may report
// to it, and returns the same shape as get_web_analytics_setup: the one-line
// tracking snippet to add to every page, the allowed hostnames and
// lastEventAt. Answers 'install tracking on this site'. Without hostnames it
// allows the client's own domain and its www form; with hostnames, those
// replace the list (up to 20; a visit from any other hostname is ignored, so
// add staging or extra domains here). Safe to repeat: an existing setup keeps
// its snippet. configured is false when the client has no site domain to
// default to and no hostnames were given. Free.
func (c *Client) SetUpWebAnalytics(ctx context.Context, in SetUpWebAnalyticsInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_up_web_analytics", in)
}

// StartCrawlInput is the input to start_crawl.
type StartCrawlInput struct {
	ClientID string `json:"clientId"`
}

// StartCrawl calls start_crawl: Start a site audit crawl. Crawls the client's
// site now rather than waiting for the schedule. Answers 'audit the site
// again'. If a crawl is already running it is returned instead of starting a
// second. Poll get_crawl_status with the crawl id; read results with
// list_crawl_issues.
func (c *Client) StartCrawl(ctx context.Context, in StartCrawlInput) (json.RawMessage, error) {
	return c.Call(ctx, "start_crawl", in)
}

// StrikingDistanceInput is the input to striking_distance.
type StrikingDistanceInput struct {
	ClientID       string   `json:"clientId"`
	Days           *int64   `json:"days,omitempty"`
	MinPosition    *float64 `json:"minPosition,omitempty"`
	MaxPosition    *float64 `json:"maxPosition,omitempty"`
	MinImpressions *int64   `json:"minImpressions,omitempty"`
}

// StrikingDistance calls striking_distance: Queries just off the top, by page.
// Search Console queries where the client already ranks around positions four
// to twenty with meaningful impressions, grouped by the page that ranks. Each
// has the expected click-through rate now and at the top three from the
// measured click curve, the extra monthly clicks a move to third (and to
// first) would bring, and an opportunity score that favours queries nearest
// the top. Answers 'which existing pages should we improve first'. Read from
// stored console data; buys nothing. Null when the client has no site
// configured.
func (c *Client) StrikingDistance(ctx context.Context, in StrikingDistanceInput) (json.RawMessage, error) {
	return c.Call(ctx, "striking_distance", in)
}

// SubmitBaiduURLsInput is the input to submit_baidu_urls.
type SubmitBaiduURLsInput struct {
	ClientID      string   `json:"clientId"`
	URLs          []string `json:"urls"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// SubmitBaiduURLs calls submit_baidu_urls: Push pages to Baidu. Pushes
// addresses on the client's own site to Baidu through Ziyuan's link
// submission, the only API Baidu offers a site owner. Answers 'tell Baidu
// about these pages'. Returns how many Baidu accepted and how many more it
// will take today — Baidu reports its allowance only in reply to a push.
// Refuses addresses not on the client's registered site, and refuses entirely
// when no Baidu token has been saved.
func (c *Client) SubmitBaiduURLs(ctx context.Context, in SubmitBaiduURLsInput) (json.RawMessage, error) {
	return c.Call(ctx, "submit_baidu_urls", in)
}

// SubmitIndexnowInput is the input to submit_indexnow.
type SubmitIndexnowInput struct {
	ClientID      string   `json:"clientId"`
	URLs          []string `json:"urls"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// SubmitIndexnow calls submit_indexnow: Tell search engines a page changed.
// Submits changed addresses to IndexNow, which Bing, Yandex, Seznam and Naver
// share. Google does not take part. Only accepts addresses on a site the
// agency manages — anything else is refused and named. Answers 'we have just
// republished these pages, tell the engines'. Refuses entirely until the
// client's key file has been published and checked.
func (c *Client) SubmitIndexnow(ctx context.Context, in SubmitIndexnowInput) (json.RawMessage, error) {
	return c.Call(ctx, "submit_indexnow", in)
}

// SubmitSitemapInput is the input to submit_sitemap.
type SubmitSitemapInput struct {
	ClientID   string `json:"clientId"`
	SitemapURL string `json:"sitemapUrl"`
}

// SubmitSitemap calls submit_sitemap: Submit a sitemap to Google. Submits a
// sitemap address to the client's Google Search Console property. Answers
// 'tell Google about this sitemap'. Refuses an address that is not on a site
// the agency manages for this client, Connected by the agency's Google account
// (verify_site, or the account added as a user), it submits directly — no
// further permission; Google itself refuses if that account is only a
// Restricted user, and the refusal says what to change. Connected by someone's
// Google sign-in, it needs sitemap management allowed separately (a second
// Google permission, never assumed from the read-only one).
func (c *Client) SubmitSitemap(ctx context.Context, in SubmitSitemapInput) (json.RawMessage, error) {
	return c.Call(ctx, "submit_sitemap", in)
}

// SubmitURLsInput is the input to submit_urls.
type SubmitURLsInput struct {
	ClientID      string   `json:"clientId"`
	URLs          []string `json:"urls"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// SubmitURLs calls submit_urls: Submit URLs for crawling. Adds addresses to
// the crawl frontier for a client. Only accepts addresses on a site the agency
// manages — anything else is refused and named. Answers 'crawl these pages I
// found'.
func (c *Client) SubmitURLs(ctx context.Context, in SubmitURLsInput) (json.RawMessage, error) {
	return c.Call(ctx, "submit_urls", in)
}

// SubmitYandexRecrawlInput is the input to submit_yandex_recrawl.
type SubmitYandexRecrawlInput struct {
	ClientID      string   `json:"clientId"`
	URLs          []string `json:"urls"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// SubmitYandexRecrawl calls submit_yandex_recrawl: Ask Yandex to recrawl
// pages. Asks Yandex Webmaster to recrawl pages of the client's verified
// Yandex host, one task per page. Answers 'get Yandex to look at these pages
// again'. Returns the task id Yandex gave each queued page;
// get_yandex_performance shows the queue and the day's remaining quota.
// Refuses addresses not on the verified host (Yandex treats www and the bare
// domain as different hosts). The first quota refusal stops the rest, which
// come back as not sent.
func (c *Client) SubmitYandexRecrawl(ctx context.Context, in SubmitYandexRecrawlInput) (json.RawMessage, error) {
	return c.Call(ctx, "submit_yandex_recrawl", in)
}

// SuggestAIPromptsInput is the input to suggest_ai_prompts.
type SuggestAIPromptsInput struct {
	ClientID      string `json:"clientId"`
	Topic         string `json:"topic,omitempty"`
	Count         *int64 `json:"count,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// SuggestAIPrompts calls suggest_ai_prompts: Suggest questions to track.
// Drafts questions worth tracking, grouped by topic, from the client brief,
// its brands and what is already tracked — phrased as a buyer would ask an
// assistant, never naming a tracked brand, never repeating a tracked question.
// Optionally around one topic. Nothing is saved: show them, then save the
// chosen ones with save_ai_visibility_setup. For questions people actually
// search, use keyword_explorer with mode questions. Answers 'what should we be
// tracking'. A metered Claude call.
func (c *Client) SuggestAIPrompts(ctx context.Context, in SuggestAIPromptsInput) (json.RawMessage, error) {
	return c.Call(ctx, "suggest_ai_prompts", in)
}

// SyncLocalReviewsInput is the input to sync_local_reviews.
type SyncLocalReviewsInput struct {
	ClientID      string `json:"clientId"`
	LocationID    string `json:"locationId"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// SyncLocalReviews calls sync_local_reviews: Collect a location's Google
// reviews. Collects the location's latest Google reviews (up to 50) from the
// vendor. Paid: quote it with estimate_cost first. Runs in the background and
// takes a few minutes; read the analysis with get_local_visibility afterwards.
// A second request within the hour lands on the collection already running
// rather than paying twice. With maxCostMicros, it is refused before it starts
// when its estimate is over the limit. Answers 'pull in our latest reviews'.
func (c *Client) SyncLocalReviews(ctx context.Context, in SyncLocalReviewsInput) (json.RawMessage, error) {
	return c.Call(ctx, "sync_local_reviews", in)
}

// TagKeywordsInput is the input to tag_keywords.
type TagKeywordsInput struct {
	ClientID   string   `json:"clientId"`
	TrackerID  string   `json:"trackerId"`
	KeywordIDs []string `json:"keywordIds"`
	Tags       []string `json:"tags"`
	Remove     *bool    `json:"remove,omitempty"`
}

// TagKeywords calls tag_keywords: Tag tracked keywords. Puts tags on a
// tracker's keywords, or takes them off with remove: true. Answers 'label
// these as our brand terms'. Tags are case-insensitive labels such as 'brand'
// or 'money terms' that the tracker view filters by. Keywords the tracker does
// not track are returned in notTracked rather than tagged.
func (c *Client) TagKeywords(ctx context.Context, in TagKeywordsInput) (json.RawMessage, error) {
	return c.Call(ctx, "tag_keywords", in)
}

// UpdateAlertRuleInput is the input to update_alert_rule.
type UpdateAlertRuleInput struct {
	ClientID        string   `json:"clientId"`
	RuleID          string   `json:"ruleId"`
	Threshold       *int64   `json:"threshold,omitempty"`
	CooldownHours   *int64   `json:"cooldownHours,omitempty"`
	WebhookURL      string   `json:"webhookUrl,omitempty"`
	Channel         string   `json:"channel,omitempty"`
	EmailRecipients []string `json:"emailRecipients,omitempty"`
	TelegramChatID  string   `json:"telegramChatId,omitempty"`
	Paused          *bool    `json:"paused,omitempty"`
}

// UpdateAlertRule calls update_alert_rule: Update an alert rule. Changes a
// rule's threshold, cooldown or delivery (channel, webhookUrl — null removes
// it — emailRecipients, telegramChatId), or pauses and resumes it. Answers
// 'this alert is too noisy', 'send this to email instead' or 'mute this for
// now'. Pausing keeps the rule's history.
func (c *Client) UpdateAlertRule(ctx context.Context, in UpdateAlertRuleInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_alert_rule", in)
}

// UpdateClientInput is the input to update_client.
type UpdateClientInput struct {
	ClientID         string   `json:"clientId"`
	Name             string   `json:"name,omitempty"`
	MonthlyBudgetUsd *float64 `json:"monthlyBudgetUsd,omitempty"`
}

// UpdateClient calls update_client: Update a client. Renames a client or
// changes its monthly vendor spend cap (null removes the cap). Answers 'rename
// this client' or 'cap this client's spend'. The cap is enforced on every paid
// lookup from then on.
func (c *Client) UpdateClient(ctx context.Context, in UpdateClientInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_client", in)
}

// UpdateClientBriefInput is the input to update_client_brief.
type UpdateClientBriefInput struct {
	ClientID    string           `json:"clientId"`
	Overview    string           `json:"overview,omitempty"`
	Goal        string           `json:"goal,omitempty"`
	Positioning string           `json:"positioning,omitempty"`
	Audience    string           `json:"audience,omitempty"`
	Voice       string           `json:"voice,omitempty"`
	Add         []map[string]any `json:"add,omitempty"`
	Update      []map[string]any `json:"update,omitempty"`
	Remove      []string         `json:"remove,omitempty"`
}

// UpdateClientBrief calls update_client_brief: Update the client brief.
// Changes the client brief that every teammate and agent reads. Answers
// 'remember this about the client'. Set fields (overview, goal, positioning,
// audience, voice; null or empty clears one, fields left out are untouched);
// add entries (kind competitor, key_page, section or research, each with a
// title and optional body and url); edit entries by id; remove entries by id
// (unknown ids are ignored). Applied in that order, so removing an old
// competitor and adding its replacement in one call works. Record a research
// entry when you finish a piece of work whose conclusions the next one should
// start from; the 50 newest research notes are kept. Returns the brief as it
// now stands. Free.
func (c *Client) UpdateClientBrief(ctx context.Context, in UpdateClientBriefInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_client_brief", in)
}

// UpdateRankTrackerInput is the input to update_rank_tracker.
type UpdateRankTrackerInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId"`
	Name      string `json:"name,omitempty"`
	Device    string `json:"device,omitempty"`
	Frequency string `json:"frequency,omitempty"`
}

// UpdateRankTracker calls update_rank_tracker: Update a rank tracker. Renames
// a tracker or changes its device (desktop, mobile or both) or cadence (daily
// or weekly). Answers 'check these weekly instead'. Weekly costs a seventh as
// much.
func (c *Client) UpdateRankTracker(ctx context.Context, in UpdateRankTrackerInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_rank_tracker", in)
}

// UpdateReportScheduleInput is the input to update_report_schedule.
type UpdateReportScheduleInput struct {
	ClientID   string   `json:"clientId"`
	ScheduleID string   `json:"scheduleId"`
	Cadence    string   `json:"cadence,omitempty"`
	Recipients []string `json:"recipients,omitempty"` // Email addresses the report goes to. At least one.
	Paused     *bool    `json:"paused,omitempty"`
}

// UpdateReportSchedule calls update_report_schedule: Change a report schedule.
// Changes a schedule's cadence or recipients, or pauses or resumes it. Answers
// 'add the new marketing lead to the monthly report' and 'stop the weekly
// email while the client is away'. Recipients, when given, replace the whole
// list. A paused schedule sends nothing until resumed.
func (c *Client) UpdateReportSchedule(ctx context.Context, in UpdateReportScheduleInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_report_schedule", in)
}

// VerifyIndexnowKeyInput is the input to verify_indexnow_key.
type VerifyIndexnowKeyInput struct {
	ClientID string `json:"clientId"`
}

// VerifyIndexnowKey calls verify_indexnow_key: Verify the IndexNow key file.
// Fetches the key file from the client's site and records whether it is there
// and correct. Answers 'is IndexNow set up'. submit_indexnow submits only for
// a verified key. A failure says what was found instead (not served, wrong
// content, unreachable); a site that stops serving the file loses its
// verification.
func (c *Client) VerifyIndexnowKey(ctx context.Context, in VerifyIndexnowKeyInput) (json.RawMessage, error) {
	return c.Call(ctx, "verify_indexnow_key", in)
}

// VerifySiteInput is the input to verify_site.
type VerifySiteInput struct {
	ClientID string `json:"clientId"`
}

// VerifySite calls verify_site: Verify a deployed site and finish onboarding.
// Answers 'I deployed it, is everything connected'. After the onboard_site
// manifest is live, checks and completes everything in one call: the IndexNow
// key file, Google ownership (and then adds and connects the Search Console
// property, with its sixteen months of history), whether analytics has
// received its first event, that robots.txt and the sitemap answer, and starts
// the first crawl. Returns a checklist where every item not done carries the
// exact fix and the next call. Safe to call again. Free.
func (c *Client) VerifySite(ctx context.Context, in VerifySiteInput) (json.RawMessage, error) {
	return c.Call(ctx, "verify_site", in)
}

// VerifySiteOwnershipInput is the input to verify_site_ownership.
type VerifySiteOwnershipInput struct {
	ClientID  string   `json:"clientId"`
	Method    string   `json:"method,omitempty"`
	SiteURL   string   `json:"siteUrl,omitempty"` // A URL-prefix site other than https://<domain>/, such as https://www.example.com/. Not used for dns.
	AddOwners []string `json:"addOwners,omitempty"`
}

// VerifySiteOwnership calls verify_site_ownership: Verify site ownership and
// connect Search Console. Answers 'the verification file is live, connect
// Search Console'. Has Google check the deployed file, tag or record (with no
// method, each is tried), which makes the agency's Google account a verified
// owner; then adds the property to its Search Console and connects it to the
// client, history included. addOwners makes other Google accounts owners too,
// so the site's people keep their own access. verify_site does this for you.
// Free.
func (c *Client) VerifySiteOwnership(ctx context.Context, in VerifySiteOwnershipInput) (json.RawMessage, error) {
	return c.Call(ctx, "verify_site_ownership", in)
}

// WhoamiInput is the input to whoami.
type WhoamiInput struct {
}

// Whoami calls whoami: Who am I connected as. Says who this connection is: the
// agency, whether it is an API key or a person, the role, every permission
// that role holds, and the clients it can reach. Answers 'what can this key
// do'. Call it first when unsure whether a task is allowed. Free.
func (c *Client) Whoami(ctx context.Context, in WhoamiInput) (json.RawMessage, error) {
	return c.Call(ctx, "whoami", in)
}
