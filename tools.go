// Code generated from the Seolyticize tool registry. DO NOT EDIT.
// Regenerate with `pnpm --filter @seolyticize/sdk generate:languages`.

package seolyticize

import (
	"context"
	"encoding/json"
)

// ToolCount is how many tools this SDK was generated with.
const ToolCount = 499

// readOnlyTools are retried on server errors; every other call is not, since
// a write that failed with a 500 may already have happened.
var readOnlyTools = map[string]bool{
	"analyse_server_log":              true,
	"audit_merchant_feed":             true,
	"audit_page":                      true,
	"audit_product_page":              true,
	"batch_domain_metrics":            true,
	"build_disavow_file":              true,
	"check_answer_readiness":          true,
	"check_brand_entity":              true,
	"check_programmatic_pages":        true,
	"check_templated_metadata":        true,
	"check_url_safety":                true,
	"compare_crawls":                  true,
	"describe_datasets":               true,
	"draft_agent_files":               true,
	"estimate_ai_visibility_cost":     true,
	"estimate_cost":                   true,
	"explain_keyword_difficulty":      true,
	"export_ai_visibility_data":       true,
	"export_data":                     true,
	"export_search_console_data":      true,
	"export_sitemap":                  true,
	"find_journalist_requests":        true,
	"find_tools":                      true,
	"forecast_keyword_traffic":        true,
	"forecast_search_clicks":          true,
	"generate_hreflang":               true,
	"generate_llms_txt":               true,
	"generate_schema":                 true,
	"get_agent_readiness":             true,
	"get_ai_agent_visits":             true,
	"get_ai_answer_share":             true,
	"get_ai_citation_history":         true,
	"get_ai_cited_sources":            true,
	"get_ai_crawler_access":           true,
	"get_ai_crawler_activity":         true,
	"get_ai_overview_citations":       true,
	"get_ai_referrals":                true,
	"get_ai_visibility":               true,
	"get_ai_visibility_answer":        true,
	"get_ai_visibility_results":       true,
	"get_ai_visibility_runs":          true,
	"get_ai_visibility_sentiment":     true,
	"get_ai_visibility_setup":         true,
	"get_ai_visibility_trend":         true,
	"get_anchor_overlap":              true,
	"get_audit_performance":           true,
	"get_audit_summary":               true,
	"get_backlink_audit":              true,
	"get_backlink_health":             true,
	"get_billing":                     true,
	"get_bing_crawl_health":           true,
	"get_bing_inbound_links":          true,
	"get_bing_keyword_research":       true,
	"get_bing_performance":            true,
	"get_bing_search_stats":           true,
	"get_brand_split":                 true,
	"get_branding":                    true,
	"get_budget_status":               true,
	"get_cannibalization":             true,
	"get_click_curve":                 true,
	"get_client":                      true,
	"get_client_brief":                true,
	"get_client_costs":                true,
	"get_client_integrations":         true,
	"get_competitor_new_pages":        true,
	"get_competitor_page_changes":     true,
	"get_connection_link":             true,
	"get_content_audit":               true,
	"get_content_decay":               true,
	"get_content_gap":                 true,
	"get_content_groups":              true,
	"get_conversions":                 true,
	"get_costs":                       true,
	"get_crawl_budget":                true,
	"get_crawl_settings":              true,
	"get_crawl_status":                true,
	"get_ctr_outliers":                true,
	"get_deploy_hook":                 true,
	"get_disavow_history":             true,
	"get_domain_facts":                true,
	"get_domain_technologies":         true,
	"get_engine_capabilities":         true,
	"get_entity_pageviews":            true,
	"get_error_catalogue":             true,
	"get_extraction_results":          true,
	"get_freshness_queue":             true,
	"get_funnel_report":               true,
	"get_ga4_audience":                true,
	"get_ga4_ecommerce":               true,
	"get_ga4_key_events":              true,
	"get_ga4_landing_pages":           true,
	"get_ga4_measurement_health":      true,
	"get_ga4_overview":                true,
	"get_ga4_page_performance":        true,
	"get_ga4_site_search":             true,
	"get_ga4_traffic_acquisition":     true,
	"get_geo_readiness":               true,
	"get_goal_report":                 true,
	"get_google_service_account":      true,
	"get_hreflang_parity":             true,
	"get_hreflang_report":             true,
	"get_index_coverage":              true,
	"get_indexnow_key":                true,
	"get_internal_link_graph":         true,
	"get_issue_templates":             true,
	"get_job":                         true,
	"get_keyword_difficulty":          true,
	"get_keyword_gap":                 true,
	"get_keyword_list":                true,
	"get_landing_pages":               true,
	"get_link_checks":                 true,
	"get_link_intersect":              true,
	"get_link_velocity":               true,
	"get_local_visibility":            true,
	"get_log_orphans":                 true,
	"get_long_tail_candidates":        true,
	"get_low_ctr_queries":             true,
	"get_matomo_report":               true,
	"get_media_report":                true,
	"get_naver_keyword_volumes":       true,
	"get_next_steps":                  true,
	"get_opportunities":               true,
	"get_outbound_domains":            true,
	"get_outreach_mailboxes":          true,
	"get_outreach_status":             true,
	"get_page_details":                true,
	"get_page_indexing_reasons":       true,
	"get_page_issues":                 true,
	"get_panel_reach":                 true,
	"get_parameter_report":            true,
	"get_parasite_risk":               true,
	"get_personal_keyword_difficulty": true,
	"get_portfolio_metrics":           true,
	"get_position_distribution":       true,
	"get_preload_readiness":           true,
	"get_price_list":                  true,
	"get_query_intents":               true,
	"get_query_page_swaps":            true,
	"get_questions":                   true,
	"get_rank_history":                true,
	"get_rank_summary":                true,
	"get_recipe":                      true,
	"get_rendering_report":            true,
	"get_report_data":                 true,
	"get_robots_history":              true,
	"get_run_status":                  true,
	"get_search_anomalies":            true,
	"get_search_appearance":           true,
	"get_search_changes":              true,
	"get_search_console_hourly":       true,
	"get_search_console_performance":  true,
	"get_search_demand":               true,
	"get_search_opportunities":        true,
	"get_search_timeline":             true,
	"get_security_report":             true,
	"get_seo_changes":                 true,
	"get_seo_revenue":                 true,
	"get_seo_scorecard":               true,
	"get_serp_features":               true,
	"get_serp_intent_shifts":          true,
	"get_serp_newcomers":              true,
	"get_serp_volatility":             true,
	"get_share_of_search":             true,
	"get_share_of_voice":              true,
	"get_shopping_keyword_gap":        true,
	"get_shopping_serp":               true,
	"get_site_conventions":            true,
	"get_site_status":                 true,
	"get_top_searches":                true,
	"get_topical_authority":           true,
	"get_toxic_link_risk":             true,
	"get_translation_gap":             true,
	"get_trust_signals":               true,
	"get_url_swaps":                   true,
	"get_vendor_reconciliation":       true,
	"get_vitals_history":              true,
	"get_volume_calibration":          true,
	"get_web_analytics_breakdown":     true,
	"get_web_analytics_setup":         true,
	"get_web_analytics_summary":       true,
	"get_written_report":              true,
	"get_yandex_diagnostics":          true,
	"get_yandex_external_links":       true,
	"get_yandex_important_pages":      true,
	"get_yandex_indexing":             true,
	"get_yandex_performance":          true,
	"get_yandex_query_history":        true,
	"get_zero_click_estimates":        true,
	"inspect_bing_url":                true,
	"list_accepted_issues":            true,
	"list_alert_rules":                true,
	"list_annotations":                true,
	"list_client_facts":               true,
	"list_clients":                    true,
	"list_content_plan":               true,
	"list_conversion_goals":           true,
	"list_crawl_issues":               true,
	"list_crawls":                     true,
	"list_data_syncs":                 true,
	"list_digest_schedules":           true,
	"list_funnels":                    true,
	"list_goals":                      true,
	"list_integrations":               true,
	"list_issue_pages":                true,
	"list_keyword_lists":              true,
	"list_keyword_owners":             true,
	"list_link_prospects":             true,
	"list_local_locations":            true,
	"list_log_sources":                true,
	"list_markets":                    true,
	"list_outreach_sequences":         true,
	"list_path_owners":                true,
	"list_portfolios":                 true,
	"list_projects":                   true,
	"list_rank_trackers":              true,
	"list_recent_alerts":              true,
	"list_report_schedules":           true,
	"list_report_templates":           true,
	"list_reports":                    true,
	"list_sitemaps":                   true,
	"list_task_links":                 true,
	"list_tasks":                      true,
	"list_tracked_keywords":           true,
	"list_written_reports":            true,
	"page_vitals":                     true,
	"plan_redirect_map":               true,
	"preview_weekly_digest":           true,
	"propose_keyword_owners":          true,
	"query_data":                      true,
	"query_keyword_list":              true,
	"score_eeat":                      true,
	"search_locations":                true,
	"simulate_serp_snippet":           true,
	"striking_distance":               true,
	"suggest_internal_links":          true,
	"sync_search_console":             true,
	"test_robots":                     true,
	"validate_product_schema":         true,
	"validate_sitemap":                true,
	"validate_structured_data":        true,
	"wait_for_job":                    true,
	"whoami":                          true,
}

// AcceptIssueInput is the input to accept_issue.
type AcceptIssueInput struct {
	ClientID string   `json:"clientId"`
	RuleID   string   `json:"ruleId"`
	Template string   `json:"template,omitempty"`
	URLs     []string `json:"urls,omitempty"`
	Reason   string   `json:"reason"`
	Until    string   `json:"until,omitempty"`
	DryRun   *bool    `json:"dryRun,omitempty"` // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
}

// AcceptIssue calls accept_issue: Accept an audit issue as intentional. Marks
// one audit rule's findings as kept on purpose, on the pages of a URL template
// (template: /login, /legal/*, /docs/*?sort=*; * is one path segment, a final
// ** the rest; a template without a query matches any query) or on listed
// addresses (urls), with the reason and, optionally, the last day it holds
// (until, YYYY-MM-DD). From then on those findings leave get_seo_scorecard,
// the counts of get_audit_summary (shown apart as accepted, each with its
// reason), the weekly digest, create_tasks_from_audit and the alerts;
// list_issue_pages marks them. ruleId comes from get_audit_summary. matchedNow
// is how many of the latest crawl's findings of the rule it covers, out of
// rulePages (both null without a crawl), with a sample of the covered
// addresses; warning is set when it covers nothing, or when a template covers
// more than half the rule's pages (likely too broad). dryRun: true saves
// nothing and answers the same coverage, to check a template first. Examples:
// login pages noindexed by design, old versions of the legal pages, sort
// variants. unaccept_issue undoes it. Answers 'stop flagging our login pages,
// they are noindexed on purpose'. Free.
func (c *Client) AcceptIssue(ctx context.Context, in AcceptIssueInput) (json.RawMessage, error) {
	return c.Call(ctx, "accept_issue", in)
}

// AddAnnotationInput is the input to add_annotation.
type AddAnnotationInput struct {
	ClientID string `json:"clientId"`
	On       string `json:"on"` // The day it happened (UTC).
	Kind     string `json:"kind"`
	Title    string `json:"title"`
	Detail   string `json:"detail,omitempty"`
	URL      string `json:"url,omitempty"`
}

// AddAnnotation calls add_annotation: Note something on the client's timeline.
// Records something that happened on a day — a deploy, a page published or
// rewritten, a campaign, a technical change, a note — so traffic charts and
// get_search_timeline can be read as cause and effect. Answers 'remember that
// we shipped this today'. Deploy hooks note their runs on their own. Free.
func (c *Client) AddAnnotation(ctx context.Context, in AddAnnotationInput) (json.RawMessage, error) {
	return c.Call(ctx, "add_annotation", in)
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

// AddDefaultAlertsInput is the input to add_default_alerts.
type AddDefaultAlertsInput struct {
	ClientID   string `json:"clientId,omitempty"`
	AllClients *bool  `json:"allClients,omitempty"`
}

// AddDefaultAlerts calls add_default_alerts: Add the default alerts. Gives a
// client the four alerts every new client now starts with: the site goes down
// (site-unreachable), non-branded Search Console clicks fall 30% or more week
// on week (traffic-drop 30), a crawl finds critical issues
// (new-critical-issues 1), and a tracked keyword drops 5 or more places
// (rank-drop 5). Each goes to the agency owner's email until changed with
// update_alert_rule. A kind the client already has a rule for is kept as it
// is, so it is safe to repeat. Pass clientId for one client, or allClients:
// true for every live client the key may manage. Answers 'set up the usual
// alerts'. Free.
func (c *Client) AddDefaultAlerts(ctx context.Context, in AddDefaultAlertsInput) (json.RawMessage, error) {
	return c.Call(ctx, "add_default_alerts", in)
}

// AddKeywordsToListInput is the input to add_keywords_to_list.
type AddKeywordsToListInput struct {
	ListID   string           `json:"listId"`
	Keywords []map[string]any `json:"keywords"`
}

// AddKeywordsToList calls add_keywords_to_list: Add keywords to a list. Adds
// keywords to an existing list. Keywords already on it are counted as already
// present and otherwise left alone, except that metrics passed with them
// replace the fields given. Answers 'save these to that list'. Each keyword
// may carry the metrics keyword_explorer or get_keyword_metrics already
// returned (volume, difficulty, CPC in micro-USD, intent, SERP features); they
// are stored with their read time, so refresh_keyword_list_metrics need not
// buy them again. Location code defaults to 2840 (United States), language to
// en, engine to google.
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
// competitors but not the client, with that as the reason. Give targetUrl when
// the prospect is one page — a dead page to reclaim, a thread to answer, an
// article that mentions the brand — so each page is its own row; without it
// the prospect is the whole site. A site or page already on the list is left
// at its current stage, never reset. Answers 'add these sites to our
// outreach'. Free.
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
	PlaceID  string  `json:"placeId,omitempty"`
	Cid      string  `json:"cid,omitempty"`
}

// AddLocalLocation calls add_local_location: Add a business location. Adds a
// business location to a client for local SEO, at a latitude and longitude —
// the centre of every grid scan around it. Give its Google placeId (or cid)
// when known — get_business_profile finds both, with the coordinates — so
// scans match the listing itself, not any business with the same name. Answers
// 'start tracking this branch in the map pack'. Free.
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
// requests that returned no content), aiCrawlerShare, aiByPurpose (AI agents'
// requests split into training crawls — GPTBot, ClaudeBot, CCBot, Meta,
// Bytespider, Amazonbot — search indexing — OAI-SearchBot, Claude-SearchBot,
// PerplexityBot, Applebot, DuckAssistBot — and user fetches — ChatGPT-User,
// Claude-User, Perplexity-User, MistralAI-User, an assistant reading the page
// because someone asked it something), aiUserPaths (the pages assistants
// fetched for people: the nearest a log gets to citations), totalHits and
// totalBytes. Also linesRead, linesUnparsed (lines in a format it did not
// recognise — many means the wrong format, not a quiet site), humanHits and
// truncated. Up to 64 MB per call; split a bigger log and call once per part.
// Nothing is stored. Answers 'what is Googlebot actually crawling on our site'
// and 'which of our pages are AI assistants reading'. Free.
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

// AssessAIVisibilitySentimentInput is the input to assess_ai_visibility_sentiment.
type AssessAIVisibilitySentimentInput struct {
	ClientID      string `json:"clientId"`
	CapturedAt    string `json:"capturedAt,omitempty"`
	MaxAnswers    *int64 `json:"maxAnswers,omitempty"`
	Rejudge       *bool  `json:"rejudge,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// AssessAIVisibilitySentiment calls assess_ai_visibility_sentiment: Judge
// sentiment for answers already captured. Judges how each answer in one check
// (the latest unless capturedAt names one) describes the tracked brands it
// names, for answers not yet judged — those captured before sentiment was on,
// or when the judge was unavailable. One metered Claude call per answer, up to
// maxAnswers (default 50); rejudge: true judges again. Returns judged,
// skipped, left and costMicros. Get the user's go-ahead first: it spends.
// Answers 'fill in sentiment for last week's check'.
func (c *Client) AssessAIVisibilitySentiment(ctx context.Context, in AssessAIVisibilitySentimentInput) (json.RawMessage, error) {
	return c.Call(ctx, "assess_ai_visibility_sentiment", in)
}

// AuditMerchantFeedInput is the input to audit_merchant_feed.
type AuditMerchantFeedInput struct {
	ClientID   string `json:"clientId"`
	MerchantID string `json:"merchantId,omitempty"` // Default: the account the consent opened, or the first of several.
	Limit      *int64 `json:"limit,omitempty"`
}

// AuditMerchantFeed calls audit_merchant_feed: Audit the Merchant Center
// product feed. Reads the client's connected Google Merchant Center account
// (Merchant API) and audits its products: Google's own item-level issues
// grouped by code, worst first (disapproved, then demoted), with how many
// products each hits, examples and Google's help link; approval counts per
// destination; and what Google does not flag but Shopping placement rewards —
// titles too short or without the brand, thin descriptions, missing GTINs,
// single images, out of stock. Answers 'why are our products disapproved' and
// 'how good is our feed'. Needs Merchant Center connected (get_connection_link
// names the link; /api/connect/merchant/start). Free.
func (c *Client) AuditMerchantFeed(ctx context.Context, in AuditMerchantFeedInput) (json.RawMessage, error) {
	return c.Call(ctx, "audit_merchant_feed", in)
}

// AuditPageInput is the input to audit_page.
type AuditPageInput struct {
	ClientID string `json:"clientId"`
	URL      string `json:"url"`
	Render   string `json:"render,omitempty"`
	PageType string `json:"pageType,omitempty"`
}

// AuditPage calls audit_page: Audit one page. Fetches one page of the client's
// site now and runs the crawler's full rule engine over it, plus page-level
// thresholds: title 30–60 characters, meta description 120–160, exactly one H1
// and no skipped heading levels, og:title, og:description, og:image and
// og:url, twitter:card, twitter:title and twitter:description, alt text 10–125
// characters, images over 200KB (warning) or 500KB (critical), and a
// word-count floor for the page type (home 500, service 800, primary location
// 600, secondary location 500, blog 1,500, product, category and about 400,
// landing 600, FAQ 800; pageType overrides the guess from the URL and markup),
// Flesch reading ease (60–70 suits a general audience; only under 30 is
// flagged, since it is not a ranking factor), query parameters in the URL
// (tracking parameters aside), links from the copy itself to the site's own
// pages against the page type's target (blog 5–10, service 3–5, product 2–4;
// navigation, header and footer links not counted), and templated metadata (a
// description that restates the title, closes on a stock call to action or
// repeats the title's brand suffix). Returns an overall score and five
// sub-scores (on-page, content, technical, schema, images; null where not
// measured), every issue by severity with its fix, and generate_schema's
// suggestions for the page. render 'auto' renders the page in a browser when
// its HTML looks like an unrendered shell. A heuristic. Answers 'what is wrong
// with this page' and 'is this page ready'. Free.
func (c *Client) AuditPage(ctx context.Context, in AuditPageInput) (json.RawMessage, error) {
	return c.Call(ctx, "audit_page", in)
}

// AuditProductPageInput is the input to audit_product_page.
type AuditProductPageInput struct {
	ClientID         string `json:"clientId"`
	URL              string `json:"url"`
	ManufacturerText string `json:"manufacturerText,omitempty"` // The manufacturer's own description, to check the page is not a copy.
}

// AuditProductPage calls audit_product_page: Audit one product page. Reads one
// product page now and scores it 0–100 (a heuristic, said so in the answer)
// from six weighted parts, each check pass, warn or fail with the reason, or
// null when it cannot be judged (then left out and the weights rescaled):
// Product schema (25, the completeness validate_product_schema gives), title
// and meta (15: a title under 60 characters, a meta description under 155 with
// a price and a call to action), images (20: at least 3, each at least 50,000
// pixels where the page declares sizes, alt text), content (20: 200 words or
// more, not the manufacturer's copy — pass manufacturerText to compare — a
// specifications table, reviews), linking (10: breadcrumbs, a link to its
// category, related products) and technical (10: status 200, indexable, a self
// canonical, a viewport, HTTPS). Answers 'what is wrong with this product
// page' and 'why does this product not rank'. Free.
func (c *Client) AuditProductPage(ctx context.Context, in AuditProductPageInput) (json.RawMessage, error) {
	return c.Call(ctx, "audit_product_page", in)
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

// BuildComparisonPageInput is the input to build_comparison_page.
type BuildComparisonPageInput struct {
	ClientID      string           `json:"clientId"`
	PageType      string           `json:"pageType"`
	Competitors   []map[string]any `json:"competitors"`
	Features      []string         `json:"features,omitempty"`
	SourceURLs    []string         `json:"sourceUrls,omitempty"`
	Category      string           `json:"category,omitempty"` // For a roundup: what the tools are, e.g. 'CRM'.
	Demand        *bool            `json:"demand,omitempty"`
	Instructions  string           `json:"instructions,omitempty"`
	MaxCostMicros *int64           `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool            `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// BuildComparisonPage calls build_comparison_page: Build a comparison page.
// Plans an honest comparison page — vs (the client against one rival),
// alternatives (to the first competitor), roundup (best tools in a category)
// or table — from what each product publishes: its pricing and feature pages
// (and any sourceUrls) are read and a feature matrix filled, each cell ✅, ⚠️
// partial, ❌ or 'Not publicly available' with its source URL, and every price
// 'as of' the read. Gives the title from the pageType's template, an H1 under
// 70 characters, a sourced outline of at least 1,500 words with calls to
// action above the fold, after the table and at the end (never inside a
// competitor's section), each competitor's strengths acknowledged, the
// client's product named as the client's, JSON-LD (SoftwareApplication with
// offers; ItemList in ranked order for roundups; no AggregateRating without
// real review data) checked against Google's requirements; each product's
// hidden costs from its pricing page (setup fees, per-user pricing, overages,
// transaction fees, annual-only billing, minimum seats), a 'Last updated'
// stamp, related comparisons to link, and checks that the outline carries a
// methodology section, the hidden costs and related comparisons; and the
// comparison searches by pattern ('[a] vs [b]', '[a] vs [b] pricing', 'is [a]
// better than [b]', '[a] alternatives', '[a] alternatives [year]', '[a] review
// [year]', 'best [category] tools', '[category] comparison'), with demand:
// true their volumes (one paid keyword-metrics read). Answers 'write a vs page
// against our rival' and 'what should our alternatives page say'. The outline
// is one metered Claude call; quote with estimate_cost.
func (c *Client) BuildComparisonPage(ctx context.Context, in BuildComparisonPageInput) (json.RawMessage, error) {
	return c.Call(ctx, "build_comparison_page", in)
}

// BuildContentBriefInput is the input to build_content_brief.
type BuildContentBriefInput struct {
	ClientID        string   `json:"clientId"`
	Keyword         string   `json:"keyword"`
	RelatedKeywords []string `json:"relatedKeywords,omitempty"`
	TargetURL       string   `json:"targetUrl,omitempty"` // The page this brief is for, if it exists, so it is not suggested as its own internal link.
	PageType        string   `json:"pageType,omitempty"`  // The kind of page, for its section template: service, blog, case_study, category, landing, faq, location, about or homepage. Default: guessed from the search int
	Depth           *int64   `json:"depth,omitempty"`
	Save            *bool    `json:"save,omitempty"`
	LocationCode    *int64   `json:"locationCode,omitempty"`  // Google location code (2840 United States, 2826 United Kingdom, 2724 Spain). Default: the client's project.
	LanguageCode    string   `json:"languageCode,omitempty"`  // Language code (en, es). Default: the client's project.
	MaxCostMicros   *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun          *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// BuildContentBrief calls build_content_brief: Build a content brief. A
// writer's brief for a keyword (or a cluster: the keyword plus
// relatedKeywords), built from the pages that rank: search intent and the
// format that wins (list, how-to, comparison, review, landing page, article),
// target length, the outline to cover, must-cover and nice-to-have terms,
// questions to answer (People Also Ask and the ranking pages'), structure and
// schema to include, the client's own pages to link to and from, and each
// competing page with what it misses — and ownership, a warning for each
// keyword the keyword registry already gives another page ('this already
// belongs to /x'). Also: each competitor scored 1–10 on depth, format, SEO and
// UX (out of 40), with claude-seo's full excluded list left out as
// non-competitors (encyclopedias, social and content platforms, search
// engines, marketplaces, forums, news and media, data platforms, directories,
// comparison aggregators, job boards, SEO tools, AI platforms, government and
// academic domains, and tag, author, login, legal and feed paths on any
// domain), and thinCompetition set when fewer than 3 real competitors remain;
// a section template for the page type (pageType: service, blog, case_study,
// category, landing, faq, location, about or homepage; default from the
// intent) with its goal, sections, schema (never FAQPage) and keyword
// placement; gaps typed topic, depth or quality with priority = impact ×
// advantage ÷ effort; a required information-gain statement; the E-E-A-T the
// page needs; words per section; the featured-snippet section and format;
// title 50–60 and description 130–150 character targets; and with targetUrl,
// improve mode marking each section keep, strengthen or add against the page
// as it stands. Answers 'brief me on this keyword' and 'what should this post
// cover'. save: true keeps it as a research note on the client brief. Same
// cost as get_serp_content_profile.
func (c *Client) BuildContentBrief(ctx context.Context, in BuildContentBriefInput) (json.RawMessage, error) {
	return c.Call(ctx, "build_content_brief", in)
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

// CapturePageScreenshotsInput is the input to capture_page_screenshots.
type CapturePageScreenshotsInput struct {
	ClientID  string   `json:"clientId"`
	URL       string   `json:"url"`
	Viewports []string `json:"viewports,omitempty"`
	Parity    *bool    `json:"parity,omitempty"`
}

// CapturePageScreenshots calls capture_page_screenshots: Screenshots and
// layout checks at four screen sizes. Renders one of the client's pages in a
// real browser at 1920×1080, 1366×768, 768×1024 and 375×812 (viewports narrows
// it), stores a PNG of each first screen (links last a day) and checks the
// layout there: whether the H1, the primary call to action and the hero image
// are above the fold, horizontal scrolling, tap targets under 24×24 px (fail)
// and under 48×48 (warn) — inline text links excepted — the base font size
// under 16px, overlapping controls and clipped text, and a fixed overlay
// covering over half the screen on load (cookie and consent banners excepted).
// parity (default on) also fetches the page as Googlebot smartphone and as
// desktop Chrome and compares words, links, headings and structured data
// types, since Google indexes the mobile version. Answers 'does the page work
// on a phone', 'is the CTA above the fold' and 'does Google's mobile crawler
// see the same page'. No vendor charge; it renders pages on the client's
// behalf, so it needs project:manage.
func (c *Client) CapturePageScreenshots(ctx context.Context, in CapturePageScreenshotsInput) (json.RawMessage, error) {
	return c.Call(ctx, "capture_page_screenshots", in)
}

// CheckAIBrandAccuracyInput is the input to check_ai_brand_accuracy.
type CheckAIBrandAccuracyInput struct {
	ClientID      string   `json:"clientId"`
	Facts         []string `json:"facts,omitempty"`
	Question      string   `json:"question,omitempty"`
	Platforms     []string `json:"platforms,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// CheckAIBrandAccuracy calls check_ai_brand_accuracy: Do assistants describe
// the brand correctly. Asks assistants about the client's brand now (the
// ChatGPT app and Perplexity by default, or chat_gpt, claude, gemini) and
// checks every claim they make against the facts given — or, by default, the
// client's registered facts (set_client_facts), and failing those the client
// brief's overview and positioning — each claim marked supported, contradicted
// or unverified (the facts do not say), with the fact that decides it. Answers
// 'what does ChatGPT get wrong about us' and 'is the AI description of our
// brand accurate'. Paid: the live answers on DataForSEO plus a metered Claude
// call; quote it with estimate_cost.
func (c *Client) CheckAIBrandAccuracy(ctx context.Context, in CheckAIBrandAccuracyInput) (json.RawMessage, error) {
	return c.Call(ctx, "check_ai_brand_accuracy", in)
}

// CheckAnswerReadinessInput is the input to check_answer_readiness.
type CheckAnswerReadinessInput struct {
	ClientID string   `json:"clientId"`
	URLs     []string `json:"urls,omitempty"`
	TopPages *int64   `json:"topPages,omitempty"`
}

// CheckAnswerReadiness calls check_answer_readiness: Is the page written to be
// quoted by assistants. Checks pages of the client's site for what makes an
// assistant quote them: a direct answer in the first 100 words (not a
// preamble), a plain definition, sourced claims (links to sources outside the
// site, against the figures stated), a named author, a published or updated
// date, and structured data — a 0–100 score and a fix list per page, and the
// site's fixes in order of how many pages need them. Give urls, or check the
// top pages by Search Console clicks (topPages). Answers 'why is this page not
// cited' and 'how do we make our pages answer-ready'. Free.
func (c *Client) CheckAnswerReadiness(ctx context.Context, in CheckAnswerReadinessInput) (json.RawMessage, error) {
	return c.Call(ctx, "check_answer_readiness", in)
}

// CheckBrandEntityInput is the input to check_brand_entity.
type CheckBrandEntityInput struct {
	ClientID  string   `json:"clientId"`
	Brand     string   `json:"brand,omitempty"`
	Languages []string `json:"languages,omitempty"`
}

// CheckBrandEntity calls check_brand_entity: Does Google know the brand as an
// entity. Looks the client's brand (default its name) up in Google's Knowledge
// Graph Search API: whether there is an entity for it, its types, description,
// Wikipedia article and official URL, whether that URL is the client's own
// site, and how far it scores ahead of the next candidate — matched, ambiguous
// or absent — with what to do about each (Organization schema with sameAs, a
// Wikidata item, one consistent name). Answers 'does Google have a knowledge
// panel entity for us' and 'is our brand an entity'. Free.
func (c *Client) CheckBrandEntity(ctx context.Context, in CheckBrandEntityInput) (json.RawMessage, error) {
	return c.Call(ctx, "check_brand_entity", in)
}

// CheckDeployInput is the input to check_deploy.
type CheckDeployInput struct {
	ClientID        string   `json:"clientId"`
	PreviewURL      string   `json:"previewUrl,omitempty"`
	URLs            []string `json:"urls,omitempty"`
	MaxURLs         *int64   `json:"maxUrls,omitempty"`
	FailOn          string   `json:"failOn,omitempty"`
	BaselineCrawlID string   `json:"baselineCrawlId,omitempty"` // A complete crawl of this client to compare with, from list_crawls; default the latest complete one.
}

// CheckDeploy calls check_deploy: Check a deploy before it ships. Fetches
// pages now and compares what they tell search engines (status, robots,
// canonical, hreflang, title, H1, structured data types; and, against crawls
// that recorded them, meta description, Open Graph tags, structured data
// values, H2 outline and copy) with the latest complete crawl's snapshot of
// the same paths, or the complete crawl baselineCrawlId names (from
// list_crawls), to compare against a crawl from before a release; baseline
// says which crawl was used and when it ran, and notes warns when the latest
// crawl started before the last recorded deploy (deploy hook or deploy
// annotation), so the baseline may already include that release. Give
// previewUrl (a deploy preview or staging origin: the crawl's pages, one or
// more per template, are fetched on that host, up to maxUrls) or urls
// (specific addresses, on production or a preview; matched by path).
// Canonicals and hreflang on the preview host are read as the production host,
// so only real changes show. passed is false when anything at or above failOn
// (default critical: noindex added, canonical changed, title dropped, a page
// now 4xx/5xx, hreflang lost, a page unreachable) changed; markdown is a ready
// pull-request comment. The client's set_crawl_auth credentials are sent to
// the preview host too, so a password-protected preview can be checked. Used
// by the Seolyticize GitHub Action. Answers 'will this deploy hurt SEO' and
// 'did the release noindex anything'. No vendor charge; it fetches pages on
// the client's behalf, so it needs project:manage.
func (c *Client) CheckDeploy(ctx context.Context, in CheckDeployInput) (json.RawMessage, error) {
	return c.Call(ctx, "check_deploy", in)
}

// CheckNapConsistencyInput is the input to check_nap_consistency.
type CheckNapConsistencyInput struct {
	ClientID            string   `json:"clientId"`
	LocationID          string   `json:"locationId"`
	UseListing          *bool    `json:"useListing,omitempty"`
	Directories         *bool    `json:"directories,omitempty"`
	IndustryDirectories []string `json:"industryDirectories,omitempty"` // Extra directory domains for the trade, e.g. avvo.com, healthgrades.com. Null searches the trade's own directories, detected from the listing's categories and th
	MaxCostMicros       *int64   `json:"maxCostMicros,omitempty"`       // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun              *bool    `json:"dryRun,omitempty"`              // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// CheckNapConsistency calls check_nap_consistency: Do the site and the Google
// listing agree. Compares the business's name, address and phone on its Google
// listing with what its site says — the homepage and its contact, about and
// location pages, text and structured data — and says which agree, which
// differ and how. Inconsistent details weaken local rankings and confuse
// assistants. useListing reads the live Google listing (paid, by place id)
// instead of the location as saved. directories also audits citations: Google
// is searched for the business's name and phone on Yelp, Foursquare, Facebook,
// Yellow Pages, Superpages, BBB, Tripadvisor, MapQuest, Apple Maps, Nextdoor
// and more, plus the trade's own directories (industryDirectories, detected
// when not given), and each directory's name, address and phone as shown is
// compared with the listing (null where the result did not show it), with a
// severity for the worst disagreement — a different name critical, address
// warning, phone notice; directories with no listing found are named. Bing
// Places and Apple Business Connect have no self-serve API, so neither is read
// directly. Answers 'is our NAP consistent' and 'which directories list us
// wrongly'. Free with the saved location; quote useListing and directories
// (two results pages) with estimate_cost.
func (c *Client) CheckNapConsistency(ctx context.Context, in CheckNapConsistencyInput) (json.RawMessage, error) {
	return c.Call(ctx, "check_nap_consistency", in)
}

// CheckProgrammaticPagesInput is the input to check_programmatic_pages.
type CheckProgrammaticPagesInput struct {
	ClientID    string `json:"clientId"`
	CrawlID     string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	Template    string `json:"template,omitempty"`
	MinVariants *int64 `json:"minVariants,omitempty"`
	Limit       *int64 `json:"limit,omitempty"`
}

// CheckProgrammaticPages calls check_programmatic_pages: Do the template pages
// earn their place. Finds the site's template pages (the same path with one
// variable segment, minVariants or more of them; or one template given as a
// glob such as /sms/*) in a crawl and, for every variant: how much of its text
// is its own (the share no more than a third of its siblings repeat) and how
// close its nearest sibling is (simhash bits), whether it is thin against the
// template's median, and its demand (Search Console clicks and impressions
// over 90 days, stored search volume for its H1) — with a verdict: keep, merge
// (all but a duplicate of a sibling with more demand: fold it in) or noindex
// (thin or boilerplate, and no demand). Each template also gets checks with
// six 0–100 sub-scores (data, template uniqueness, URL, linking, thin risk,
// index management; null where unmeasured): a gate (100+ unreviewed pages is a
// warning, 500+ says stop), thin content (under 40% own text is thin, under
// 30% fails, under 300 words flagged), descriptions that open with the title,
// end on a stock call to action or share one shape with the name swapped in,
// URL rules (lowercase, hyphens, under 100 characters, no query string, one
// trailing-slash style, no duplicate slugs), links to 3–5 siblings,
// self-referencing canonicals, and Google's indexing (URL Inspection) against
// the pages meant to be indexed. checks.extras adds near-duplicate records
// (title, H1 and description sharing more than 80% of their words),
// boilerplate (pages whose copy is over 60% words most of the template's pages
// share), a review sample of 5% of the template, thinnest first (up to 10% is
// better), noindexed variants the sitemaps still list, and a crawl-budget
// notice past 10,000 pages. Answers 'which of our location pages should exist'
// and 'are these programmatic pages thin'. For the demand of combinations not
// yet built, use get_demand_matrix. Free.
func (c *Client) CheckProgrammaticPages(ctx context.Context, in CheckProgrammaticPagesInput) (json.RawMessage, error) {
	return c.Call(ctx, "check_programmatic_pages", in)
}

// CheckShoppingSerpInput is the input to check_shopping_serp.
type CheckShoppingSerpInput struct {
	ClientID      string   `json:"clientId"`
	Keywords      []string `json:"keywords"`
	LocationCode  *int64   `json:"locationCode,omitempty"`
	LanguageCode  string   `json:"languageCode,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// CheckShoppingSerp calls check_shopping_serp: Queue Google Shopping results
// for keywords. Queues a Google Shopping results page for each keyword (up to
// 20) at DataForSEO, which answers in a few minutes; collect them with
// get_shopping_serp and the task ids returned. Answers 'track our Google
// Shopping positions'. One purchase per keyword. A paid lookup on the client's
// account (spend says how much); quote it with estimate_cost.
func (c *Client) CheckShoppingSerp(ctx context.Context, in CheckShoppingSerpInput) (json.RawMessage, error) {
	return c.Call(ctx, "check_shopping_serp", in)
}

// CheckTemplatedMetadataInput is the input to check_templated_metadata.
type CheckTemplatedMetadataInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"`
}

// CheckTemplatedMetadata calls check_templated_metadata: Descriptions a
// template wrote. Reads every crawled page's title and meta description
// (latest crawl, or crawlId) for the shapes bulk generation leaves: a
// description that opens by restating its own title and closes on a stock call
// to action ('Try it free now.', 'Get started today!'; templated, high), one
// that only echoes the title (medium), the title's brand suffix pasted into
// the description (low), and a description that is the title again (medium).
// Site risk is high when 3 or more pages and at least 30% of those checked are
// templated, medium from 10%; a stock closer shared by 3 or more pages is
// flagged too. Returns the counts, the shared closers, and the flagged pages
// worst first with a fix each. Deterministic string comparison: a heuristic,
// which claims no Google update targeted the pattern. Answers 'are our meta
// descriptions templated' and 'which descriptions should we rewrite'. Free.
func (c *Client) CheckTemplatedMetadata(ctx context.Context, in CheckTemplatedMetadataInput) (json.RawMessage, error) {
	return c.Call(ctx, "check_templated_metadata", in)
}

// CheckURLSafetyInput is the input to check_url_safety.
type CheckURLSafetyInput struct {
	ClientID string   `json:"clientId"`
	URLs     []string `json:"urls,omitempty"`
	Outbound *bool    `json:"outbound,omitempty"` // Also check the external sites the homepage links to (up to 40).
}

// CheckURLSafety calls check_url_safety: Is the site, or a site it links to,
// flagged as unsafe. Looks up the client's homepage, the addresses given, and
// (outbound, on by default) the external sites its homepage links to on
// Google's Web Risk lists — malware, social engineering (phishing) and
// unwanted software — and says which are flagged, with the threat types and
// until when the match holds. A flagged own page is what turns into a browser
// warning and lost traffic; a flagged outbound link is a hacked partner or a
// planted spam link. Answers 'has our site been flagged by Google as unsafe'
// and 'do we link to any malware sites'. Free.
func (c *Client) CheckURLSafety(ctx context.Context, in CheckURLSafetyInput) (json.RawMessage, error) {
	return c.Call(ctx, "check_url_safety", in)
}

// ClusterKeywordsInput is the input to cluster_keywords.
type ClusterKeywordsInput struct {
	ClientID      string   `json:"clientId"`
	Keywords      []string `json:"keywords,omitempty"`
	ListID        string   `json:"listId,omitempty"`
	MinShared     *int64   `json:"minShared,omitempty"`
	Depth         *int64   `json:"depth,omitempty"`
	Mode          string   `json:"mode,omitempty"`          // head: every keyword overlaps the topic's top keyword (plan pages with this). connected: chains of overlap join topics.
	LocationCode  *int64   `json:"locationCode,omitempty"`  // Default: the client's project.
	LanguageCode  string   `json:"languageCode,omitempty"`  // Default: the client's project.
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// ClusterKeywords calls cluster_keywords: Group keywords into topics. Groups
// keywords into topics by how much their Google results overlap: keywords
// whose top results share at least minShared pages are one topic, which one
// page should target. Each topic has its parent topic (the highest-volume
// keyword), total volume, intent, the pages that rank across it, and where the
// client already ranks. overlap lists keyword pairs by shared top-ten results:
// 7–10 merge (one page), 4–6 same cluster, 2–3 interlink, 0–1 separate, after
// leaving out up to the 5 commonest domains that rank on over half the results
// pages (Wikipedia, Reddit and the like). architecture leaves navigational
// keywords out and lays the topics out as one pillar (2,500–4,000 words) with
// 2–5 clusters of 2–4 spokes (1,200–1,800 words), one of eight templates
// picked by intent and wording (ultimate guide, how-to, listicle, explainer,
// comparison, review, best-of, landing page), a link matrix {from, to, type,
// anchor} whose anchors vary so no one carries over 40% of a page's inbound
// links: spoke ↔ pillar mandatory, 2–3 spoke-to-spoke, at most one
// cross-cluster, at least 3 inbound per post, every post within 2 clicks of
// the pillar, 5–21 posts in all, and JSON-LD for the pillar (Article,
// BreadcrumbList, ItemList of spokes) and each spoke (Article isPartOf the
// pillar). scorecard checks the client's live pages against it from the latest
// crawl: no orphans, pillar links 100% both ways, 80% of the planned links
// present. Answers 'which of these keywords belong on one page' and 'how many
// pages do we need'. Give keywords (up to 200) or a keyword listId. One
// results page per keyword, each a paid lookup on the client's account unless
// someone bought it in the last day — quote with estimate_cost first.
func (c *Client) ClusterKeywords(ctx context.Context, in ClusterKeywordsInput) (json.RawMessage, error) {
	return c.Call(ctx, "cluster_keywords", in)
}

// CompareCrawlsInput is the input to compare_crawls.
type CompareCrawlsInput struct {
	ClientID       string `json:"clientId"`
	CrawlID        string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	AgainstCrawlID string `json:"againstCrawlId,omitempty"`
	Sample         *int64 `json:"sample,omitempty"`
}

// CompareCrawls calls compare_crawls: What changed since the last crawl.
// Compares a crawl's issues with the crawl before it (or againstCrawlId), page
// by page: issues that are new, issues that were fixed, and checks whose count
// went up (regressed) or down (improved) — a check that flagged a few new
// pages while clearing more counts as improved — or stayed level with
// different pages (shifted), each with the pages that came and went — plus
// pages added, removed and changed. Answers 'did the fixes work', 'what broke
// this week' and 'what should the weekly audit report'. Free.
func (c *Client) CompareCrawls(ctx context.Context, in CompareCrawlsInput) (json.RawMessage, error) {
	return c.Call(ctx, "compare_crawls", in)
}

// CompareSitesInput is the input to compare_sites.
type CompareSitesInput struct {
	ClientID        string   `json:"clientId"`
	A               string   `json:"a"` // The site under test, e.g. staging.
	B               string   `json:"b"` // The reference, e.g. production.
	Paths           []string `json:"paths,omitempty"`
	MaxPages        *int64   `json:"maxPages,omitempty"`
	SendCrawlAuthTo string   `json:"sendCrawlAuthTo,omitempty"`
	FollowRedirect  *bool    `json:"followRedirect,omitempty"` // When one side's home redirects to the other's host, follow it and check each path lands on the same path there, instead of stopping.
}

// CompareSites calls compare_sites: Compare staging with production. Fetches
// the same paths on two origins — staging (a) and production (b), say — and
// lists every difference in what they tell search engines: status, robots,
// canonical (read on each side's own host), hreflang, title, H1 and structured
// data, with critical ones first and a Markdown summary. Paths default to the
// client's latest crawl (its shallowest pages, spread across templates) or,
// without a crawl, the links on b's home page; paths narrows it.
// sendCrawlAuthTo sends the client's set_crawl_auth credentials (basic auth,
// headers, cookies) to that side only, for a password-protected staging site.
// When one side's home redirects to the other's host (www to the apex),
// redirect names the status and target and nothing is diffed, since that side
// has no pages of its own; single paths that redirect across are listed in
// redirected. followRedirect: true follows that redirect instead of stopping:
// each path on the redirecting side is fetched to where it lands and compared
// with the same path on the other side (the redirect's destination),
// redirectMap lists where each landed, and a path the redirect sends elsewhere
// (the home page, a 404) or a temporary redirect shows as a change; redirect
// stays reported and followedRedirect says it was followed. Answers 'is
// staging ready to go live', 'what will the migration change' and 'does www
// still redirect every page to the same path'. No vendor charge; it fetches
// pages on the client's behalf, so it needs project:manage.
func (c *Client) CompareSites(ctx context.Context, in CompareSitesInput) (json.RawMessage, error) {
	return c.Call(ctx, "compare_sites", in)
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

// ConnectIntegrationInput is the input to connect_integration.
type ConnectIntegrationInput struct {
	ClientID   string `json:"clientId"`
	Provider   string `json:"provider"`
	Token      string `json:"token,omitempty"`
	WebhookURL string `json:"webhookUrl,omitempty"`
	Channel    string `json:"channel,omitempty"`
	ChatID     string `json:"chatId,omitempty"`
	Site       string `json:"site,omitempty"`
	Email      string `json:"email,omitempty"`
	ProjectKey string `json:"projectKey,omitempty"`
	IssueType  string `json:"issueType,omitempty"`
	TeamID     string `json:"teamId,omitempty"`
	Owner      string `json:"owner,omitempty"`
	Repo       string `json:"repo,omitempty"`
	IDSite     string `json:"idSite,omitempty"` // For Matomo: the site's id (idSite).
}

// ConnectIntegration calls connect_integration: Connect Slack, Discord,
// Telegram, Airtable, a tracker or Matomo. Connects one of the client's own
// tools by token. Slack: webhookUrl (an incoming webhook), or token (a bot
// token, xoxb-…) with channel. Discord: webhookUrl. Telegram: token (from
// BotFather) with chatId. Airtable: token (a personal access token with
// data.records:read and write on the base). Jira: site
// (https://….atlassian.net), email, token (an API token), projectKey, optional
// issueType. Linear: token (a personal API key) and teamId. GitHub: token
// (fine-grained, Issues read and write), owner and repo. Matomo (analytics
// instead of GA4): site (the Matomo address, https://…), idSite (the site's id
// in Matomo) and token (a token_auth with view access, from Administration →
// Personal → Security; a POST-only token works); read with get_matomo_report.
// Checked with one call, then sealed. For Linear and GitHub it returns a
// webhook URL and secret to add in the tracker for instant status sync
// (polling hourly works without). Answers 'post the digest to our Slack',
// 'mirror tasks into Jira'. Free.
func (c *Client) ConnectIntegration(ctx context.Context, in ConnectIntegrationInput) (json.RawMessage, error) {
	return c.Call(ctx, "connect_integration", in)
}

// ConnectLogSourceInput is the input to connect_log_source.
type ConnectLogSourceInput struct {
	ClientID     string         `json:"clientId"`
	Kind         string         `json:"kind"`
	Label        string         `json:"label,omitempty"`
	S3           map[string]any `json:"s3,omitempty"`
	VercelVerify string         `json:"vercelVerify,omitempty"` // The verification value Vercel shows when adding the drain, echoed back as x-vercel-verify.
}

// ConnectLogSource calls connect_log_source: Connect a server log feed.
// Connects a feed of the client's access logs so crawler reports stay current
// without uploads. kind: cloudflare-logpush (an HTTP destination for a Logpush
// http_requests job), cloudflare-worker (a Worker snippet for zones without
// Logpush), vercel (a log drain), fastly (an HTTPS logging endpoint, with the
// JSON format string to use), s3 (a bucket and prefix the worker pulls new
// objects from every hour — combined log format or Logpush NDJSON, gzipped or
// not; give s3 with read-only keys, stored encrypted) or upload (a URL to POST
// log files to). Returns the ingest URL (shown once: it carries the secret),
// the setup steps and any snippet. Only crawler requests are kept — visitors'
// requests are counted and dropped — and each crawler's claim is checked
// against reverse DNS or its published IP ranges. Answers 'stream our
// Cloudflare logs in' and 'set up a Vercel log drain'. Free.
func (c *Client) ConnectLogSource(ctx context.Context, in ConnectLogSourceInput) (json.RawMessage, error) {
	return c.Call(ctx, "connect_log_source", in)
}

// ConnectRevenueSourceInput is the input to connect_revenue_source.
type ConnectRevenueSourceInput struct {
	ClientID    string `json:"clientId"`
	Provider    string `json:"provider"`
	APIKey      string `json:"apiKey,omitempty"`
	Shop        string `json:"shop,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
	Sandbox     *bool  `json:"sandbox,omitempty"`
}

// ConnectRevenueSource calls connect_revenue_source: Connect Stripe, Shopify
// or Paddle. Connects the client's own store so its orders can be counted
// against SEO. Stripe: a restricted key (rk_…) with read access to Charges and
// PaymentIntents — or use get_integration_connect_link for Stripe Connect.
// Shopify: the store's myshopify.com domain and a custom app's Admin API
// access token with read_orders — or the connect link for OAuth. Paddle: an
// API key with transaction.read (sandbox: true for a sandbox account). The key
// is checked, sealed, and the last 90 days imported straight away. Answers
// 'connect the client's Shopify'. Free.
func (c *Client) ConnectRevenueSource(ctx context.Context, in ConnectRevenueSourceInput) (json.RawMessage, error) {
	return c.Call(ctx, "connect_revenue_source", in)
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
	PaidSource      *bool    `json:"paidSource,omitempty"`
}

// CreateAlertRule calls create_alert_rule: Create an alert rule. Raises an
// alert when a client's rankings, traffic, issues or backlinks cross a
// threshold, or the site goes down, and delivers it by channel. Threshold by
// kind: rank-drop positions fallen; traffic-drop percent lost;
// new-critical-issues and backlinks-lost a count; serp-features features lost
// (or won); competitor-new-pages new pages one tracked competitor's sitemap
// gained; page-changes own pages whose title, meta description, H1 or
// canonical changed, went noindex or 404 between the last two audits;
// ai-mention-drop percentage points the client's AI mention or citation rate
// fell from the previous check; competitor-new-keywords tracked keywords one
// competitor started ranking for; web-mentions new pages our crawl found
// naming the brand; link-velocity referring domains gained or lost in one day,
// far above that domain's own baseline, for the client or a tracked
// competitor. The watch kinds look back to their last firing (a week at most)
// and stay silent until there is something to compare. Channels: slack
// (webhookUrl, Slack's format; the default), webhook (webhookUrl, plain JSON),
// email (emailRecipients, up to ten) or telegram (telegramChatId, with the
// agency's bot from Settings). Webhooks must be https and public. With no
// channel and no destination, the alert goes to the agency owner's email and
// deliveryNote says so; a channel named without its field is refused as
// no_delivery_target. paidSource (competitor-new-keywords and web-mentions
// only; off by default; needs spend permission) also buys the signal from
// DataForSEO on every sweep: for competitor-new-keywords, Labs ranked keywords
// each tracker competitor (up to 5) newly ranks for at the tracker's market,
// not only tracked keywords; for web-mentions, pages DataForSEO's content
// index first saw naming the brand, off the client's own domains. The sweep is
// hourly but a purchase is reused from the archive while fresh, so it buys at
// most once a week per competitor (once a day for mentions); it is capped by
// the client's monthly budget, and a failed or over-budget read falls back to
// the free signal. Quote it first with estimate_cost. Answers 'tell me when
// rankings drop'.
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
	Mode              string   `json:"mode,omitempty"`              // test mints a sandbox key: it acts in the agency's sandbox, with sample clients and DataForSEO's documented answers at $0, no model calls, and no reach into live
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

// CreateBusinessPostInput is the input to create_business_post.
type CreateBusinessPostInput struct {
	ClientID     string         `json:"clientId"`
	LocationID   string         `json:"locationId,omitempty"` // The client's location, from list_local_locations. Default: its one connected location.
	Summary      string         `json:"summary"`
	CallToAction map[string]any `json:"callToAction,omitempty"`
	DryRun       *bool          `json:"dryRun,omitempty"` // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
}

// CreateBusinessPost calls create_business_post: Publish a Google Business
// Profile post. Publishes an update post (topic STANDARD) on the client's own
// Google Business Profile listing through its connected Business Profile:
// summary (up to 1,500 characters) and an optional callToAction — BOOK, ORDER,
// SHOP, LEARN_MORE or SIGN_UP with a url, or CALL with none (it dials the
// listing's phone). dryRun: true publishes nothing and returns the exact
// localPosts body. The answer carries Google's name for the post and its state
// (PROCESSING until Google has reviewed it, LIVE, or REJECTED for its content
// policy). Refused with missing_scope or not_connected as
// update_business_profile is. get_business_posts reads the posts already
// there. Answers 'post this offer on our Google listing'. Free.
func (c *Client) CreateBusinessPost(ctx context.Context, in CreateBusinessPostInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_business_post", in)
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

// CreateDataSyncInput is the input to create_data_sync.
type CreateDataSyncInput struct {
	ClientID      string         `json:"clientId"`
	Destination   string         `json:"destination"`
	Name          string         `json:"name"`
	Query         map[string]any `json:"query"` // The same query query_data takes, without clientId: see describe_datasets.
	Mode          string         `json:"mode,omitempty"`
	WindowDays    *int64         `json:"windowDays,omitempty"`
	Cadence       string         `json:"cadence,omitempty"`
	RunHour       *int64         `json:"runHour,omitempty"`
	SpreadsheetID string         `json:"spreadsheetId,omitempty"`
	Sheet         string         `json:"sheet,omitempty"`
	BaseID        string         `json:"baseId,omitempty"`
	Table         string         `json:"table,omitempty"`
	ProjectID     string         `json:"projectId,omitempty"`
	DatasetID     string         `json:"datasetId,omitempty"`
	TableID       string         `json:"tableId,omitempty"`
}

// CreateDataSync calls create_data_sync: Sync data to a sheet, base or
// warehouse. Copies any dataset query (the query_data shape; see
// describe_datasets) to the client's Google Sheet (spreadsheetId, sheet),
// Airtable table (baseId, table — its fields named as the columns) or BigQuery
// table (projectId, datasetId, tableId — created on first run) on a schedule:
// daily, weekly or monthly at a UTC hour. replace swaps the contents each run;
// append adds rows, and with windowDays reads only the last N days ending
// yesterday so no day is added twice. Needs that destination connected first
// (get_integration_connect_link for Google, connect_integration for Airtable).
// Answers 'put our Search Console pages in a sheet every Monday', 'send the
// analytics to BigQuery daily'. Free.
func (c *Client) CreateDataSync(ctx context.Context, in CreateDataSyncInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_data_sync", in)
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

// CreateFunnelInput is the input to create_funnel.
type CreateFunnelInput struct {
	ClientID string           `json:"clientId"`
	Name     string           `json:"name"`
	Steps    []map[string]any `json:"steps"`
}

// CreateFunnel calls create_funnel: Define a funnel. Defines a funnel: 2 to 8
// steps (a page reached or an event sent) a visit should pass through in order
// — landing page, product, cart, thank-you. Measured within one visit, because
// the cookieless visitor id resets daily. Answers 'how many who see pricing go
// on to sign up'. Free.
func (c *Client) CreateFunnel(ctx context.Context, in CreateFunnelInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_funnel", in)
}

// CreateGoalInput is the input to create_goal.
type CreateGoalInput struct {
	ClientID   string `json:"clientId"`
	Name       string `json:"name"`
	Kind       string `json:"kind"`
	Match      string `json:"match"` // A path for a pageview ("/thanks", "/signup/*"; * matches anything) or an event name.
	ValueMinor *int64 `json:"valueMinor,omitempty"`
}

// CreateGoal calls create_goal: Define a conversion goal. Defines a goal for
// the cookieless analytics: reaching a page (kind pageview, match a path such
// as /thanks or /signup/*) or sending a custom event (kind event, match its
// name). Optionally worth a fixed value in minor units (2500 = 25.00); without
// one, an event's own 'value' property is summed. Counts every visit since the
// tag went in, not only from now. Answers 'count sign-ups as reaching
// /welcome'. Free.
func (c *Client) CreateGoal(ctx context.Context, in CreateGoalInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_goal", in)
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
// country are a different keyword. Each keyword may carry the metrics another
// tool already returned for it (volume, difficulty, CPC, intent, SERP
// features), stored with their read time so the list starts with numbers
// without buying them twice. Returns created: false with the reason when the
// name is already taken, rather than failing.
func (c *Client) CreateKeywordList(ctx context.Context, in CreateKeywordListInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_keyword_list", in)
}

// CreateLeadWidgetInput is the input to create_lead_widget.
type CreateLeadWidgetInput struct {
	Name           string   `json:"name"`
	Heading        string   `json:"heading,omitempty"`
	AccentColor    string   `json:"accentColor,omitempty"`    // The button colour, as hex.
	AllowedOrigins []string `json:"allowedOrigins,omitempty"` // Sites allowed to show the form, as origins (https://agency.example). Empty allows any site, for testing.
	NotifyEmails   []string `json:"notifyEmails,omitempty"`   // Who is emailed when a lead arrives. Empty keeps leads without telling anyone.
}

// CreateLeadWidget calls create_lead_widget: Create a lead-gen audit widget.
// Makes an audit form for the agency's own website: a visitor enters their URL
// and email, sees a short audit of the page on the spot, and lands in
// list_leads. Returns embedCode to paste into the site and frameUrl for the
// form alone. Answers 'put a free audit form on our site'. Limited per visitor
// and guarded against internal addresses. Free.
func (c *Client) CreateLeadWidget(ctx context.Context, in CreateLeadWidgetInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_lead_widget", in)
}

// CreateOutreachSequenceInput is the input to create_outreach_sequence.
type CreateOutreachSequenceInput struct {
	ClientID string           `json:"clientId"`
	Name     string           `json:"name"`
	Steps    []map[string]any `json:"steps"`
}

// CreateOutreachSequence calls create_outreach_sequence: Create an outreach
// sequence. Saves a sequence for the client: the first email and up to five
// follow-ups, each sent delayDays after the one before unless the prospect has
// replied (the first step's delay is ignored: it goes when send_outreach is
// called). Subjects and bodies may use {{site}} (the prospect's domain),
// {{page}} (its page, or the domain) and {{client}} (the client's name).
// Follow-ups go into the same thread. Answers 'set up a three-step outreach
// sequence'. Free.
func (c *Client) CreateOutreachSequence(ctx context.Context, in CreateOutreachSequenceInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_outreach_sequence", in)
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

// CreatePortfolioInput is the input to create_portfolio.
type CreatePortfolioInput struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Members     []map[string]any `json:"members,omitempty"`
}

// CreatePortfolio calls create_portfolio: Create a portfolio. Groups clients,
// or single sites of clients, under a name, to read their traffic, rankings,
// issues and AI visibility together with get_portfolio_metrics. Answers 'group
// the dental practices'. Names are unique in the agency. Every member must be
// a client the caller can reach.
func (c *Client) CreatePortfolio(ctx context.Context, in CreatePortfolioInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_portfolio", in)
}

// CreateProjectInput is the input to create_project.
type CreateProjectInput struct {
	ClientID     string `json:"clientId"`
	Name         string `json:"name"`
	Domain       string `json:"domain,omitempty"`
	LocationCode *int64 `json:"locationCode,omitempty"` // A DataForSEO location: a country, or a region, county or city (find one with search_locations).
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
	LocationCode *int64 `json:"locationCode,omitempty"` // A DataForSEO location: a country, or a region, county or city (find one with search_locations).
	LanguageCode string `json:"languageCode,omitempty"`
	RunHour      *int64 `json:"runHour,omitempty"` // The UTC hour scheduled runs happen at (default 3).
	RunDay       *int64 `json:"runDay,omitempty"`  // Weekly: the weekday, 0 Sunday – 6 Saturday. Monthly: the day of the month, 1–28.
	Depth        *int64 `json:"depth,omitempty"`   // Results read per keyword (default 100). Deeper costs more per check.
	Match        string `json:"match,omitempty"`   // domain: the exact host; subdomains: it and every subdomain; prefix: only addresses under urlPrefix (a folder or one page).
	URLPrefix    string `json:"urlPrefix,omitempty"`
}

// CreateRankTracker calls create_rank_tracker: Create a rank tracker. Starts
// tracking a client's positions on one engine, device and locale. Answers
// 'track this client on Bing too' or 'track mobile separately'. Defaults to
// the client's first project and site domain. Engine and locale cannot be
// changed later — they are part of every keyword's identity — so a different
// market is a new tracker. A locale can be a city, county or region, not only
// a country: search_locations finds the code. Cadence daily, weekly, monthly
// or manual, at a chosen UTC hour and day; depth 10–200; match the exact host,
// its subdomains, or a URL prefix.
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

// CreateReportLinkInput is the input to create_report_link.
type CreateReportLinkInput struct {
	ClientID      string `json:"clientId"`
	ReportID      string `json:"reportId"`
	ExpiresInDays *int64 `json:"expiresInDays,omitempty"`
	PeriodDays    *int64 `json:"periodDays,omitempty"`
}

// CreateReportLink calls create_report_link: Make a shareable report page. A
// signed, expiring, read-only link to one of a client's saved reports (from
// list_reports) as a web page under the agency's own branding, plus
// downloadUrl for the same report as a self-contained HTML file. Answers 'send
// the client a link to the report' or 'give me the report as HTML'. Covers the
// last periodDays days (default 30) as of now, fixed: the page does not change
// as time passes. Anyone holding the link can read it until it expires
// (expiresInDays, default 14, at most 90), so only make one when the user
// asks. The url is shown only in this answer and is never stored; linkId names
// it for list_report_links and revoke_report_link. revoke_report_links turns
// every earlier link off. Free.
func (c *Client) CreateReportLink(ctx context.Context, in CreateReportLinkInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_report_link", in)
}

// CreateTaskInput is the input to create_task.
type CreateTaskInput struct {
	ClientID       string         `json:"clientId"`
	Title          string         `json:"title"`
	Notes          string         `json:"notes,omitempty"`
	Status         string         `json:"status,omitempty"`
	AssigneeUserID string         `json:"assigneeUserId,omitempty"` // A user id from list_team_members (the member's userId). They must be able to open this client.
	DueDate        string         `json:"dueDate,omitempty"`        // YYYY-MM-DD.
	Link           map[string]any `json:"link,omitempty"`           // What the task is about: an audit issue (value: its ruleId from list_crawl_issues), a page (an absolute URL) or a keyword.
	CrawlID        string         `json:"crawlId,omitempty"`
}

// CreateTask calls create_task: Add a task. Adds a task to a client's list,
// optionally assigned, dated and linked to an audit issue, a page or a
// keyword. Answers 'remind us to rewrite the pricing page by Friday'. One open
// task per audit issue: a second is refused while the first is open.
func (c *Client) CreateTask(ctx context.Context, in CreateTaskInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_task", in)
}

// CreateTasksFromAuditInput is the input to create_tasks_from_audit.
type CreateTasksFromAuditInput struct {
	ClientID       string   `json:"clientId"`
	CrawlID        string   `json:"crawlId,omitempty"`
	RuleIDs        []string `json:"ruleIds,omitempty"` // Only these issues; null takes every issue of the chosen severities.
	Severities     []string `json:"severities,omitempty"`
	AssigneeUserID string   `json:"assigneeUserId,omitempty"` // A user id from list_team_members (the member's userId). They must be able to open this client.
	DueDate        string   `json:"dueDate,omitempty"`        // YYYY-MM-DD.
	Limit          *int64   `json:"limit,omitempty"`
	DryRun         *bool    `json:"dryRun,omitempty"` // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
}

// CreateTasksFromAudit calls create_tasks_from_audit: Turn audit issues into
// tasks. Makes one task per issue a crawl found — titled from the issue, with
// its page count and howToFix in the notes — for the chosen rule ids or every
// issue of the chosen severities (default critical and warning). Issues that
// already have an open task are skipped, so this is safe after every crawl.
// Findings accepted as intentional (accept_issue) are left out, and an issue
// they cover entirely makes no task. With path owners set (set_path_owners),
// each task names the owner of most of its pages, and its notes say how the
// pages split between owners. crawlId defaults to the latest completed crawl.
// dryRun: true makes nothing and lists wouldCreate (each task's title, notes,
// assignee and due date) and skipped. Answers 'put the audit on our to-do
// list'. Free.
func (c *Client) CreateTasksFromAudit(ctx context.Context, in CreateTasksFromAuditInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_tasks_from_audit", in)
}

// CreateWebhookInput is the input to create_webhook.
type CreateWebhookInput struct {
	ClientID string   `json:"clientId,omitempty"`
	URL      string   `json:"url"`
	Events   []string `json:"events"`
}

// CreateWebhook calls create_webhook: Subscribe a URL to events. Sends events
// to an https URL as they happen, signed: each POST carries
// Seolyticize-Signature: t=<unix>,v1=<hex HMAC-SHA256 of "<t>.<body>" with the
// secret>, plus Seolyticize-Event and Seolyticize-Delivery headers. Retried
// with backoff (1 min to a day) until a 2xx; replay_webhooks sends missed ones
// again. Events: crawl.completed, gsc.synced, rank.moved, backlink.new,
// backlink.lost, issue.new_critical, ai.citation.gained, ai.citation.lost,
// competitor.page.new, budget.threshold, alert.fired. clientId limits it to
// one client; null is every client. The secret is in the response once.
// Answers 'tell my system when a crawl finishes'. Free.
func (c *Client) CreateWebhook(ctx context.Context, in CreateWebhookInput) (json.RawMessage, error) {
	return c.Call(ctx, "create_webhook", in)
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

// DeleteDataSyncInput is the input to delete_data_sync.
type DeleteDataSyncInput struct {
	ClientID string `json:"clientId"`
	SyncID   string `json:"syncId"`
}

// DeleteDataSync calls delete_data_sync: Delete a data sync. Stops and removes
// a sync. What it already wrote stays where it is. Answers 'stop the Airtable
// sync'. Safe to repeat.
func (c *Client) DeleteDataSync(ctx context.Context, in DeleteDataSyncInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_data_sync", in)
}

// DeleteDigestScheduleInput is the input to delete_digest_schedule.
type DeleteDigestScheduleInput struct {
	ClientID string `json:"clientId"`
	DigestID string `json:"digestId"`
}

// DeleteDigestSchedule calls delete_digest_schedule: Stop a digest. Stops
// posting the weekly digest to that chat. Answers 'stop the Monday Slack
// digest'. Safe to repeat.
func (c *Client) DeleteDigestSchedule(ctx context.Context, in DeleteDigestScheduleInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_digest_schedule", in)
}

// DeleteFunnelInput is the input to delete_funnel.
type DeleteFunnelInput struct {
	ClientID string `json:"clientId"`
	FunnelID string `json:"funnelId"`
}

// DeleteFunnel calls delete_funnel: Delete a funnel. Deletes a funnel
// definition. Answers 'remove that funnel'. Safe to repeat.
func (c *Client) DeleteFunnel(ctx context.Context, in DeleteFunnelInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_funnel", in)
}

// DeleteGoalInput is the input to delete_goal.
type DeleteGoalInput struct {
	ClientID string `json:"clientId"`
	GoalID   string `json:"goalId"`
}

// DeleteGoal calls delete_goal: Delete a goal. Deletes a goal. The events it
// counted are untouched. Answers 'stop counting that goal'. Safe to repeat.
func (c *Client) DeleteGoal(ctx context.Context, in DeleteGoalInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_goal", in)
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

// DeletePortfolioInput is the input to delete_portfolio.
type DeletePortfolioInput struct {
	PortfolioID string `json:"portfolioId"`
}

// DeletePortfolio calls delete_portfolio: Delete a portfolio. Deletes a
// portfolio; its clients and sites are untouched. Needs access to every client
// in it. Answers 'drop that portfolio'.
func (c *Client) DeletePortfolio(ctx context.Context, in DeletePortfolioInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_portfolio", in)
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

// DeleteTaskInput is the input to delete_task.
type DeleteTaskInput struct {
	ClientID string `json:"clientId"`
	TaskID   string `json:"taskId"`
}

// DeleteTask calls delete_task: Delete a task. Deletes a task outright. Prefer
// update_task with status wont_do, which keeps the record. Answers 'remove
// that task'. Safe to repeat.
func (c *Client) DeleteTask(ctx context.Context, in DeleteTaskInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_task", in)
}

// DeleteWebhookInput is the input to delete_webhook.
type DeleteWebhookInput struct {
	WebhookID string `json:"webhookId"`
}

// DeleteWebhook calls delete_webhook: Delete a webhook. Stops a webhook for
// good, with its queued deliveries. Answers 'stop sending to that URL'. Free.
func (c *Client) DeleteWebhook(ctx context.Context, in DeleteWebhookInput) (json.RawMessage, error) {
	return c.Call(ctx, "delete_webhook", in)
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

// DescribeDatasetsInput is the input to describe_datasets.
type DescribeDatasetsInput struct {
}

// DescribeDatasets calls describe_datasets: What can be queried. The datasets
// query_data and export_data read — Search Console (every day kept),
// cookieless analytics events, crawl pages, issues and links, rank captures,
// backlinks, imported revenue (orders with their channel and landing page), AI
// assistant visits — with their columns, metrics, time column and how long
// each is kept. Answers 'what data is there to ask about'. Free.
func (c *Client) DescribeDatasets(ctx context.Context, in DescribeDatasetsInput) (json.RawMessage, error) {
	return c.Call(ctx, "describe_datasets", in)
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

// DisconnectIntegrationInput is the input to disconnect_integration.
type DisconnectIntegrationInput struct {
	ClientID string `json:"clientId"`
	Provider string `json:"provider"`
}

// DisconnectIntegration calls disconnect_integration: Disconnect a tool.
// Removes a connection and its sealed key, with the syncs, digests and task
// links that ran through it. Revoke the key at the provider too. Answers
// 'disconnect their Stripe'. Safe to repeat.
func (c *Client) DisconnectIntegration(ctx context.Context, in DisconnectIntegrationInput) (json.RawMessage, error) {
	return c.Call(ctx, "disconnect_integration", in)
}

// DraftAgentFilesInput is the input to draft_agent_files.
type DraftAgentFilesInput struct {
	ClientID string   `json:"clientId"`
	AITrain  string   `json:"aiTrain"`           // Whether the content may be used to train AI models.
	AIInput  string   `json:"aiInput,omitempty"` // Whether it may ground AI answers (retrieval).
	Search   string   `json:"search,omitempty"`  // Whether it may be indexed for search.
	URLs     []string `json:"urls,omitempty"`    // Pages whose forms to stub, besides the homepage.
}

// DraftAgentFiles calls draft_agent_files: Draft the site's files for AI
// agents. Drafts, for the client's site, the files AI agents read — nothing is
// published: robots.txt with a Content-Signal line (search, ai-input,
// ai-train) in every user-agent group, every Allow and Disallow kept exactly
// as it is, with aiTrain the user's own choice; an ai-catalog.json in the
// Agentic Resource Discovery 1.0 shape (specVersion 1.0, standard media types
// and placeholder representativeQueries) listing the agent card the site
// already has, with resources ARD has no type for (llms.txt, sitemaps, API
// catalogs) named in skipped; and WebMCP tool stubs (document.modelContext,
// falling back to navigator.modelContext) carrying readOnlyHint,
// consequentialHint and untrustedContentHint, one per real form found on the
// homepage and the pages given or its contact, booking and sign-up pages, with
// forms that send, buy or delete marked consequential. Answers 'write our
// Content-Signal lines', 'make us an ai-catalog.json' and 'how do we expose
// our forms to browser agents'. Free.
func (c *Client) DraftAgentFiles(ctx context.Context, in DraftAgentFilesInput) (json.RawMessage, error) {
	return c.Call(ctx, "draft_agent_files", in)
}

// DraftArticleInput is the input to draft_article.
type DraftArticleInput struct {
	ClientID        string   `json:"clientId"`
	Keyword         string   `json:"keyword"`
	RelatedKeywords []string `json:"relatedKeywords,omitempty"`
	WordCount       *int64   `json:"wordCount,omitempty"`    // Length to aim for. Default: the median of the ranking pages.
	Instructions    string   `json:"instructions,omitempty"` // Anything the writer should know: angle, product, what to avoid.
	TargetURL       string   `json:"targetUrl,omitempty"`    // The page this draft is for, when it rewrites one.
	Depth           *int64   `json:"depth,omitempty"`
	LocationCode    *int64   `json:"locationCode,omitempty"`  // Google location code (2840 United States, 2826 United Kingdom). Default: the client's project.
	LanguageCode    string   `json:"languageCode,omitempty"`  // Language code (en, es). Default: the client's project.
	MaxCostMicros   *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun          *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// DraftArticle calls draft_article: Draft an article from the brief. Writes a
// full article draft for a keyword, grounded in the content brief built from
// the pages that rank: the intent and format that win, the target length, the
// outline, the must-cover terms and the questions to answer, in the client's
// voice from its brief. Returns Markdown with H2/H3 headings, a meta title and
// description with pixel-length checks, and which must-cover terms the draft
// uses and misses (checked by us, not claimed). When the keyword registry
// already gives the keyword to another page, ownership warns 'this already
// belongs to /x' (targetUrl: the owner, to draft its rewrite instead). Answers
// 'write me a first draft for this keyword'. A paid Claude call on the
// client's budget, plus the brief's results page (free when anyone bought it
// in the last day). Run score_content on the result before publishing.
func (c *Client) DraftArticle(ctx context.Context, in DraftArticleInput) (json.RawMessage, error) {
	return c.Call(ctx, "draft_article", in)
}

// DraftClientBriefInput is the input to draft_client_brief.
type DraftClientBriefInput struct {
	ClientID      string `json:"clientId"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// DraftOutreachEmailsInput is the input to draft_outreach_emails.
type DraftOutreachEmailsInput struct {
	ClientID      string   `json:"clientId"`
	ProspectIDs   []string `json:"prospectIds"`
	Offer         string   `json:"offer"` // What the client offers: the replacement page, the data, the guest piece.
	OurURL        string   `json:"ourUrl,omitempty"`
	Sender        string   `json:"sender,omitempty"`        // Who signs it.
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// DraftOutreachEmails calls draft_outreach_emails: Draft outreach emails to
// link prospects. Drafts one short, personal pitch per prospect on the
// client's outreach list (up to 20): it names why this site was chosen — the
// dead page it links to, the thread, the article, the rivals it links to (from
// the prospect's reason and page) — and what is offered in its place (offer,
// and ourUrl when given), in the client's voice from its brief. Plain text
// with a subject; no templates stuffed with merge fields, no false flattery,
// nothing claimed that the prospect data does not show. Nothing is sent: copy
// them into your mail tool, then move the prospect to contacted with
// set_link_prospect_status. Answers 'write the outreach for these sites'. A
// metered Claude call on the client's budget; quote it with estimate_cost.
func (c *Client) DraftOutreachEmails(ctx context.Context, in DraftOutreachEmailsInput) (json.RawMessage, error) {
	return c.Call(ctx, "draft_outreach_emails", in)
}

// DraftReviewRepliesInput is the input to draft_review_replies.
type DraftReviewRepliesInput struct {
	ClientID        string `json:"clientId"`
	LocationID      string `json:"locationId"`
	Limit           *int64 `json:"limit,omitempty"`
	Voice           string `json:"voice,omitempty"`
	IncludeAnswered *bool  `json:"includeAnswered,omitempty"`
	Trade           string `json:"trade,omitempty"`
	MaxCostMicros   *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun          *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// DraftReviewReplies calls draft_review_replies: Draft replies to a location's
// reviews. Drafts owner replies to a location's Google reviews that have no
// reply yet (newest first, up to limit): thanks that name what the reviewer
// praised, and for a poor review an apology, what will change and an offline
// route — never arguing, never inventing facts, no keyword stuffing. voice
// sets the tone in a sentence. trade (detected from the business by default)
// holds a healthcare practice's replies to patient privacy (HIPAA: never
// confirm the reviewer is a patient, never mention care) and a legal
// practice's to client confidentiality, and each draft that still slips
// carries a caution. Nothing is posted: read them, then post_review_reply.
// canPost says whether the listing is connected so they can be posted from
// here. Answers 'reply to our reviews'. A metered Claude call on the client's
// budget; quote it with estimate_cost.
func (c *Client) DraftReviewReplies(ctx context.Context, in DraftReviewRepliesInput) (json.RawMessage, error) {
	return c.Call(ctx, "draft_review_replies", in)
}

// DraftSEOStrategyInput is the input to draft_seo_strategy.
type DraftSEOStrategyInput struct {
	ClientID      string   `json:"clientId"`
	Competitors   []string `json:"competitors,omitempty"`
	Instructions  string   `json:"instructions,omitempty"`
	Save          *bool    `json:"save,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// DraftSEOStrategy calls draft_seo_strategy: Draft an SEO strategy. Drafts a
// 12-month SEO strategy for the client and saves it as a written report (save:
// false returns it without saving): a plan for its business type (SaaS, local,
// ecommerce, publisher or agency, as get_seo_scorecard detects it) in four
// phases — weeks 1–4, weeks 5–12, weeks 13–24 and months 7–12 — a teardown of
// the top five competitors (competitors, else the brief's and the ones we see
// ranking beside it) from the data we hold, a KPI table at baseline, 3, 6 and
// 12 months with baselines from Search Console and targets from
// forecast_search_clicks, plus indexed pages (URL Inspection, against the
// pages meant to be indexed) and phone Core Web Vitals (Chrome UX Report p75
// against Google's good lines), the site structure from the latest crawl, the
// URL architecture and a schema plan by page type for the business type (no
// FAQPage, no self-serving ratings), and a calendar from plan_content (the
// stored plan, or a new one built when there is none). Every figure comes from
// the client's data and is named with its source; Claude writes the prose
// around them and invents none. instructions adds the agency's own direction.
// One metered Claude call: quote it with estimate_cost. Answers 'write us an
// SEO strategy' and 'what is the plan for this client this year'.
func (c *Client) DraftSEOStrategy(ctx context.Context, in DraftSEOStrategyInput) (json.RawMessage, error) {
	return c.Call(ctx, "draft_seo_strategy", in)
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
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// EstimateDomainTraffic calls estimate_domain_traffic: Estimate search traffic
// for domains. A vendor's current estimate of monthly organic search traffic
// and ranking keywords for up to 100 domains, bought now. get_batch_analysis
// gives the same traffic beside rank, backlinks and referring domains for up
// to 200. Answers 'how much search traffic do these competitors get'. Modelled
// from rankings, not measured visits. batch_domain_metrics reads what is
// already stored for free; use this when those figures are missing or stale. A
// paid lookup on the client's account: quote it with estimate_cost.
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

// ExplainResultInput is the input to explain_result.
type ExplainResultInput struct {
	ClientID      string `json:"clientId"`
	Tool          string `json:"tool"`
	Result        any    `json:"result"`
	Audience      string `json:"audience,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// ExplainResult calls explain_result: Explain a result in plain words. Reads
// any earlier tool result (pass the tool's name and its data) and says, in
// plain language for a client or for the team, what it shows, what matters in
// it and what to do next — for handing off to people who don't read JSON.
// Answers 'what does this mean for the client'. A metered Claude call on the
// client's budget; quote it with estimate_cost.
func (c *Client) ExplainResult(ctx context.Context, in ExplainResultInput) (json.RawMessage, error) {
	return c.Call(ctx, "explain_result", in)
}

// ExploreAIPromptInput is the input to explore_ai_prompt.
type ExploreAIPromptInput struct {
	ClientID      string   `json:"clientId,omitempty"`
	Prompt        string   `json:"prompt"`
	Engines       []string `json:"engines,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// ExportDataInput is the input to export_data.
type ExportDataInput struct {
	ClientID   string           `json:"clientId"`
	Dataset    string           `json:"dataset"`
	Dimensions []string         `json:"dimensions,omitempty"`
	Metrics    []string         `json:"metrics,omitempty"`
	Columns    []string         `json:"columns,omitempty"`
	Filters    []map[string]any `json:"filters,omitempty"`
	From       string           `json:"from,omitempty"`
	To         string           `json:"to,omitempty"`
	OrderBy    []map[string]any `json:"orderBy,omitempty"` // Sort by a selected dimension or metric: [{ "column": "sum(clicks)", "desc": true }].
	CrawlID    string           `json:"crawlId,omitempty"` // For crawl datasets: a crawl id or latest.
	Limit      *int64           `json:"limit,omitempty"`
	Format     string           `json:"format,omitempty"`
}

// ExportData calls export_data: Export stored data to a file. The same query
// as query_data, written to a file — CSV, JSON lines or Parquet, up to a
// million rows — with a download link valid for 24 hours, for BigQuery, a
// notebook or a spreadsheet. Answers 'give me all of our Search Console
// history' and 'export the crawl'. Free.
func (c *Client) ExportData(ctx context.Context, in ExportDataInput) (json.RawMessage, error) {
	return c.Call(ctx, "export_data", in)
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

// ExportSitemapInput is the input to export_sitemap.
type ExportSitemapInput struct {
	ClientID string `json:"clientId"`
}

// ExportSitemap calls export_sitemap: Write the sitemap from the latest audit.
// Writes sitemap.xml from the client's latest complete crawl: every 200 HTML
// page that is indexable and canonical to itself, deduplicated and sorted,
// with no lastmod (the crawl knows when it read a page, not when it changed).
// Past 50,000 URLs or 50 MB a file, it splits into sitemap-1.xml,
// sitemap-2.xml… under a sitemap.xml index whose locations assume the parts
// sit at the site's root. Returns a download link per file, valid for a day.
// Answers 'make us a clean sitemap' and 'what should our sitemap list'. Free.
func (c *Client) ExportSitemap(ctx context.Context, in ExportSitemapInput) (json.RawMessage, error) {
	return c.Call(ctx, "export_sitemap", in)
}

// FindJournalistRequestsInput is the input to find_journalist_requests.
type FindJournalistRequestsInput struct {
	ClientID  string   `json:"clientId"`
	Feeds     []string `json:"feeds,omitempty"`
	Keywords  []string `json:"keywords,omitempty"`
	Days      *int64   `json:"days,omitempty"`
	Limit     *int64   `json:"limit,omitempty"`
	AddToList *bool    `json:"addToList,omitempty"`
}

// FindJournalistRequests calls find_journalist_requests: Journalist requests
// for sources, from feeds. Reads RSS or Atom feeds of journalist requests for
// sources and keeps those matching the keywords (all of them when none are
// given), newest first, with the deadline where the request states one. With
// no feeds given it reads the public #journorequest, #journorequests and
// #prrequest feeds on Mastodon (mastodon.social and journa.host) and says so
// in defaultFeeds, so keywords alone work. Source of Sources sends its queries
// by email only, Qwoted's and Featured's requests need an account, and
// SourceBottle publishes no feed — for those, give an email-to-RSS address
// subscribed to one, or any niche request board's feed. addToList puts each
// request with a site of its own on the outreach list (needs client:manage);
// social posts have none. Answers 'which journalists want a source on our
// topic this week'. Free.
func (c *Client) FindJournalistRequests(ctx context.Context, in FindJournalistRequestsInput) (json.RawMessage, error) {
	return c.Call(ctx, "find_journalist_requests", in)
}

// FindLinkOpportunitiesInput is the input to find_link_opportunities.
type FindLinkOpportunitiesInput struct {
	ClientID       string   `json:"clientId"`
	Topic          string   `json:"topic"`
	Kinds          []string `json:"kinds,omitempty"`
	QueriesPerKind *int64   `json:"queriesPerKind,omitempty"`
	Depth          *int64   `json:"depth,omitempty"`
	ExcludeDomains []string `json:"excludeDomains,omitempty"`
	AddToList      *bool    `json:"addToList,omitempty"`
	LocationCode   *int64   `json:"locationCode,omitempty"`
	LanguageCode   string   `json:"languageCode,omitempty"`
	MaxCostMicros  *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun         *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// FindLinkOpportunities calls find_link_opportunities: Resource pages, best-of
// lists, roundups, podcasts and newsletters for a topic. Finds places that
// link out by design, from Google results for the search footprints link
// builders use: resource pages ("useful links", intitle:resources), "best X"
// lists, link roundups, podcasts inviting guests, and newsletters taking
// sponsors or guest pieces — each with the query that found it, its title,
// snippet and position. The client's own domain and the giant platforms are
// left out. addToList puts every page on the client's outreach list (as
// add_link_prospects would; needs client:manage). Answers 'which resource
// pages could list us', 'which podcasts take guests in our niche' and 'where
// can we sponsor a newsletter'. One results page per query: queriesPerKind ×
// kinds reads. A paid lookup on the client's account unless anyone read the
// same thing recently (spend says which); quote it with estimate_cost.
func (c *Client) FindLinkOpportunities(ctx context.Context, in FindLinkOpportunitiesInput) (json.RawMessage, error) {
	return c.Call(ctx, "find_link_opportunities", in)
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
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// FindToolsInput is the input to find_tools.
type FindToolsInput struct {
	Task  string `json:"task"`
	Limit *int64 `json:"limit,omitempty"`
}

// FindTools calls find_tools: Find the tools for a task. Given a task in plain
// words ('why did traffic drop', 'write an article about X', 'who links to my
// rivals'), returns the 5–10 tools for it, best first, each with what it does
// and whether it costs money — and the named recipe that covers the task, when
// one does. Use it instead of reading the whole catalogue. Answers 'which tool
// do I use for this'. Free.
func (c *Client) FindTools(ctx context.Context, in FindToolsInput) (json.RawMessage, error) {
	return c.Call(ctx, "find_tools", in)
}

// FindUnlinkedMentionsInput is the input to find_unlinked_mentions.
type FindUnlinkedMentionsInput struct {
	ClientID      string `json:"clientId"`
	Brand         string `json:"brand,omitempty"`
	Since         string `json:"since,omitempty"` // Only pages published since; default the last year.
	Limit         *int64 `json:"limit,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// FindUnlinkedMentions calls find_unlinked_mentions: Pages that name the brand
// but don't link to it. Finds pages that mention the client's brand (brand, or
// the brand terms from its brief) from DataForSEO's content index, one per
// site, then reads each (up to limit) and keeps those with no link to the
// client's site — the easiest links to win, since the writer already chose to
// talk about it. Each with its domain rank, date and the sentence that names
// the brand. Answers 'who mentions us without linking'. A paid DataForSEO read
// plus free page reads; quote it with estimate_cost.
func (c *Client) FindUnlinkedMentions(ctx context.Context, in FindUnlinkedMentionsInput) (json.RawMessage, error) {
	return c.Call(ctx, "find_unlinked_mentions", in)
}

// ForecastKeywordTrafficInput is the input to forecast_keyword_traffic.
type ForecastKeywordTrafficInput struct {
	ClientID     string           `json:"clientId"`
	Keywords     []map[string]any `json:"keywords"`
	Positions    []int64          `json:"positions,omitempty"`
	LocationCode *int64           `json:"locationCode,omitempty"` // Google location code (2840 United States, 2826 United Kingdom). Default: the client's project.
	LanguageCode string           `json:"languageCode,omitempty"` // Language code (en, es). Default: the client's project.
}

// ForecastKeywordTraffic calls forecast_keyword_traffic: What a keyword set is
// worth at target positions. Estimated organic clicks for a set of keywords at
// the positions aimed for (default 1, 3, 5 and 10): volume × the click curve
// for the market (measured by the panel where it can be, published otherwise,
// and said which), shaped month by month by each keyword's seasonality from
// stored monthly volumes, with a time-to-rank estimate from difficulty and the
// client's current position, and first-year clicks counting only from the
// month the page is likely to arrive. Totals for the set. keywords are
// strings, or { term, searchVolume, difficulty } when the figures are known.
// Answers 'what is this cluster worth if we rank'. Free: volumes and
// difficulty are read from what is stored, or given; a keyword never looked up
// is reported with unknown volume rather than zero.
func (c *Client) ForecastKeywordTraffic(ctx context.Context, in ForecastKeywordTrafficInput) (json.RawMessage, error) {
	return c.Call(ctx, "forecast_keyword_traffic", in)
}

// ForecastSearchClicksInput is the input to forecast_search_clicks.
type ForecastSearchClicksInput struct {
	ClientID    string           `json:"clientId"`
	HorizonDays *int64           `json:"horizonDays,omitempty"`
	Trends      []map[string]any `json:"trends,omitempty"`
	Scenario    *bool            `json:"scenario,omitempty"`
}

// ForecastSearchClicks calls forecast_search_clicks: The next 90 days of
// non-branded clicks. Forecasts daily non-branded Google clicks for the next
// 90 days (horizonDays) from the site's own history: level, a damped trend,
// the weekly rhythm and last year's seasonality (needs 13 months of history;
// with less, pass trends — Google Trends interest for the site's head term
// from get_google_trends, past_5_years — to stand in), with an 80% band.
// scenario adds 'what if we win the striking-distance set': every query at
// positions 4–15 moved to 3, using the site's own click-through curve, the
// gain arriving over 60 days. Answers 'how much traffic should we expect next
// quarter' and 'what is the striking-distance work worth'. Free.
func (c *Client) ForecastSearchClicks(ctx context.Context, in ForecastSearchClicksInput) (json.RawMessage, error) {
	return c.Call(ctx, "forecast_search_clicks", in)
}

// GenerateHreflangInput is the input to generate_hreflang.
type GenerateHreflangInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	XDefault string `json:"xDefault,omitempty"`
	Method   string `json:"method,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GenerateHreflang calls generate_hreflang: Write the hreflang for the site.
// Generates hreflang for every group of language versions in a crawl (tied by
// existing hreflang, or the same path under a language prefix such as /es/ or
// /de-at/), each with a self-reference and an x-default (xDefault names the
// fallback language; default the language of the unprefixed pages): <link>
// tags for groups under 50 versions, a Link HTTP header for PDFs, and an
// xhtml:link sitemap for groups of 50 or more (method forces one). Lists
// groups with two pages claiming one language, and pages whose language could
// not be told. Answers 'write our hreflang tags' and 'give me an hreflang
// sitemap'. Free.
func (c *Client) GenerateHreflang(ctx context.Context, in GenerateHreflangInput) (json.RawMessage, error) {
	return c.Call(ctx, "generate_hreflang", in)
}

// GenerateLlmsTxtInput is the input to generate_llms_txt.
type GenerateLlmsTxtInput struct {
	ClientID string `json:"clientId"`
	MaxLinks *int64 `json:"maxLinks,omitempty"`
	Summary  string `json:"summary,omitempty"` // The one-line summary; default the homepage's description.
}

// GenerateLlmsTxt calls generate_llms_txt: Write the site's llms.txt. Builds
// an llms.txt for the client's site (the llmstxt.org format: a title, a
// one-line summary, then sections of links with a description each) from its
// latest complete crawl — the pages meant to rank, strongest by internal
// PageRank, grouped by section, each with its own title and description. Says
// whether the site already publishes one and whether it differs. Nothing is
// published: put the markdown at publishAt. Answers 'make us an llms.txt'.
// Free.
func (c *Client) GenerateLlmsTxt(ctx context.Context, in GenerateLlmsTxtInput) (json.RawMessage, error) {
	return c.Call(ctx, "generate_llms_txt", in)
}

// GenerateSchemaInput is the input to generate_schema.
type GenerateSchemaInput struct {
	ClientID          string           `json:"clientId"`
	URL               string           `json:"url"`
	Type              string           `json:"type,omitempty"`              // A type to write even where the page would not suggest it.
	LocalBusinessType string           `json:"localBusinessType,omitempty"` // The schema.org LocalBusiness subtype, e.g. Dentist or LegalService.
	OpeningHours      []map[string]any `json:"openingHours,omitempty"`      // As the Business Profile lists them, HH:MM.
	AreaServed        []string         `json:"areaServed,omitempty"`        // For a service-area business: the towns or regions it serves.
	Shipping          map[string]any   `json:"shipping,omitempty"`          // The shop's shipping to one country, for Product offers.
	Returns           map[string]any   `json:"returns,omitempty"`           // The return policy: days null for no limit, 0 for no returns.
}

// GenerateSchema calls generate_schema: Write a page's structured data. Reads
// one page of the client's site and writes the JSON-LD it should carry: a
// BreadcrumbList from its path; on the homepage an Organization (with an @id,
// the social profiles the page itself links to, and the logo from its Open
// Graph image) and a WebSite; a LocalBusiness from the client's saved
// location, as the subtype its vertical calls for (from the page's words, or
// localBusinessType; never the deprecated Attorney), with the address split
// into a PostalAddress, an @id, branchOf when the client has several
// locations, openingHoursSpecification from openingHours, areaServed for a
// service-area business, and never a self-serving aggregateRating; an Article
// for a post, with dates and author only where the page states them. Nothing
// is invented — a field the page does not state is left out and named in notes
// — and types the page already has are listed. Also Product (price, currency,
// availability and brand from the page's product tags, with
// OfferShippingDetails and MerchantReturnPolicy from shipping and returns),
// VideoObject for an embedded YouTube or Vimeo video (Clip per chapter the
// page lists, else SeekToAction; BroadcastEvent for a live stream),
// ProfilePage on an author or team page, QAPage on a question page,
// DiscussionForumPosting on a forum thread, an ItemList of a hub's own pages,
// SoftwareSourceCode for a page linking its repository, HowTo (its ordered
// list of steps; Google retired the rich result), Service (for /services/
// pages), and on request (type) ProductGroup, FAQPage, SoftwareApplication,
// Organization, LocalBusiness or Article on any page. Every block is checked
// against Google's required and recommended properties for its type
// (validation), so it is clear whether it can earn a rich result as written.
// FAQPage is written only when asked for: Google stopped showing FAQ rich
// results for every site on 7 May 2026, and any benefit with assistants is
// unconfirmed; FAQPage and HowTo blocks carry richResultEligible: false with
// the reason. Answers 'what schema should this page have', 'write product
// schema for this page' and 'mark up our locations'. Free.
func (c *Client) GenerateSchema(ctx context.Context, in GenerateSchemaInput) (json.RawMessage, error) {
	return c.Call(ctx, "generate_schema", in)
}

// GetAgentReadinessInput is the input to get_agent_readiness.
type GetAgentReadinessInput struct {
	ClientID string `json:"clientId"`
	Render   *bool  `json:"render,omitempty"` // Render the homepage in a browser for the accessibility tree.
}

// GetAgentReadiness calls get_agent_readiness: Get agent readiness. Whether
// the client's site can be read, cited and used by AI assistants and browsing
// agents. The original findings (crawler access, rendering, structured data,
// speed) plus prioritised checks: P0 a made-up path returns a real 404, CLS ≤
// 0.1 (Chrome UX Report), robots.txt answers (a 5xx fails), and the homepage's
// primary content is there without JavaScript (under 50 words in an app shell
// fails, under 150 warns); P1 a Content-Signal line in every robots.txt group
// with valid syntax (search, ai-input, ai-train, each yes or no), llms.txt as
// Lighthouse checks it (an H1, a Markdown link, 50+ characters; a catch-all
// HTML page fails; /llms-full.txt reported), Markdown delivery (a .md twin, or
// Accept: text/markdown answered with Vary: Accept), no /.well-known path
// answering 200 with HTML, and no WebMCP tool telling the agent to skip the
// user's confirmation; P2 /.well-known/api-catalog and OAuth metadata where
// the site runs APIs, WebMCP tools registered with document.modelContext
// (navigator.modelContext alone warns), and an accessibility-tree heuristic
// (0–100: accessible names, valid roles, nothing focusable hidden) from the
// rendered homepage; P3 ai-catalog.json found in Lighthouse's order
// (robots.txt Agentmap, <link rel=ai-catalog>, Link header,
// /.well-known/ai-catalog.json) and validated against the ARD 1.0 rules (P1
// when a signalled catalog fails or errors), agent-card.json, /.well-known/ucp
// for stores validated against the UCP shape, and declarative WebMCP form
// attributes. Each check is pass, fail, warn, unmeasured or not applicable; no
// fix below P0 is recommended while a measured P0 fails. Answers 'why are
// assistants not citing us' and 'is our site ready for AI agents'. Free.
func (c *Client) GetAgentReadiness(ctx context.Context, in GetAgentReadinessInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_agent_readiness", in)
}

// GetAIAgentVisitsInput is the input to get_ai_agent_visits.
type GetAIAgentVisitsInput struct {
	ClientID string `json:"clientId"`
	By       string `json:"by,omitempty"`
	Days     *int64 `json:"days,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetAIAgentVisits calls get_ai_agent_visits: AI assistants reading the site.
// How often AI assistants read the client's pages on somebody's behalf —
// ChatGPT-User, ChatGPT's agent mode, Claude-User, Perplexity-User, Mistral —
// by assistant, operator, page or day, seen by the tag or a server-side
// collector and kept apart from visits by people. Answers 'which pages is
// ChatGPT reading' and 'are assistants citing us more'. Training and index
// crawlers are not here: analyse_server_log reads those from the logs. Free.
func (c *Client) GetAIAgentVisits(ctx context.Context, in GetAIAgentVisitsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_agent_visits", in)
}

// GetAIAnswerShareInput is the input to get_ai_answer_share.
type GetAIAnswerShareInput struct {
	ClientID string `json:"clientId"`
	Weeks    *int64 `json:"weeks,omitempty"`
	Engine   string `json:"engine,omitempty"` // One assistant (chatgpt, perplexity, gemini, claude, google-aio, …); default all.
}

// GetAIAnswerShare calls get_ai_answer_share: Share of AI answers per topic,
// week by week. For each topic of the tracked questions, week by week across
// assistants (or one): how many answers there were, how many mentioned or
// cited the client, its share of answers, and its share of every tracked
// brand's mentions in them (share of voice). Answers 'are we winning the
// pricing questions in ChatGPT' and 'is our AI share growing by topic'. Free.
func (c *Client) GetAIAnswerShare(ctx context.Context, in GetAIAnswerShareInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_answer_share", in)
}

// GetAIBrandMentionsInput is the input to get_ai_brand_mentions.
type GetAIBrandMentionsInput struct {
	ClientID       string           `json:"clientId"`
	Brands         []map[string]any `json:"brands"`
	Platform       string           `json:"platform,omitempty"`       // google: AI Overviews; chat_gpt: ChatGPT (US English only) — both from DataForSEO's database plus Seolyticize's samples. claude, gemini, perplexity, copilot, gro
	IncludeSampled *bool            `json:"includeSampled,omitempty"` // Add Seolyticize's own sampled answers to the client's tracked questions (run_ai_visibility_check), tagged source sampled.
	SampledDays    *int64           `json:"sampledDays,omitempty"`
	LocationCode   *int64           `json:"locationCode,omitempty"`  // Default: the client's project (2840 United States).
	LanguageCode   string           `json:"languageCode,omitempty"`  // Default: the client's project.
	MaxCostMicros  *int64           `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun         *bool            `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetAIBrandMentions calls get_ai_brand_mentions: Brand share of voice in AI
// answers. How often AI answers (Google AI Overviews, or ChatGPT) mention or
// cite each brand — the client and up to 9 competitors, each given as domains
// and/or brand keywords — across DataForSEO's database of real AI answers,
// with share of voice, the AI search volume behind the mentions, and the
// domains cited. No question tracking needed; works for any brand. Answers
// 'what is our share of voice in AI search' and 'who do AI answers recommend'.
// Rows carry source: dataforseo (real users' prompts, with AI search volume)
// or sampled (the client's tracked questions asked by Seolyticize, no volume);
// the two are never summed. For platforms other than google and chat_gpt only
// sampled rows exist and the call is free. A paid lookup on the client's
// account unless anyone read the same thing recently (spend says which); quote
// it with estimate_cost. (about $0.10 a call)
func (c *Client) GetAIBrandMentions(ctx context.Context, in GetAIBrandMentionsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_brand_mentions", in)
}

// GetAICitationGapInput is the input to get_ai_citation_gap.
type GetAICitationGapInput struct {
	ClientID      string   `json:"clientId"`
	OurURL        string   `json:"ourUrl"`
	PromptID      string   `json:"promptId,omitempty"`
	Question      string   `json:"question,omitempty"`
	CitedURLs     []string `json:"citedUrls,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetAICitationGap calls get_ai_citation_gap: Why assistants cite others and
// not this page. For a question assistants answer without citing the client:
// reads the pages they did cite and the client's page (ourUrl), and says what
// the cited pages cover that the client's does not — the terms, the sections,
// the questions they answer — with a 0–100 coverage score and the fixes, as
// content scoring does against search results. The cited pages come from a
// tracked question's stored answers (promptId, free), from citedUrls (free),
// or by asking the ChatGPT app and Perplexity now (question, paid). Answers
// 'why does ChatGPT cite them and not us' and 'what should this page add to be
// cited'. Quote a live question with estimate_cost.
func (c *Client) GetAICitationGap(ctx context.Context, in GetAICitationGapInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_citation_gap", in)
}

// GetAICitationHistoryInput is the input to get_ai_citation_history.
type GetAICitationHistoryInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
	Engine   string `json:"engine,omitempty"` // One assistant (chatgpt, perplexity, gemini, claude, google-aio, …); default all.
	Limit    *int64 `json:"limit,omitempty"`
}

// GetAICitationHistory calls get_ai_citation_history: Which of our pages
// assistants cite, over time. The client's pages cited in the tracked
// questions' AI answers over the last N days: per page, how many answers cited
// it, by which assistants, for which questions, first and last seen, and
// citations per week — and, per question and assistant, the pages gained and
// lost between its last two checks. Answers 'which pages does ChatGPT cite',
// 'did we lose a citation this week' and 'which content earns AI citations'.
// Free.
func (c *Client) GetAICitationHistory(ctx context.Context, in GetAICitationHistoryInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_citation_history", in)
}

// GetAICitedSourcesInput is the input to get_ai_cited_sources.
type GetAICitedSourcesInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
	Engine   string `json:"engine,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	Topic    string `json:"topic,omitempty"`   // Only questions in this topic.
	Segment  string `json:"segment,omitempty"` // neutral: only questions naming no tracked brand (who gets recommended unprompted); branded: only those naming one.
}

// GetAICitedSources calls get_ai_cited_sources: Who owns the pages AI answers
// cite. The pages assistants cited for the client's tracked questions over the
// last N days (default 30), each domain classed by owner — own (the client's
// site), competitor (a tracked competitor's site, named), ugc (Reddit, Quora,
// YouTube, forums, social) or third_party (everyone else) — with each class's
// share of citations, each competitor's share, and the domains in order with
// their page type, questions, assistants and most-cited page. Counted once per
// answer. Filter by topic, segment or one assistant. Answers 'whose content
// are AI answers built from' and 'where should we get mentioned'. Free.
func (c *Client) GetAICitedSources(ctx context.Context, in GetAICitedSourcesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_cited_sources", in)
}

// GetAICrawlerAccessInput is the input to get_ai_crawler_access.
type GetAICrawlerAccessInput struct {
	ClientID string `json:"clientId"`
	Pages    *int64 `json:"pages,omitempty"`
	Log      string `json:"log,omitempty"` // An access log in combined format, as analyse_server_log takes it.
}

// GetAICrawlerAccess calls get_ai_crawler_access: Can each AI crawler reach
// and read the pages that matter. Per AI agent — GPTBot, OAI-SearchBot,
// ChatGPT-User, ClaudeBot, Claude-SearchBot, Claude-User, PerplexityBot,
// Perplexity-User, Google-Extended, Applebot-Extended, Meta-ExternalAgent,
// CCBot — for the client's top pages by Search Console clicks: which
// robots.txt allows and blocks, whether the server or CDN turns it away when
// the homepage is fetched under its name, which pages need JavaScript to show
// their content (from the crawl's rendering, which AI crawlers do not do),
// and, given an access log, how often it came, how many pages it read and how
// many requests failed. A verdict (open, partial, blocked) and fixes per
// agent, its purpose (training, search, user) said. Answers 'can ChatGPT read
// our site' and 'are we blocking AI search by accident'. Free.
func (c *Client) GetAICrawlerAccess(ctx context.Context, in GetAICrawlerAccessInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_crawler_access", in)
}

// GetAICrawlerActivityInput is the input to get_ai_crawler_activity.
type GetAICrawlerActivityInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
}

// GetAICrawlerActivity calls get_ai_crawler_activity: What AI crawlers read.
// From the client's connected logs: which pages GPTBot, ClaudeBot,
// PerplexityBot and every other AI agent fetched and how often — split by
// purpose (training crawls, search indexing, and fetches made for a user right
// now, the nearest a log gets to a citation) — with verified counts, error
// shares and when each was last seen. Google-Extended is a robots.txt opt-out
// token that Googlebot reads, not a crawler, so its robots.txt status is
// reported instead. Answers 'is ChatGPT reading our site', 'which pages does
// Perplexity fetch' and 'are AI crawlers hitting errors'. Free.
func (c *Client) GetAICrawlerActivity(ctx context.Context, in GetAICrawlerActivityInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_crawler_activity", in)
}

// GetAIModeAnswerInput is the input to get_ai_mode_answer.
type GetAIModeAnswerInput struct {
	ClientID      string `json:"clientId"`
	Keyword       string `json:"keyword"`
	Device        string `json:"device,omitempty"`
	LocationCode  *int64 `json:"locationCode,omitempty"`  // Default: the client's project (2840 United States).
	LanguageCode  string `json:"languageCode,omitempty"`  // Default: the client's project.
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetAIModeAnswer calls get_ai_mode_answer: Google AI Mode for a keyword. What
// Google's AI Mode answers for a keyword, with every source it cites and
// whether the client is among them. Answers 'are we cited in AI Mode' and 'who
// does AI Mode trust for this'. A paid lookup on the client's account unless
// anyone read the same thing recently (spend says which); quote it with
// estimate_cost.
func (c *Client) GetAIModeAnswer(ctx context.Context, in GetAIModeAnswerInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_mode_answer", in)
}

// GetAIOverviewCitationsInput is the input to get_ai_overview_citations.
type GetAIOverviewCitationsInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId,omitempty"` // Default: every tracker of the client.
	Feature   string `json:"feature,omitempty"`
	Days      *int64 `json:"days,omitempty"`
}

// GetAIOverviewCitations calls get_ai_overview_citations: Who the AI Overview
// cites. For each tracked keyword: how often Google showed an AI Overview (or
// a featured snippet, or People Also Ask — feature), whether the client is
// cited, who is, and how that changed; plus the domains cited most across the
// keywords — the outreach and content targets. Answers 'are we in the AI
// Overview' and 'who does Google trust on this topic'. From the results pages
// already captured; free.
func (c *Client) GetAIOverviewCitations(ctx context.Context, in GetAIOverviewCitationsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_overview_citations", in)
}

// GetAIPromptResearchInput is the input to get_ai_prompt_research.
type GetAIPromptResearchInput struct {
	ClientID       string `json:"clientId"`
	Keyword        string `json:"keyword"`
	Platform       string `json:"platform,omitempty"` // google: AI Overviews; chat_gpt: ChatGPT (US English only) — both from DataForSEO's database plus Seolyticize's samples. claude, gemini, perplexity, copilot, gro
	Limit          *int64 `json:"limit,omitempty"`
	IncludeSampled *bool  `json:"includeSampled,omitempty"` // Add Seolyticize's own sampled answers to the client's tracked questions (run_ai_visibility_check), tagged source sampled.
	SampledDays    *int64 `json:"sampledDays,omitempty"`
	LocationCode   *int64 `json:"locationCode,omitempty"`  // Default: the client's project (2840 United States).
	LanguageCode   string `json:"languageCode,omitempty"`  // Default: the client's project.
	MaxCostMicros  *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun         *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetAIPromptResearch calls get_ai_prompt_research: Real questions people ask
// AI about a topic. Questions real people put to AI assistants that contain a
// keyword, from DataForSEO's database of AI answers: each with its AI search
// volume, an excerpt of the answer, the sources cited, the brands named, the
// follow-up searches the assistant ran (fan-out), and whether the client is
// cited. Answers 'what do people ask ChatGPT about this' and 'which prompts
// should we track'. Rows carry source: dataforseo (real users' prompts, with
// AI search volume) or sampled (the client's tracked questions asked by
// Seolyticize, no volume); the two are never summed. For platforms other than
// google and chat_gpt only sampled rows exist and the call is free. A paid
// lookup on the client's account unless anyone read the same thing recently
// (spend says which); quote it with estimate_cost. (about $0.10 a call)
func (c *Client) GetAIPromptResearch(ctx context.Context, in GetAIPromptResearchInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_prompt_research", in)
}

// GetAIReferralsInput is the input to get_ai_referrals.
type GetAIReferralsInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetAIReferrals calls get_ai_referrals: Visits AI assistants send. Visits
// from AI assistants — ChatGPT (chatgpt.com, chat.openai.com), Perplexity,
// Gemini, Copilot, Claude, Google AI Mode (where the referrer shows it), Grok,
// Meta AI, DeepSeek and others — as an AI channel: from the cookieless
// analytics, visits by assistant, their share of all visits, the landing pages
// they arrive on and the conversions they bring (goal events); from GA4,
// sessions by source / medium, landing pages and key events and revenue for
// the same sources. Answers 'how much traffic does ChatGPT send us', 'which
// pages do assistants send people to' and 'do AI visitors convert'. Free.
func (c *Client) GetAIReferrals(ctx context.Context, in GetAIReferralsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_referrals", in)
}

// GetAISearchVolumeInput is the input to get_ai_search_volume.
type GetAISearchVolumeInput struct {
	ClientID      string   `json:"clientId"`
	Keywords      []string `json:"keywords"`
	LocationCode  *int64   `json:"locationCode,omitempty"`  // Default: the client's project (2840 United States).
	LanguageCode  string   `json:"languageCode,omitempty"`  // Default: the client's project.
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetAISearchVolume calls get_ai_search_volume: How often a topic is asked of
// AI. Estimated monthly AI-assistant search volume for up to 1,000 keywords,
// with 12 months of history. Answers 'how much is this asked in AI search' and
// 'which topics matter in AI answers'. A paid lookup on the client's account
// unless anyone read the same thing recently (spend says which); quote it with
// estimate_cost.
func (c *Client) GetAISearchVolume(ctx context.Context, in GetAISearchVolumeInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_search_volume", in)
}

// GetAITopCitedDomainsInput is the input to get_ai_top_cited_domains.
type GetAITopCitedDomainsInput struct {
	ClientID       string `json:"clientId"`
	Topic          string `json:"topic"`
	Platform       string `json:"platform,omitempty"` // google: AI Overviews; chat_gpt: ChatGPT (US English only) — both from DataForSEO's database plus Seolyticize's samples. claude, gemini, perplexity, copilot, gro
	Limit          *int64 `json:"limit,omitempty"`
	IncludeSampled *bool  `json:"includeSampled,omitempty"` // Add Seolyticize's own sampled answers to the client's tracked questions (run_ai_visibility_check), tagged source sampled.
	SampledDays    *int64 `json:"sampledDays,omitempty"`
	LocationCode   *int64 `json:"locationCode,omitempty"`  // Default: the client's project (2840 United States).
	LanguageCode   string `json:"languageCode,omitempty"`  // Default: the client's project.
	MaxCostMicros  *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun         *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetAITopCitedDomains calls get_ai_top_cited_domains: Who AI answers cite on
// a topic. The domains AI answers cite most for questions about a topic, with
// how often and the AI search volume behind it, the client's own marked — the
// outreach, PR and content targets for AI visibility. Answers 'which sites
// does AI trust on this topic'. Rows carry source: dataforseo (real users'
// prompts, with AI search volume) or sampled (the client's tracked questions
// asked by Seolyticize, no volume); the two are never summed. For platforms
// other than google and chat_gpt only sampled rows exist and the call is free.
// A paid lookup on the client's account unless anyone read the same thing
// recently (spend says which); quote it with estimate_cost. (about $0.10 a
// call)
func (c *Client) GetAITopCitedDomains(ctx context.Context, in GetAITopCitedDomainsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_top_cited_domains", in)
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
// the pages it cited, where each tracked brand appears in it, and how it
// described each (sentiment by brand, empty when not judged). found is false
// when that answer was never captured or has aged out (answers are kept 13
// months). Answers 'what exactly did ChatGPT say'. Free.
func (c *Client) GetAIVisibilityAnswer(ctx context.Context, in GetAIVisibilityAnswerInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_visibility_answer", in)
}

// GetAIVisibilityResultsInput is the input to get_ai_visibility_results.
type GetAIVisibilityResultsInput struct {
	ClientID   string `json:"clientId"`
	CapturedAt string `json:"capturedAt,omitempty"`
	Topic      string `json:"topic,omitempty"`   // Only questions in this topic.
	Segment    string `json:"segment,omitempty"` // neutral: only questions naming no tracked brand (who gets recommended unprompted); branded: only those naming one.
}

// GetAIVisibilityResults calls get_ai_visibility_results: AI visibility
// results, answer by answer. One check (the latest unless capturedAt names one
// from get_ai_visibility_runs), as a row per question and assistant: the
// question's topic and whether it is branded, whether the client was named,
// where in the answer (character offset; earlier is more prominent), whether
// its site was cited, how the answer described it (clientSentiment positive /
// neutral / negative with a one-line note; null when not named or not judged),
// and which other brands were named, in order. summary gives mention rate,
// citations and sentiment overall, neutral against branded, and per topic.
// Filter by topic or segment. Answers 'which questions are we missing from'
// and 'how do assistants talk about us'. Free.
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
// null when no brand is marked as the client's), and the same rate on neutral
// questions only (neutralMentionRate), the one branded questions cannot
// flatter. Answers 'when did we last check, and how did it go'. Free.
func (c *Client) GetAIVisibilityRuns(ctx context.Context, in GetAIVisibilityRunsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_visibility_runs", in)
}

// GetAIVisibilitySentimentInput is the input to get_ai_visibility_sentiment.
type GetAIVisibilitySentimentInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
	Topic    string `json:"topic,omitempty"`   // Only questions in this topic.
	Segment  string `json:"segment,omitempty"` // neutral: only questions naming no tracked brand (who gets recommended unprompted); branded: only those naming one.
}

// GetAIVisibilitySentiment calls get_ai_visibility_sentiment: How AI answers
// describe each brand. Sentiment over the last N days (default 30): for every
// tracked brand, how many answers that named it described it positively,
// neutrally or negatively, net score ((positive − negative) / judged) and how
// many were never judged; the client's split per assistant; and the latest
// negative and positive descriptions of the client, each with its question and
// assistant. Answers 'do assistants recommend us or warn people off'. Free;
// unjudged answers can be filled in with assess_ai_visibility_sentiment.
func (c *Client) GetAIVisibilitySentiment(ctx context.Context, in GetAIVisibilitySentimentInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_visibility_sentiment", in)
}

// GetAIVisibilitySetupInput is the input to get_ai_visibility_setup.
type GetAIVisibilitySetupInput struct {
	ClientID string `json:"clientId"`
}

// GetAIVisibilitySetup calls get_ai_visibility_setup: AI visibility setup.
// Reads what is tracked for a client: the questions (with their market, topic,
// and whether each is branded — names a tracked brand — or neutral, worked out
// automatically unless overridden), the topics and the branded/neutral split,
// the brands looked for (which one is the client's own), the schedule
// (cadence, which assistants, last run, whether sentiment is judged), and
// which assistants this agency can ask at all and how (its own key or its
// proxy pool). Answers 'what are we measuring in AI assistants'. Free.
func (c *Client) GetAIVisibilitySetup(ctx context.Context, in GetAIVisibilitySetupInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_visibility_setup", in)
}

// GetAIVisibilityTrendInput is the input to get_ai_visibility_trend.
type GetAIVisibilityTrendInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
	Topic    string `json:"topic,omitempty"`   // Only questions in this topic.
	Segment  string `json:"segment,omitempty"` // neutral: only questions naming no tracked brand (who gets recommended unprompted); branded: only those naming one.
}

// GetAIVisibilityTrend calls get_ai_visibility_trend: Is AI visibility
// improving. The client's own mention rate over the last N days (default 28)
// against the N days before, overall and per assistant, counted by answer,
// with how many answers each side rests on. Filter by topic, or segment
// neutral for the rate on questions naming no brand. Answers 'are we being
// recommended more than last month'. Free.
func (c *Client) GetAIVisibilityTrend(ctx context.Context, in GetAIVisibilityTrendInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ai_visibility_trend", in)
}

// GetAnchorDistributionInput is the input to get_anchor_distribution.
type GetAnchorDistributionInput struct {
	ClientID      string   `json:"clientId"`
	Domain        string   `json:"domain,omitempty"` // Any domain; default the client's own.
	Source        string   `json:"source,omitempty"`
	Limit         *int64   `json:"limit,omitempty"`
	MoneyTerms    []string `json:"moneyTerms,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetAnchorDistribution calls get_anchor_distribution: Anchor text mix. The
// anchor texts a domain is linked with, each classed as branded, a bare URL,
// generic (click here, read more), empty (images) or other (usually
// keyword-rich), and the mix by referring domains, with bands: each share
// against its natural range — branded 30–50% (flagged under 15%), URL 15–25%,
// generic 10–20%, exact match 3–10% (flagged over 15%), partial match 5–15%
// (flagged over 25%); exact and partial are told apart from other only with
// moneyTerms, the commercial phrases targeted. Answers 'does our link profile
// look natural' and 'what do they rank with'. source index is free: our own
// link index (our crawls, client crawls' outlinks, Common Crawl, Bing). auto
// answers from the index when it holds 50+ links to the domain (or, for
// referring domains, a complete pass of Common Crawl's domain graph), and
// otherwise asks the vendor — a paid lookup on the client's account unless
// anyone bought it in the last week; coverage says which answered. vendor
// always asks the vendor.
func (c *Client) GetAnchorDistribution(ctx context.Context, in GetAnchorDistributionInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_anchor_distribution", in)
}

// GetAnchorOverlapInput is the input to get_anchor_overlap.
type GetAnchorOverlapInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	Limit    *int64 `json:"limit,omitempty"`
}

// GetAnchorOverlap calls get_anchor_overlap: One anchor, several pages.
// Descriptive anchor texts used for links to more than one page — the site
// telling search engines that two pages are about the same words, which splits
// their ranking (cannibalization through links). Generic anchors (read more,
// click here) are left out. Answers 'are our internal links confusing Google'.
// Free.
func (c *Client) GetAnchorOverlap(ctx context.Context, in GetAnchorOverlapInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_anchor_overlap", in)
}

// GetAppKeywordGapInput is the input to get_app_keyword_gap.
type GetAppKeywordGapInput struct {
	ClientID      string   `json:"clientId"`
	Store         string   `json:"store"`
	AppID         string   `json:"appId"` // App Store: the numeric id (id123456789 or 123456789). Google Play: the package name.
	Competitors   []string `json:"competitors,omitempty"`
	LocationCode  *int64   `json:"locationCode,omitempty"`
	LanguageCode  string   `json:"languageCode,omitempty"`
	Limit         *int64   `json:"limit,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetAppKeywordGap calls get_app_keyword_gap: Keywords competitor apps rank
// for and ours does not. The store search keywords competitor apps rank for
// and the app does not, with search volume and each competitor's rank, from
// DataForSEO Labs (keywords for app; US English is what the vendor covers).
// With no competitors given, the vendor's own list of the app's competitors
// (top five) is used. Difficulty is not published for app keywords, so the
// best competitor rank stands in. Answers 'which keywords do rival apps rank
// for that we don't'. One read per app. A paid lookup on the client's account
// (spend says how much); quote it with estimate_cost.
func (c *Client) GetAppKeywordGap(ctx context.Context, in GetAppKeywordGapInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_app_keyword_gap", in)
}

// GetAppMetricsInput is the input to get_app_metrics.
type GetAppMetricsInput struct {
	ClientID      string   `json:"clientId"`
	Engine        string   `json:"engine"`
	AppIDs        []string `json:"appIds"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetAppMetrics calls get_app_metrics: How visible apps are in an app store.
// For up to 1,000 apps in the App Store (numeric ids, from /id… in the store
// URL) or Google Play (package names): how many search results pages each
// appears on, how many at #1, #2–3, #4–10 and #11–100, and the total search
// volume of those keywords, from DataForSEO Labs (US English only). Answers
// 'how visible is our app against these competitors'. get_ranked_keywords with
// engine app_store or google_play lists one app's keywords. A paid lookup on
// the client's account unless anyone read the same thing recently (spend says
// which); quote it with estimate_cost.
func (c *Client) GetAppMetrics(ctx context.Context, in GetAppMetricsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_app_metrics", in)
}

// GetAuditPerformanceInput is the input to get_audit_performance.
type GetAuditPerformanceInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest".
}

// GetAuditPerformance calls get_audit_performance: Speed in the audit. How
// fast the audited site is, as recorded when the crawl finished. Lab: a
// Lighthouse run (mobile) of the homepage and one page of each of the largest
// URL templates, with performance, accessibility, SEO and best-practices
// scores (0–100), LCP, TBT and CLS — one synthetic load each, from a Google
// datacentre. Field: the origin's Core Web Vitals from the Chrome UX Report —
// LCP, INP and CLS at the 75th percentile of real visits over 28 days, per
// phone and desktop, graded and with Google's pass/fail. Each page also
// carries diagnostics (where the time went: the LCP element and its subparts —
// TTFB, resource load delay, resource load duration, element render delay —
// DOM size, long tasks, main-thread work by kind, script execution,
// render-blocking requests, unused JavaScript and CSS, third parties, and the
// INP breakdown when the run interacted, which a PageSpeed run never does),
// and lab.fixes merges the pages' findings into one list ranked by estimated
// saving, with the pages each affects. Missing figures say why (no key, too
// few visits). crawlId defaults to the latest complete crawl. Answers 'is the
// site fast', 'which template is slow', 'what should we fix first for speed'
// and 'do we pass Core Web Vitals'. Free.
func (c *Client) GetAuditPerformance(ctx context.Context, in GetAuditPerformanceInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_audit_performance", in)
}

// GetAuditSummaryInput is the input to get_audit_summary.
type GetAuditSummaryInput struct {
	ClientID      string `json:"clientId"`
	CrawlID       string `json:"crawlId,omitempty"`
	IndexableOnly *bool  `json:"indexableOnly,omitempty"`
}

// GetAuditSummary calls get_audit_summary: Get site audit summary. Technical
// issues one crawl found, grouped by rule: each with its name, severity,
// category, how many pages it affects and howToFix. Answers 'what is broken on
// the site'. crawlId defaults to the latest completed crawl. Then
// list_issue_pages shows which pages an issue is on. indexableOnly (default
// true) leaves out findings on pages search engines will not index as
// themselves — noindexed, not answering 200, or canonicalised elsewhere —
// except the findings that explain why; pass false to see everything. Findings
// accepted as intentional (accept_issue) are taken out of the counts and
// listed apart in accepted, each with its reason; with path owners set
// (set_path_owners), each issue says how its pages split between owners
// (owners) and byOwner totals them.
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

// GetBacklinkHealthInput is the input to get_backlink_health.
type GetBacklinkHealthInput struct {
	ClientID        string         `json:"clientId"`
	MoneyTerms      []string       `json:"moneyTerms,omitempty"`
	SpamScores      map[string]any `json:"spamScores,omitempty"`      // Domain → DataForSEO spam score, from get_backlink_spam_scores.
	FetchSpamScores *bool          `json:"fetchSpamScores,omitempty"` // True to buy the vendor's spam score for every linking site inside this call (one paid request of up to 1,000 sites, the price of get_backlink_spam_scores; sites
	Industry        string         `json:"industry,omitempty"`
	MaxCostMicros   *int64         `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun          *bool          `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetBacklinkHealth calls get_backlink_health: How healthy is the link
// profile. Scores the client's link profile 0–100 (a heuristic, said so) from
// seven factors, each good, warn or critical with its figures and the source
// they came from (the vendor's synced index, our own link index, the domain
// graph): referring domains (20: over 100 good, under 20 critical), quality
// distribution (20: the linking sites' authority, and domain diversity — one
// site over 10% of the links warns, over 25% is critical), anchor naturalness
// (15: the anchor bands of get_anchor_distribution), toxic ratio (20: the
// share of linking sites get_toxic_link_risk calls toxic), velocity (10: a
// fall in referring domains of over 20% a quarter is critical), follow ratio
// (5: over 60% followed good, under 40% critical) and geographic relevance
// (10: the linking sites' TLD and country mix against the client's market; 80%
// or more from other countries flags a private network). Each factor carries
// its source's confidence (DataForSEO 1.0, our link index 0.7, Common Crawl
// 0.5), which scales its weight; a number only when at least four factors have
// data, and never from presence-only sources alone (NOT ASSESSED); otherwise
// INSUFFICIENT DATA (X/7). industry (saas, ecommerce, local, publisher,
// agency) judges the anchors against that industry's natural ranges. Also the
// profile's toxic signals: 50+ directory links, linking pages with 10,000+
// outbound links, more than 50 links from one site, linking pages under 100
// words, reciprocal links. Pass moneyTerms to tell exact and partial-match
// anchors apart. Spam scores sharpen the toxic ratio: fetchSpamScores: true
// buys them inside this call (one paid DataForSEO request of up to 1,000
// linking sites, the price of get_backlink_spam_scores; sites scored in the
// last 30 days are reused free; quote it with estimate_cost or dryRun first),
// or pass spamScores you already have. Answers 'how healthy are our backlinks'
// and 'is our link profile at risk'. Free without fetchSpamScores.
func (c *Client) GetBacklinkHealth(ctx context.Context, in GetBacklinkHealthInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_backlink_health", in)
}

// GetBacklinkHistoryInput is the input to get_backlink_history.
type GetBacklinkHistoryInput struct {
	ClientID      string `json:"clientId"`
	Target        string `json:"target,omitempty"`        // A domain (competitor.com), a subfolder (competitor.com/blog/) or a page URL. Default: the client's own domain.
	From          string `json:"from,omitempty"`          // Default: two years back.
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetBacklinkHistory calls get_backlink_history: A domain's backlinks over
// time. Month by month for any domain: backlinks, referring domains, new and
// lost backlinks and referring domains, and rank, from DataForSEO's index
// (back to 2019; new and lost from May 2021). Answers 'is their link profile
// growing', 'when did we lose those links'. A paid lookup on the client's
// account unless anyone read the same thing recently (spend says which); quote
// it with estimate_cost.
func (c *Client) GetBacklinkHistory(ctx context.Context, in GetBacklinkHistoryInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_backlink_history", in)
}

// GetBacklinkSpamScoresInput is the input to get_backlink_spam_scores.
type GetBacklinkSpamScoresInput struct {
	ClientID      string   `json:"clientId"`
	Domains       []string `json:"domains"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// GetBacklinksInput is the input to get_backlinks.
type GetBacklinksInput struct {
	ClientID      string `json:"clientId"`
	Domain        string `json:"domain,omitempty"` // Any domain; default the client's own.
	Source        string `json:"source,omitempty"`
	Limit         *int64 `json:"limit,omitempty"`
	Mode          string `json:"mode,omitempty"`
	Days          *int64 `json:"days,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetBacklinks calls get_backlinks: Backlinks, new and lost. The links to a
// domain, page to page, with anchor and follow status; mode new lists links
// first seen in the last days, lost those no longer seen since. Answers 'which
// links did we gain or lose this month' and 'what links does this competitor
// have'. source index is free: our own link index (our crawls, client crawls'
// outlinks, Common Crawl, Bing). auto answers from the index when it holds 50+
// links to the domain (or, for referring domains, a complete pass of Common
// Crawl's domain graph), and otherwise asks the vendor — a paid lookup on the
// client's account unless anyone bought it in the last week; coverage says
// which answered. vendor always asks the vendor.
func (c *Client) GetBacklinks(ctx context.Context, in GetBacklinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_backlinks", in)
}

// GetBatchAnalysisInput is the input to get_batch_analysis.
type GetBatchAnalysisInput struct {
	ClientID      string   `json:"clientId"`
	Targets       []string `json:"targets"`
	LocationCode  *int64   `json:"locationCode,omitempty"` // Default: the client's project.
	LanguageCode  string   `json:"languageCode,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetBatchAnalysis calls get_batch_analysis: Compare many domains at once. For
// up to 200 domains (or pages) at once: DataForSEO rank, backlinks, referring
// domains, organic traffic and organic keywords, in one table. For stored
// figures without a purchase, batch_domain_metrics. Answers 'size these sites
// up side by side' and 'which of these prospects are worth it'. Four bulk
// reads, whatever the number of targets. A paid lookup on the client's account
// unless anyone read the same thing this week (spend says which); quote it
// with estimate_cost.
func (c *Client) GetBatchAnalysis(ctx context.Context, in GetBatchAnalysisInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_batch_analysis", in)
}

// GetBestPagesByLinksInput is the input to get_best_pages_by_links.
type GetBestPagesByLinksInput struct {
	ClientID      string `json:"clientId"`
	Domain        string `json:"domain,omitempty"` // Any domain; default the client's own.
	Source        string `json:"source,omitempty"`
	Limit         *int64 `json:"limit,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetBestPagesByLinks calls get_best_pages_by_links: Pages that attract links.
// A domain's pages by how many domains link to them. Run it on competitors to
// see what kind of content earns links in a niche — the content worth making.
// Answers 'what do people link to'. source index is free: our own link index
// (our crawls, client crawls' outlinks, Common Crawl, Bing). auto answers from
// the index when it holds 50+ links to the domain (or, for referring domains,
// a complete pass of Common Crawl's domain graph), and otherwise asks the
// vendor — a paid lookup on the client's account unless anyone bought it in
// the last week; coverage says which answered. vendor always asks the vendor.
func (c *Client) GetBestPagesByLinks(ctx context.Context, in GetBestPagesByLinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_best_pages_by_links", in)
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

// GetBingCrawlHealthInput is the input to get_bing_crawl_health.
type GetBingCrawlHealthInput struct {
	ClientID string `json:"clientId"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetBingCrawlHealth calls get_bing_crawl_health: How Bing crawls the site.
// Bing's crawl stats by day (pages crawled, in index, errors by HTTP class,
// blocked by robots.txt, DNS failures and timeouts), the URLs Bing reports
// crawl issues for with the issues counted by type, and the pages or
// directories blocked from Bing's results. Answers 'is Bing having trouble
// crawling us' and 'did someone hide a page from Bing'. Bing's API has no
// penalty or manual-action feed, so none is reported. Free.
func (c *Client) GetBingCrawlHealth(ctx context.Context, in GetBingCrawlHealthInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_bing_crawl_health", in)
}

// GetBingInboundLinksInput is the input to get_bing_inbound_links.
type GetBingInboundLinksInput struct {
	ClientID string `json:"clientId"`
	URL      string `json:"url,omitempty"`
	Page     *int64 `json:"page,omitempty"`
}

// GetBingInboundLinks calls get_bing_inbound_links: Links Bing knows point at
// the site. Bing's own link data from Webmaster Tools: without url, the
// client's pages that have inbound links and how many each has; with url (a
// page on the client's site), the pages linking to it and their anchor text.
// Paged as Bing pages it (page from 0; totalPages says how many). Answers
// 'which of our pages earn links, according to Bing' and 'who links to this
// page'. Free.
func (c *Client) GetBingInboundLinks(ctx context.Context, in GetBingInboundLinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_bing_inbound_links", in)
}

// GetBingKeywordResearchInput is the input to get_bing_keyword_research.
type GetBingKeywordResearchInput struct {
	ClientID string `json:"clientId"`
	Keyword  string `json:"keyword"`
	Country  string `json:"country,omitempty"`
	Language string `json:"language,omitempty"`
	From     string `json:"from,omitempty"` // Default: 90 days before to.
	To       string `json:"to,omitempty"`   // Default: today.
}

// GetBingKeywordResearch calls get_bing_keyword_research: Bing's own search
// volume for a term. Bing Webmaster Tools' keyword research on the agency's
// key: the term's Bing impressions week by week (exact and broad match) and
// related terms with their impressions over a date range. Bing publishes real
// counts rather than bands. Answers 'how often is this searched on Bing' and
// 'what else do Bing users search around it'. country is a two-letter code,
// language a tag such as en-GB. Free.
func (c *Client) GetBingKeywordResearch(ctx context.Context, in GetBingKeywordResearchInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_bing_keyword_research", in)
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

// GetBingSearchStatsInput is the input to get_bing_search_stats.
type GetBingSearchStatsInput struct {
	ClientID  string `json:"clientId"`
	Dimension string `json:"dimension,omitempty"`
	Query     string `json:"query,omitempty"`
	Page      string `json:"page,omitempty"`
	Limit     *int64 `json:"limit,omitempty"`
}

// GetBingSearchStats calls get_bing_search_stats: Bing queries or pages, week
// by week. Bing Webmaster Tools' query stats or page stats for the client's
// site, each row with its clicks, impressions and position and its weekly
// history (Bing publishes weekly, about six months back). Give query to see
// which pages answered that one query, or page to see which queries that one
// page answered. Answers 'what do people find us for on Bing' and 'which page
// does Bing show for this'. Bing's numbers about Bing's searchers, on the
// agency's Webmaster Tools key; never add them to Google's. Free.
func (c *Client) GetBingSearchStats(ctx context.Context, in GetBingSearchStatsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_bing_search_stats", in)
}

// GetBrandMentionsInput is the input to get_brand_mentions.
type GetBrandMentionsInput struct {
	ClientID      string `json:"clientId"` // The client whose account and monthly cap pay for the lookup
	Keyword       string `json:"keyword"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// GetBrandSplitInput is the input to get_brand_split.
type GetBrandSplitInput struct {
	ClientID   string `json:"clientId"`
	SearchType string `json:"searchType,omitempty"` // Search Console's search type: web (default), image, video, news, discover or googleNews. Discover and Google News report pages but no queries.
	From       string `json:"from,omitempty"`       // Default: 16 months before to.
	To         string `json:"to,omitempty"`         // Default: the newest day archived.
	Interval   string `json:"interval,omitempty"`
	Country    string `json:"country,omitempty"`
}

// GetBrandSplit calls get_brand_split: Branded vs non-branded search. Google
// search clicks and impressions split into branded (the query names the
// client) and non-branded, by day, week or month, with the property total and
// the anonymised remainder Google withholds queries for. Answers 'is
// non-branded traffic growing' — the number content and SEO work actually
// move. Brand terms come from the client brief (update_client_brief
// brandTerms), or the client's name and domain. searchType picks web
// (default), image, video or news; Discover and Google News have no queries to
// split. Free.
func (c *Client) GetBrandSplit(ctx context.Context, in GetBrandSplitInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_brand_split", in)
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

// GetBrokenBacklinksInput is the input to get_broken_backlinks.
type GetBrokenBacklinksInput struct {
	ClientID      string `json:"clientId"`
	Source        string `json:"source,omitempty"`
	Limit         *int64 `json:"limit,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetBrokenBacklinks calls get_broken_backlinks: Links to broken pages. Links
// from other sites to the client's pages that no longer work (4xx, or not
// answering), grouped by page, with the closest live page to redirect each to
// — links already earned, reclaimed with one redirect. Answers 'which links
// are we wasting'. Status comes from the latest crawl, and pages it did not
// crawl are checked. source index is free: our own link index (our crawls,
// client crawls' outlinks, Common Crawl, Bing). auto answers from the index
// when it holds 50+ links to the domain (or, for referring domains, a complete
// pass of Common Crawl's domain graph), and otherwise asks the vendor — a paid
// lookup on the client's account unless anyone bought it in the last week;
// coverage says which answered. vendor always asks the vendor.
func (c *Client) GetBrokenBacklinks(ctx context.Context, in GetBrokenBacklinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_broken_backlinks", in)
}

// GetBudgetStatusInput is the input to get_budget_status.
type GetBudgetStatusInput struct {
	ClientID string `json:"clientId"`
}

// GetBudgetStatus calls get_budget_status: A client's budget and month-end
// forecast. A client's monthly budget (null when none is set), what it has
// spent this month, what is left, when the month resets and the alert share,
// plus projection: an estimate of month-end spend from the paid work already
// scheduled — each rank tracker's cadence × keywords × devices × price, AI
// visibility checks on their cadence, and alert rules with their paid source
// on — priced from the same ceilings estimate_cost quotes, so the real figure
// is usually lower. Answers 'how much budget is left' and 'will this client go
// over this month' before any paid call. set_budget changes the budget;
// get_client_costs says what the spend went on. Free.
func (c *Client) GetBudgetStatus(ctx context.Context, in GetBudgetStatusInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_budget_status", in)
}

// GetBusinessPostsInput is the input to get_business_posts.
type GetBusinessPostsInput struct {
	ClientID      string `json:"clientId"`
	Business      string `json:"business,omitempty"`
	TaskID        string `json:"taskId,omitempty"`
	Depth         *int64 `json:"depth,omitempty"`        // How many posts to read; each ten more costs a little more.
	LocationCode  *int64 `json:"locationCode,omitempty"` // Default: the client's project country.
	LanguageCode  string `json:"languageCode,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetBusinessPosts calls get_business_posts: Any business's Google posts. The
// posts (updates, offers, events) any business has published on its Google
// Business Profile, newest first: text, date, image, links and the post's
// address, with how many there are, when the latest went up and how many in
// the last 90 days. Collecting takes a few minutes: the first call (with
// business) queues it and returns a taskId — that call is the paid one — and
// calling again with the taskId collects it for free (status pending until
// ready). Answers 'how often does this competitor post on Google' and 'what
// are they promoting'. A paid lookup on the client's account unless anyone
// read the same thing recently; quote it with estimate_cost.
func (c *Client) GetBusinessPosts(ctx context.Context, in GetBusinessPostsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_business_posts", in)
}

// GetBusinessProfileInput is the input to get_business_profile.
type GetBusinessProfileInput struct {
	ClientID      string   `json:"clientId"`
	Business      string   `json:"business"`               // The business name (best with the town), or cid:123… / place_id:ChIJ… for an exact listing.
	LocationCode  *int64   `json:"locationCode,omitempty"` // Default: the client's project country.
	LanguageCode  string   `json:"languageCode,omitempty"`
	Lat           *float64 `json:"lat,omitempty"`
	Lng           *float64 `json:"lng,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetBusinessProfile calls get_business_profile: Any business's Google
// profile. The Google Business Profile of any business, found by name or by
// cid / place_id: categories, address, phone, site, coordinates, rating and
// its spread, review count, hours, attributes, photos, whether it is claimed,
// the topics reviewers mention and the businesses people also search. Gives
// the place_id and coordinates add_local_location needs. Answers 'what does
// this competitor's listing have that ours lacks'. A paid lookup on the
// client's account unless anyone read the same thing recently; quote it with
// estimate_cost.
func (c *Client) GetBusinessProfile(ctx context.Context, in GetBusinessProfileInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_business_profile", in)
}

// GetBusinessQuestionsInput is the input to get_business_questions.
type GetBusinessQuestionsInput struct {
	ClientID      string   `json:"clientId"`
	Business      string   `json:"business"` // The business name (best with the town), or cid:123… / place_id:ChIJ… for an exact listing.
	Limit         *int64   `json:"limit,omitempty"`
	LocationCode  *int64   `json:"locationCode,omitempty"` // Default: the client's project country.
	LanguageCode  string   `json:"languageCode,omitempty"`
	Lat           *float64 `json:"lat,omitempty"`
	Lng           *float64 `json:"lng,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetBusinessQuestions calls get_business_questions: A business's Google Q&A.
// The questions people asked on any business's Google listing, with the
// answers and who gave them, and how many are unanswered. Answers 'what do
// customers ask about us' and 'which questions should we answer on our
// profile'. A paid lookup on the client's account unless anyone read the same
// thing recently; quote it with estimate_cost.
func (c *Client) GetBusinessQuestions(ctx context.Context, in GetBusinessQuestionsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_business_questions", in)
}

// GetBusinessReviewsInput is the input to get_business_reviews.
type GetBusinessReviewsInput struct {
	ClientID      string `json:"clientId"`
	Business      string `json:"business,omitempty"`
	TaskID        string `json:"taskId,omitempty"`
	Depth         *int64 `json:"depth,omitempty"`
	LocationCode  *int64 `json:"locationCode,omitempty"` // Default: the client's project country.
	LanguageCode  string `json:"languageCode,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetBusinessReviews calls get_business_reviews: Any business's Google
// reviews. Google reviews of any business (up to 500, newest first) with the
// rating, text, date, author and whether the owner replied. Collecting takes a
// minute or two: the first call (with business) queues it and returns a taskId
// — that call is the paid one — and calling again with the taskId collects it
// for free (status pending until ready). For the client's own locations,
// sync_local_reviews keeps them stored. Answers 'what do customers complain
// about at this competitor'. A paid lookup on the client's account unless
// anyone read the same thing recently; quote it with estimate_cost.
func (c *Client) GetBusinessReviews(ctx context.Context, in GetBusinessReviewsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_business_reviews", in)
}

// GetCannibalizationInput is the input to get_cannibalization.
type GetCannibalizationInput struct {
	ClientID string `json:"clientId"`
}

// GetCannibalization calls get_cannibalization: Get cannibalization. Keywords
// where several of the client's own pages compete in the same result, each
// with the page the keyword registry says owns it (owner, null when
// unregistered) and the page that should win (shouldOwn: the registered owner,
// else the best-placed page). Answers 'where are we competing with ourselves'
// and 'which page should own this query'. set_keyword_owners registers owners.
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
// Answers 'what are we paying for on this account'. The budget, what is left
// of it and a month-end forecast are get_budget_status. Amounts are US
// dollars.
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
// and whether IndexNow is set up. backfill gives Search Console's
// sixteen-month history: the range asked for (from, to), how far it has got
// (collectedThrough, progress), its status — pending, running, retrying or
// complete — the rows stored and the days they span (rowsStored, storedFrom,
// storedThrough), and a note saying where Google's data for the property
// begins when it holds less than sixteen months. Answers 'is Search Console
// connected for this client' and 'is the history in yet'. Never returns
// tokens.
func (c *Client) GetClientIntegrations(ctx context.Context, in GetClientIntegrationsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_client_integrations", in)
}

// GetCompetitorBrokenPagesInput is the input to get_competitor_broken_pages.
type GetCompetitorBrokenPagesInput struct {
	ClientID      string `json:"clientId"`
	Domain        string `json:"domain"` // The competitor's domain.
	Source        string `json:"source,omitempty"`
	DofollowOnly  *bool  `json:"dofollowOnly,omitempty"` // Only pages with at least one dofollow referring domain.
	Limit         *int64 `json:"limit,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetCompetitorBrokenPages calls get_competitor_broken_pages: A competitor's
// dead pages that still have links. A domain's pages that answer 404 or 410
// but that other sites still link to, with how many referring domains (and how
// many of them dofollow) each still has — the broken-link-building targets:
// offer the linking sites your page instead. Run on the client's own domain it
// lists links to reclaim with a redirect (get_broken_backlinks adds the
// redirect target). source vendor reads DataForSEO's index; index reads our
// own crawl, free but only pages we have fetched; auto asks the vendor and
// falls back to the index. Answers 'which of their dead pages could we take
// the links of'. source vendor (and auto, when a provider is connected) is a
// paid lookup on the client's account unless anyone read it this week; quote
// it with estimate_cost. index is free.
func (c *Client) GetCompetitorBrokenPages(ctx context.Context, in GetCompetitorBrokenPagesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_competitor_broken_pages", in)
}

// GetCompetitorNewPagesInput is the input to get_competitor_new_pages.
type GetCompetitorNewPagesInput struct {
	ClientID    string   `json:"clientId"`
	Days        *int64   `json:"days,omitempty"`
	Competitors []string `json:"competitors,omitempty"`
	Limit       *int64   `json:"limit,omitempty"`
}

// GetCompetitorNewPages calls get_competitor_new_pages: What competitors
// published. New and updated pages on competitors' sites, from their sitemaps
// (read daily for every competitor a tracker watches, and read now for any not
// yet watched), grouped by section — 'competitor X published 14 pages under
// /guides this week'. Answers 'what are they writing about' and 'what did they
// just launch'. competitors overrides the tracked ones. Free.
func (c *Client) GetCompetitorNewPages(ctx context.Context, in GetCompetitorNewPagesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_competitor_new_pages", in)
}

// GetCompetitorPageChangesInput is the input to get_competitor_page_changes.
type GetCompetitorPageChangesInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId,omitempty"` // Default: every tracker of the client.
	Days      *int64 `json:"days,omitempty"`
}

// GetCompetitorPageChanges calls get_competitor_page_changes: Changes to pages
// that outrank us. Pages that rank above the client for its tracked keywords
// whose title, H1, length or sections changed — read from snapshots taken
// every few days. Answers 'what did the page that beat us change'. Free.
func (c *Client) GetCompetitorPageChanges(ctx context.Context, in GetCompetitorPageChangesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_competitor_page_changes", in)
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

// GetContentAuditInput is the input to get_content_audit.
type GetContentAuditInput struct {
	ClientID string `json:"clientId"`
	Verdict  string `json:"verdict,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetContentAudit calls get_content_audit: Content audit: keep, update, merge
// or review each page. A verdict for every page meant to rank in the latest
// crawl, from its Google search over the last 90 days against the 90 before
// (page-level Search Console, top query included), its length and whether it
// nearly duplicates another page: keep (earning and steady), update (losing
// clicks, seen but rarely clicked, or thin for the demand it has), merge
// (near-duplicate of another page) review (no search demand in six months —
// improve it, fold it into another, or noindex it, after checking its links
// and conversions; never deleted on this alone) or no_search_data (no Search
// Console rows for it, and too little history to call that absent demand).
// Each with the reason; excluded counts the crawled pages left out (non-200,
// noindex, canonical elsewhere, parameter copies). Without Search Console only
// merge is judged. Answers 'which pages should we refresh, merge or prune'.
// Free.
func (c *Client) GetContentAudit(ctx context.Context, in GetContentAuditInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_content_audit", in)
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

// GetContentGroupsInput is the input to get_content_groups.
type GetContentGroupsInput struct {
	ClientID   string           `json:"clientId"`
	Days       *int64           `json:"days,omitempty"`       // Length of each period, ending on the newest day Google has finished counting.
	Compare    string           `json:"compare,omitempty"`    // previous: the same number of days just before. year: the same days a year earlier.
	SearchType string           `json:"searchType,omitempty"` // Search Console's search type: web (default), image, video, news, discover or googleNews. Discover and Google News report pages but no queries.
	Groups     []map[string]any `json:"groups,omitempty"`
	Country    string           `json:"country,omitempty"`
	Device     string           `json:"device,omitempty"`
}

// GetContentGroups calls get_content_groups: Search traffic by section. Google
// search clicks, impressions and position rolled up by section of the site,
// two periods compared. Answers 'how is the blog doing against the guides' and
// 'which template is losing traffic'. groups are matched in order against each
// page's path, first match wins (* matches anything, so '/blog/*'); without
// groups, pages are grouped by their first path segment. searchType reads
// image, video, news, Discover or Google News instead of web. Free.
func (c *Client) GetContentGroups(ctx context.Context, in GetContentGroupsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_content_groups", in)
}

// GetConversionsInput is the input to get_conversions.
type GetConversionsInput struct {
	ClientID string   `json:"clientId"`
	Source   string   `json:"source,omitempty"`
	By       string   `json:"by,omitempty"`
	Channel  string   `json:"channel,omitempty"`
	Goals    []string `json:"goals,omitempty"`
	Days     *int64   `json:"days,omitempty"`
	Limit    *int64   `json:"limit,omitempty"`
}

// GetConversions calls get_conversions: Conversions by landing page and
// source. Conversions (sign-ups, purchases, any goal) by the page a visit
// landed on, its source or its referrer, with visits, converting visits,
// conversion rate and value — and, by landing page, the Search Console queries
// that bring those visits. Answers 'which pages and queries actually make
// money' and 'what is organic search worth'. source web_analytics uses the
// cookieless analytics' custom events (goals: the event names; default every
// event; value from a 'value' property); ga4 uses GA4's key events on organic
// sessions. channel narrows to search engines or AI assistants. Free.
func (c *Client) GetConversions(ctx context.Context, in GetConversionsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_conversions", in)
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

// GetCrawlBudgetInput is the input to get_crawl_budget.
type GetCrawlBudgetInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
	Bot      string `json:"bot,omitempty"`
}

// GetCrawlBudget calls get_crawl_budget: Where crawlers spend their visits.
// From the client's connected logs over the last days (default 30): crawler
// requests by URL template (the site section, not the page), each with its
// share of all crawler requests, how much went on errors and redirects, and
// which crawlers; plus per crawler how many requests were verified (reverse
// DNS or published IP ranges), spoofed (claimed the name, failed the check —
// left out of every figure) or not checkable. Assets are left out. bot narrows
// to one crawler (googlebot, bingbot, gptbot…). Answers 'what is Googlebot
// wasting its time on' and 'how much of our crawl budget goes on parameter
// URLs'. Free.
func (c *Client) GetCrawlBudget(ctx context.Context, in GetCrawlBudgetInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_crawl_budget", in)
}

// GetCrawlSettingsInput is the input to get_crawl_settings.
type GetCrawlSettingsInput struct {
	ClientID string `json:"clientId"`
}

// GetCrawlSettings calls get_crawl_settings: How the site is crawled. The
// client's crawl settings: mode (spider, list or sitemap), include and exclude
// URL patterns, page and depth limits, the user agent sent and the robots.txt
// group obeyed, speed (requests per second and concurrency), query parameters
// stripped from discovered URLs, custom extraction rules, JavaScript rendering
// (off, shells — only pages that look unrendered — or all, up to
// renderMaxPages a crawl), whether staging credentials are set (names only,
// never values), whether requests are signed with Web Bot Auth, and the crawl
// schedule (cadence, day, UTC hour and nextRunAt; weekly on Monday 02:00 UTC
// unless set_crawl_schedule changed it). Answers 'what does the crawl do
// differently for this site'. Free.
func (c *Client) GetCrawlSettings(ctx context.Context, in GetCrawlSettingsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_crawl_settings", in)
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

// GetCtrOutliersInput is the input to get_ctr_outliers.
type GetCtrOutliersInput struct {
	ClientID       string   `json:"clientId"`
	Days           *int64   `json:"days,omitempty"` // Days of Search Console history, ending on the newest day Google has finished.
	MinImpressions *int64   `json:"minImpressions,omitempty"`
	Sigma          *float64 `json:"sigma,omitempty"`
	Limit          *int64   `json:"limit,omitempty"`
}

// GetCtrOutliers calls get_ctr_outliers: Queries clicking far from what they
// should. Fits the site's own click-through rate by position bucket and device
// — and by results-page shape (AI Overview, featured snippet, map pack,
// shopping, video, plain) where rank captures say what the page showed — then
// lists Search Console queries more than sigma (default 2) standard deviations
// under it (a title or snippet problem: under, with the clicks missed) and
// over it (a snippet worth copying: over). Falls back to the published
// industry curve where the site has too few queries in a bucket, and model
// says which was used (site, published or mixed). Answers 'which snippets
// should we rewrite' and 'which titles work best'. Free.
func (c *Client) GetCtrOutliers(ctx context.Context, in GetCtrOutliersInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ctr_outliers", in)
}

// GetDemandMatrixInput is the input to get_demand_matrix.
type GetDemandMatrixInput struct {
	ClientID      string         `json:"clientId"`
	Pattern       string         `json:"pattern"` // Words with {slot} placeholders, e.g. 'receive sms {service} {country}'.
	Slots         map[string]any `json:"slots"`
	LocationCode  *int64         `json:"locationCode,omitempty"`
	LanguageCode  string         `json:"languageCode,omitempty"`
	MinVolume     *int64         `json:"minVolume,omitempty"`     // Below this monthly volume a combination is skip.
	MaxCostMicros *int64         `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool          `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetDemandMatrix calls get_demand_matrix: Demand for every combination.
// Search volume, difficulty and CPC for every combination of a pattern's slots
// — 'receive sms {service} {country}' with services and countries — up to 500,
// with the client's ranking page and position for each from Search Console and
// a verdict: write (demand, no page), improve (a page beyond the top ten),
// have (top ten) or skip (no demand). Answers 'which of these template pages
// deserve to exist'. Volumes any agency bought before are reused free; the
// rest is a paid keyword lookup on the client's account, in batches.
func (c *Client) GetDemandMatrix(ctx context.Context, in GetDemandMatrixInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_demand_matrix", in)
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

// GetDisavowHistoryInput is the input to get_disavow_history.
type GetDisavowHistoryInput struct {
	ClientID string `json:"clientId"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetDisavowHistory calls get_disavow_history: The client's saved disavow
// files. The disavow files saved for the client, newest first, with how many
// domains and URLs each names, and what the newest changed against the one
// before. Answers 'what did we disavow last time' and 'what changed in the
// disavow file'. get_toxic_link_risk lists risky sites that appeared since.
// Free.
func (c *Client) GetDisavowHistory(ctx context.Context, in GetDisavowHistoryInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_disavow_history", in)
}

// GetDomainFactsInput is the input to get_domain_facts.
type GetDomainFactsInput struct {
	Domain        string `json:"domain"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetDomainFacts calls get_domain_facts: What the open web says about a
// domain. Registration date and registrar from the registry, position on the
// Tranco top-sites list, when the Internet Archive first captured it, and the
// subdomains its certificates have announced, plus topicShift: the homepage's
// title and description as first archived against today's, crossed with the
// registration date, as a risk level (high: a different topic on a domain
// registered again after the archive first saw it, registered recently, or
// registered five or more years ago — the expired-domain abuse pattern;
// medium: a changed topic on a registration two to five years old; unknown
// when either page could not be read; a heuristic). Answers 'how old and how
// big is this, and what else is on it' and 'was this domain bought and
// repurposed' for any domain, free and without touching a vendor. Each figure
// is absent far more often than present, and an absence is never a zero: no
// registry record, not in the top million, never archived.
func (c *Client) GetDomainFacts(ctx context.Context, in GetDomainFactsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_domain_facts", in)
}

// GetDomainKeywordGapInput is the input to get_domain_keyword_gap.
type GetDomainKeywordGapInput struct {
	ClientID      string   `json:"clientId"`
	Target        string   `json:"target,omitempty"` // Default: the client's own domain.
	Competitors   []string `json:"competitors"`
	Mode          string   `json:"mode,omitempty"`
	MinVolume     *int64   `json:"minVolume,omitempty"`
	Limit         *int64   `json:"limit,omitempty"`
	LocationCode  *int64   `json:"locationCode,omitempty"` // Default: the client's project.
	LanguageCode  string   `json:"languageCode,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetDomainKeywordGap calls get_domain_keyword_gap: Keyword gap against
// competitors. The keywords up to five competitors rank for in Google and the
// client does not (mode missing), that both rank for (shared), or where the
// client ranks below them (weak) — merged across the competitors, with volume,
// difficulty, CPC, intent, the client's position and each competitor's. Works
// for any domains, tracked or not, from DataForSEO's index; get_keyword_gap is
// the free reading from what the client's trackers already captured. Answers
// 'what do our competitors rank for that we don't' and 'where are we losing to
// them'. One read per competitor. A paid lookup on the client's account unless
// anyone read the same thing this week (spend says which); quote it with
// estimate_cost.
func (c *Client) GetDomainKeywordGap(ctx context.Context, in GetDomainKeywordGapInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_domain_keyword_gap", in)
}

// GetDomainOverviewInput is the input to get_domain_overview.
type GetDomainOverviewInput struct {
	ClientID      string  `json:"clientId"`
	Target        string  `json:"target,omitempty"` // A domain (competitor.com), a subfolder (competitor.com/blog/) or a page URL. Default: the client's own domain.
	LocationCodes []int64 `json:"locationCodes,omitempty"`
	LanguageCode  string  `json:"languageCode,omitempty"`
	HistoryMonths *int64  `json:"historyMonths,omitempty"`
	MaxCostMicros *int64  `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool   `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetDomainOverview calls get_domain_overview: A domain's organic and paid
// standing. For any domain: organic and paid keywords, estimated traffic and
// its value, the split of positions (top 3, 4–10, 11–100) and keywords new,
// up, down and lost — in one or more countries (locationCodes, up to 5) — and
// a monthly history (historyMonths, up to 36). Answers 'how big is this
// competitor', 'where do they get traffic' and 'is their traffic growing'. A
// paid lookup on the client's account unless anyone read the same thing
// recently (spend says which); quote it with estimate_cost.
func (c *Client) GetDomainOverview(ctx context.Context, in GetDomainOverviewInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_domain_overview", in)
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

// GetEngineCapabilitiesInput is the input to get_engine_capabilities.
type GetEngineCapabilitiesInput struct {
	Engines      []string `json:"engines,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
}

// GetEngineCapabilities calls get_engine_capabilities: What can be done on
// each search engine. The matrix of what this product does on each engine —
// Google, Bing, YouTube, Amazon, Yandex, Baidu, Naver, DuckDuckGo, Yahoo,
// Seznam, App Store, Google Play — for rank tracking, live results, search
// volume, difficulty, keyword ideas, ranked keywords, top pages, content
// profiles, SERP features and history, webmaster data and URL submission. Each
// cell is supported (with the tools that do it), not_offered_by_engine (the
// engine publishes nothing to read) or not_built, with a note. Also lists the
// AI answer engines and whether their data is sampled by Seolyticize or from
// DataForSEO. Answers 'can you do X on Bing' and 'why is there no Naver index
// data'. Free.
func (c *Client) GetEngineCapabilities(ctx context.Context, in GetEngineCapabilitiesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_engine_capabilities", in)
}

// GetEngineKeywordIdeasInput is the input to get_engine_keyword_ideas.
type GetEngineKeywordIdeasInput struct {
	ClientID      string   `json:"clientId"`
	Engine        string   `json:"engine"`
	Seeds         []string `json:"seeds"`
	LocationCode  *int64   `json:"locationCode,omitempty"`
	LanguageCode  string   `json:"languageCode,omitempty"`
	Limit         *int64   `json:"limit,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetEngineKeywordIdeas calls get_engine_keyword_ideas: Keyword ideas for Bing
// or Amazon. Keyword ideas around seed terms on another engine: engine bing
// reads DataForSEO's Bing keywords-for-keywords (with Bing volume, CPC,
// competition and monthly history; English, French and German), engine amazon
// reads DataForSEO Labs' Amazon related keywords (with Amazon search volume;
// one seed). Answers 'what do people search on Amazon around this product' and
// 'what are the Bing variations of this term'. Google ideas are
// keyword_explorer. A paid lookup on the client's account unless anyone read
// the same thing recently (spend says which); quote it with estimate_cost.
func (c *Client) GetEngineKeywordIdeas(ctx context.Context, in GetEngineKeywordIdeasInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_engine_keyword_ideas", in)
}

// GetEntityCoverageInput is the input to get_entity_coverage.
type GetEntityCoverageInput struct {
	ClientID      string `json:"clientId"`
	Keyword       string `json:"keyword"`
	URL           string `json:"url,omitempty"`
	Html          string `json:"html,omitempty"`
	Markdown      string `json:"markdown,omitempty"`
	Resolve       *bool  `json:"resolve,omitempty"`
	LocationCode  *int64 `json:"locationCode,omitempty"`
	LanguageCode  string `json:"languageCode,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetEntityCoverage calls get_entity_coverage: The entities a page should
// cover. The named entities (people, products, places, organisations,
// concepts) at least two of the ranking pages for a keyword cover, from their
// capitalised phrases, each resolved on Wikidata (resolve: true, free),
// whether the client's page (url, html or markdown; or none, to plan one)
// names it, a 0–100 coverage score, and the JSON-LD to add: about (the main
// entity) and mentions, with sameAs links to Wikidata and Wikipedia. Answers
// 'which entities does our page miss' and 'what sameAs and about schema should
// it carry'. Same cost as get_serp_content_profile.
func (c *Client) GetEntityCoverage(ctx context.Context, in GetEntityCoverageInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_entity_coverage", in)
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

// GetErrorCatalogueInput is the input to get_error_catalogue.
type GetErrorCatalogueInput struct {
	Code string `json:"code,omitempty"`
}

// GetErrorCatalogue calls get_error_catalogue: What every error code means.
// Every refusal code the API returns, with what it means and what to do about
// it — the same list as GET /api/v1/errors. Answers 'what does
// spend_cap_exceeded mean'. Free.
func (c *Client) GetErrorCatalogue(ctx context.Context, in GetErrorCatalogueInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_error_catalogue", in)
}

// GetExtractionResultsInput is the input to get_extraction_results.
type GetExtractionResultsInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	Name     string `json:"name,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	Offset   *int64 `json:"offset,omitempty"`
}

// GetExtractionResults calls get_extraction_results: Custom extraction
// results. The values custom extraction (set_crawl_settings extractions) read
// off each page in a crawl, page by page; name narrows to one rule. Answers
// 'which pages have no author', 'what price does each product page show' and
// 'which posts are older than X'. Free.
func (c *Client) GetExtractionResults(ctx context.Context, in GetExtractionResultsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_extraction_results", in)
}

// GetForumOpportunitiesInput is the input to get_forum_opportunities.
type GetForumOpportunitiesInput struct {
	ClientID      string   `json:"clientId"`
	Keywords      []string `json:"keywords,omitempty"`
	Days          *int64   `json:"days,omitempty"` // For tracked keywords: how old a capture may be.
	MaxPosition   *int64   `json:"maxPosition,omitempty"`
	LocationCode  *int64   `json:"locationCode,omitempty"`
	LanguageCode  string   `json:"languageCode,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetForumOpportunities calls get_forum_opportunities: Reddit, Quora and forum
// threads ranking for our keywords. Threads on Reddit, Quora and forums that
// rank in Google for the client's keywords, with the thread title, address and
// position, and whether Google shows it in its Discussions and forums box —
// where a helpful answer earns visibility and referral traffic the client's
// own page cannot. Answers 'which Reddit threads rank for our keywords' and
// 'where should we join the conversation'. With no keywords given, reads the
// tracked keywords' latest captured results pages, free. With keywords (up to
// 20), reads Google's results for each now: A paid lookup on the client's
// account unless anyone read the same thing recently (spend says which); quote
// it with estimate_cost.
func (c *Client) GetForumOpportunities(ctx context.Context, in GetForumOpportunitiesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_forum_opportunities", in)
}

// GetFreshnessQueueInput is the input to get_freshness_queue.
type GetFreshnessQueueInput struct {
	ClientID string `json:"clientId"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetFreshnessQueue calls get_freshness_queue: Pages to refresh first. The
// client's pages ranked by Search Console clicks at risk (a month's clicks,
// weighted by how many stale signals the page shows and doubled when it is
// already losing them) from going stale: the results page for its top query
// moved to newer content (the ranking pages' stored dates, at the median, over
// six months newer than the page's own), a past year in the title, past years
// in the text (with the sentence), a price or limit contradicting the client's
// registered facts, or no date on a page whose rivals are recent. Each with
// its reasons and priority. Answers 'what should we update this month'. Free:
// stored results pages and the latest crawl only.
func (c *Client) GetFreshnessQueue(ctx context.Context, in GetFreshnessQueueInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_freshness_queue", in)
}

// GetFunnelReportInput is the input to get_funnel_report.
type GetFunnelReportInput struct {
	ClientID string `json:"clientId"`
	FunnelID string `json:"funnelId"`
	By       string `json:"by,omitempty"` // channel (AI assistants, organic search, paid, social, email, referral, direct), landing_page, source, campaign, ai_referrer (which assistant sent the visit), qu
	Days     *int64 `json:"days,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetFunnelReport calls get_funnel_report: Funnel steps and drop-off. How many
// visits reached each step of a funnel in order, with the share of the first
// step and of the step before, overall and per channel, landing page, source,
// campaign, AI referrer or query cluster. Answers 'where do people drop out
// between pricing and sign-up' and 'does search traffic finish the funnel'.
// Free.
func (c *Client) GetFunnelReport(ctx context.Context, in GetFunnelReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_funnel_report", in)
}

// GetGA4AudienceInput is the input to get_ga4_audience.
type GetGA4AudienceInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Compare  *bool  `json:"compare,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	By       string `json:"by,omitempty"`
	Channel  string `json:"channel,omitempty"` // organic: Organic Search; ai: sessions sent by AI assistants (ChatGPT, Perplexity, Gemini, Copilot, Claude and others, read from the session source); all: every
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
	Channel  string `json:"channel,omitempty"` // organic: Organic Search; ai: sessions sent by AI assistants (ChatGPT, Perplexity, Gemini, Copilot, Claude and others, read from the session source); all: every
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
// dimensions and metrics, the last day with sessions, and Consent Mode v2 on
// each stream's site (an inline gtag consent default naming ad_user_data and
// ad_personalization; a consent platform or Tag Manager that may set it out of
// sight is named, not marked missing). findings lists problems, warnings and
// notes in plain words; a part that could not be read is reported as unread,
// not as fine. Answers 'can we trust this client's Analytics'. Refuses with a
// sentence when the client has not connected a GA4 property. Free.
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
	Channel  string `json:"channel,omitempty"` // organic: Organic Search; ai: sessions sent by AI assistants (ChatGPT, Perplexity, Gemini, Copilot, Claude and others, read from the session source); all: every
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
	Channel  string `json:"channel,omitempty"` // organic: Organic Search; ai: sessions sent by AI assistants (ChatGPT, Perplexity, Gemini, Copilot, Claude and others, read from the session source); all: every
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
	GroupBy  string `json:"groupBy,omitempty"` // aiSource: the source / medium of AI assistants' sessions only.
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

// GetGeoReadinessInput is the input to get_geo_readiness.
type GetGeoReadinessInput struct {
	ClientID      string   `json:"clientId"`
	URLs          []string `json:"urls,omitempty"`
	TopPages      *int64   `json:"topPages,omitempty"`
	BrandMentions *bool    `json:"brandMentions,omitempty"` // Also count the brand's mentions on the web, Reddit, YouTube, LinkedIn and Wikipedia from Google search: five paid results pages (quote with estimate_cost).
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetGeoReadiness calls get_geo_readiness: How ready the pages are to be cited
// in AI answers. A 0–100 generative-engine readiness score per page, for the
// pages given or the top pages by Search Console clicks (up to 20): citability
// 25 (an answer in the first 40–60 words of each H2 section, self-contained
// 130–170-word passages, 'X is…' definitions, figures with a cited source, the
// most citable block in the first 30% of the page), structure 20 (heading
// hierarchy, question headings, 2–4-sentence paragraphs, tables and lists),
// multimodal 15 (images, video, VideoObject markup, infographics and charts,
// calculators and other interactive elements), authority 20 (author, dates,
// recency — under 3 months best, over 6 stale — links to primary sources, and
// the brand on Wikipedia, Wikidata, Reddit, YouTube and LinkedIn from the
// site's sameAs and Wikidata, expert quotes with attribution), technical 20
// (server-side rendering from the crawl, OAI-SearchBot, Claude-SearchBot and
// PerplexityBot allowed, an RSL licence, snippets allowed — nosnippet or
// max-snippet:0 keeps a page out of AI Overviews and AI Mode, data-nosnippet
// reported; llms.txt reported at weight 0). Training, search and user crawlers
// are reported apart, Bytespider, cohere-ai, Google-CloudVertexBot and
// Google-Agent (which ignores robots.txt) among them. Also a readiness score
// per answer engine (Google AI Overviews, AI Mode, ChatGPT, Perplexity,
// Claude, Bing Copilot) from the parts each leans on, 0 where its crawler is
// blocked: readiness, not measured citations. brandMentions: true adds how
// often Google finds the brand named on the web, Reddit, YouTube, LinkedIn and
// Wikipedia (five paid results pages; quote with estimate_cost). Unmeasured
// parts are null and the rest rescaled; a heuristic, and said so. Answers 'are
// our pages written to be cited by ChatGPT and Perplexity' and 'what is our
// GEO score'. Free without brandMentions.
func (c *Client) GetGeoReadiness(ctx context.Context, in GetGeoReadinessInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_geo_readiness", in)
}

// GetGoalReportInput is the input to get_goal_report.
type GetGoalReportInput struct {
	ClientID string   `json:"clientId"`
	GoalIDs  []string `json:"goalIds,omitempty"`
	By       string   `json:"by,omitempty"` // channel (AI assistants, organic search, paid, social, email, referral, direct), landing_page, source, campaign, ai_referrer (which assistant sent the visit), qu
	Days     *int64   `json:"days,omitempty"`
	Limit    *int64   `json:"limit,omitempty"`
}

// GetGoalReport calls get_goal_report: Goals by channel, page or query.
// Visits, converting visits, conversions, conversion rate and value for each
// goal (default every goal), broken down by channel, landing page, source,
// campaign, the AI assistant that referred the visit, or the Search Console
// query cluster bringing people to the landing page. Answers 'which channel
// brings sign-ups', 'what is organic search worth in conversions', 'do visits
// from ChatGPT convert'. Free.
func (c *Client) GetGoalReport(ctx context.Context, in GetGoalReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_goal_report", in)
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

// GetGoogleTrendsInput is the input to get_google_trends.
type GetGoogleTrendsInput struct {
	ClientID      string   `json:"clientId"`
	Keywords      []string `json:"keywords"`
	TimeRange     string   `json:"timeRange,omitempty"`
	Type          string   `json:"type,omitempty"`
	Worldwide     *bool    `json:"worldwide,omitempty"`    // true ignores the market and reads the whole world.
	LocationCode  *int64   `json:"locationCode,omitempty"` // Default: the client's project.
	LanguageCode  string   `json:"languageCode,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetGoogleTrends calls get_google_trends: Google Trends for up to five
// keywords. Google Trends interest over time (0–100, relative to the busiest
// point) for up to five keywords compared, by region, and — for a single
// keyword — the related searches, top and rising. Web search by default, or
// news, YouTube, images or shopping. Supersedes get_keyword_trends, which
// reads the United States only. Answers 'is this topic growing', 'when does it
// peak' and 'which of these terms is bigger'. A paid lookup on the client's
// account unless anyone read the same thing this week (spend says which);
// quote it with estimate_cost. (about $0.01)
func (c *Client) GetGoogleTrends(ctx context.Context, in GetGoogleTrendsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_google_trends", in)
}

// GetHreflangParityInput is the input to get_hreflang_parity.
type GetHreflangParityInput struct {
	ClientID   string `json:"clientId"`
	CrawlID    string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	URL        string `json:"url,omitempty"`
	SourceLang string `json:"sourceLang,omitempty"`
	MaxPairs   *int64 `json:"maxPairs,omitempty"`
}

// GetHreflangParity calls get_hreflang_parity: Are the language versions real
// translations. Reads each language version of the client's pages (the groups
// the crawl's hreflang or /es/-style paths tie together) and scores every pair
// against its source 0–100: equal H2 and H3 counts; title, meta description
// and schema inLanguage actually localised (the language detected, not just
// declared); length within the band the language should take (German 25–35%
// longer than English, Japanese 10–25% shorter, in words); a stale translation
// (modified before its source); numbers, dates, prices and phone numbers
// written for the locale; and source-language sentences left untranslated.
// Unmeasured parts are null and the rest rescale. Each pair also carries a
// cultural adaptation check against its market's profile (DACH, Francophone,
// Hispanic, Japanese: formal address, hard-sell CTAs, Impressum or Mentions
// légales, local number, date and currency formats, US brands, statistics and
// laws, length against English), each failure a Medium 'Cultural Adaptation'
// finding, and machine-translation QA (body identical to the source, <html
// lang> against the detected language, a translated meta description over 160
// characters, schema inLanguage, Google's translate widget). contentParity
// scores each group 30/30/25/15: the page exists in every declared language,
// title, meta and schema localised, structure (sections ±1, images and charts
// exact, FAQ items ±2), and freshness (≤7, 8–30, 31–90, 90+ days behind). A
// heuristic, and it says so. url picks the group containing that page;
// sourceLang which version the others are measured against (default the
// x-default one). Answers 'is our German site actually translated' and 'which
// translations are out of date'. Free: fetches the client's own pages.
func (c *Client) GetHreflangParity(ctx context.Context, in GetHreflangParityInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_hreflang_parity", in)
}

// GetHreflangReportInput is the input to get_hreflang_report.
type GetHreflangReportInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	Limit    *int64 `json:"limit,omitempty"`
}

// GetHreflangReport calls get_hreflang_report: Hreflang across the site.
// Validates hreflang as clusters of pages rather than page by page, reading it
// from the pages and from xhtml:link entries in the sitemaps (and flagging
// where the two disagree): return links (a pair Google ignores because the
// other page does not name this one back), x-default, valid
// language-script-region codes (zh-Hant-TW is one; en-uk, es-EU, en-UN and
// es-LA are not; a bare country such as gb or us is not a language),
// self-reference, targets that redirect, error, are noindexed or are
// canonicalised elsewhere (canonical conflicts), hreflang on a page that is
// itself canonicalised away, and pages in a group naming different pages for
// one language. Answers 'is our hreflang set up correctly' and 'why does
// Google show the wrong language'. Free.
func (c *Client) GetHreflangReport(ctx context.Context, in GetHreflangReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_hreflang_report", in)
}

// GetImageSerpInput is the input to get_image_serp.
type GetImageSerpInput struct {
	ClientID      string `json:"clientId"`
	Keyword       string `json:"keyword"`
	LocationCode  *int64 `json:"locationCode,omitempty"`
	LanguageCode  string `json:"languageCode,omitempty"`
	Device        string `json:"device,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetImageSerp calls get_image_serp: Who owns the image results for a keyword.
// Reads Google's organic results 100 deep for a keyword and every image pack
// in them: how many packs and where they sit, which domains hold the first ten
// images (and how dominant the top one is), the alt text pattern (how many
// have alt, average length, the share containing the keyword, the commonest
// words), the format mix (JPEG, WebP, AVIF…) and whether the client's site
// appears in a pack and where it ranks organically. Answers 'can we win image
// results for this', 'who dominates image search here' and 'how do the winners
// write their alt text'. Paid: a results page read 100 deep, billed per ten
// results; quote it with estimate_cost.
func (c *Client) GetImageSerp(ctx context.Context, in GetImageSerpInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_image_serp", in)
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

// GetIntegrationConnectLinkInput is the input to get_integration_connect_link.
type GetIntegrationConnectLinkInput struct {
	ClientID string `json:"clientId"`
	Provider string `json:"provider"`
	Shop     string `json:"shop,omitempty"`   // For Shopify: the store's myshopify.com domain.
	TeamID   string `json:"teamId,omitempty"` // For Linear: the team issues are filed in.
}

// GetIntegrationConnectLink calls get_integration_connect_link: Link to
// connect by consent. A link that connects one of the client's accounts by
// signing in rather than pasting a key: Stripe (Connect, read-only), Shopify
// (OAuth; give shop), Slack (installs the app into a channel), Linear (OAuth;
// give teamId), Google Sheets or BigQuery (the client's Google account, with
// only that scope). Valid for ten minutes, for someone signed in to this
// agency. Answers 'connect the client's Google Sheets'. Free.
func (c *Client) GetIntegrationConnectLink(ctx context.Context, in GetIntegrationConnectLinkInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_integration_connect_link", in)
}

// GetInternalLinkGraphInput is the input to get_internal_link_graph.
type GetInternalLinkGraphInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	Sort     string `json:"sort,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	Offset   *int64 `json:"offset,omitempty"`
}

// GetInternalLinkGraph calls get_internal_link_graph: The internal link graph.
// Every crawled page with its internal inlinks and outlinks, clicks from the
// home page, internal PageRank (0–100), and the anchors it is linked with —
// plus orphans (indexable pages nothing links to), pages unreachable from the
// home page, and pages more than three clicks deep. Links that redirect are
// followed to the page they land on; nofollow links do not count. Answers
// 'which pages are buried', 'what holds our internal authority' and 'which
// pages are orphaned'. Free.
func (c *Client) GetInternalLinkGraph(ctx context.Context, in GetInternalLinkGraphInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_internal_link_graph", in)
}

// GetIssueTemplatesInput is the input to get_issue_templates.
type GetIssueTemplatesInput struct {
	ClientID      string `json:"clientId"`
	CrawlID       string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	RuleID        string `json:"ruleId,omitempty"`
	IndexableOnly *bool  `json:"indexableOnly,omitempty"`
}

// GetIssueTemplates calls get_issue_templates: Issues by template. Groups each
// issue's pages by URL template (ids, slugs with numbers and query values
// replaced by *), so '62 pages of /community/es/u/*?sort=*' reads as one fix
// on one template. Answers 'how many fixes is this really' and 'which template
// causes this'. ruleId narrows to one issue. indexableOnly (default true)
// leaves out findings on pages search engines will not index as themselves —
// noindexed, not answering 200, or canonicalised elsewhere — except the
// findings that explain why; pass false to see everything. Free.
func (c *Client) GetIssueTemplates(ctx context.Context, in GetIssueTemplatesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_issue_templates", in)
}

// GetJobInput is the input to get_job.
type GetJobInput struct {
	ClientID string `json:"clientId"`
	JobID    string `json:"jobId"`
}

// GetJob calls get_job: How a slow job is going. One way to follow anything
// slow — a crawl, a rank or domain refresh, a performance run: its status
// (queued, running, done, failed), progress, an ETA where one can be told, and
// the tool that reads the finished result. jobId is kind:id — crawl:<crawlId>
// (from start_crawl), rank:<runId> (refresh_rankings), domain:<runId>
// (refresh_domain_data), performance:<runId>. An AI visibility check reports
// through get_ai_visibility_runs. Answers 'is my crawl done'. Free.
func (c *Client) GetJob(ctx context.Context, in GetJobInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_job", in)
}

// GetKeywordAdvertisersInput is the input to get_keyword_advertisers.
type GetKeywordAdvertisersInput struct {
	ClientID      string `json:"clientId"`
	Keyword       string `json:"keyword"`
	LocationCode  *int64 `json:"locationCode,omitempty"`  // Default: the client's project (2840 United States).
	LanguageCode  string `json:"languageCode,omitempty"`  // Default: the client's project.
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetKeywordAdvertisers calls get_keyword_advertisers: Who advertises on a
// keyword. The ads Google shows for a keyword right now: each advertiser's
// domain, landing page, headline and description, in order. Answers 'who bids
// on this' and 'what do their ads say'. A paid lookup on the client's account
// unless anyone read the same thing recently (spend says which); quote it with
// estimate_cost.
func (c *Client) GetKeywordAdvertisers(ctx context.Context, in GetKeywordAdvertisersInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_keyword_advertisers", in)
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
// rank for that the client does not, sized by the clicks available — from the
// client's own tracked captures, so free. For any domains and every keyword
// they rank for, get_domain_keyword_gap reads DataForSEO's index. Answers
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
// and engine, its keywordId, its note and tags, and the metrics last stored
// for it. Answers 'what is on this list'. To filter, sort or export as CSV use
// query_keyword_list. Free.
func (c *Client) GetKeywordList(ctx context.Context, in GetKeywordListInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_keyword_list", in)
}

// GetKeywordMetricsInput is the input to get_keyword_metrics.
type GetKeywordMetricsInput struct {
	ClientID      string   `json:"clientId"`
	Engine        string   `json:"engine,omitempty"`
	Keywords      []string `json:"keywords"`
	LocationCodes []int64  `json:"locationCodes,omitempty"`
	LocationName  string   `json:"locationName,omitempty"`
	LanguageCode  string   `json:"languageCode,omitempty"`
	Clickstream   *bool    `json:"clickstream,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetKeywordMetrics calls get_keyword_metrics: Metrics for a list of keywords.
// Volume (with 12 months of history), CPC, competition, difficulty, intent and
// SERP features for up to 700 known keywords — in one or more countries
// (locationCodes, up to 5), or a city, county or region (locationName, from
// Google Ads, e.g. 'London,England,United Kingdom'). Countries the research
// index does not cover are answered from Google Ads. clickstream: true adds
// clickstream-refined volume (twice the cost). engine bing gives Bing's
// volume, CPC, competition and history from Bing Ads (English, French and
// German); engine amazon gives Amazon search volume from DataForSEO Labs
// (volume only). YouTube publishes no volume and is refused with that reason.
// Answers 'what is the volume of these keywords', 'how many people search this
// in Madrid' and 'how often is this searched on Bing or Amazon'. A paid lookup
// on the client's account unless anyone read the same thing recently (spend
// says which); quote it with estimate_cost.
func (c *Client) GetKeywordMetrics(ctx context.Context, in GetKeywordMetricsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_keyword_metrics", in)
}

// GetKeywordTrendsInput is the input to get_keyword_trends.
type GetKeywordTrendsInput struct {
	ClientID      string   `json:"clientId"` // The client whose account and monthly cap pay for the lookup
	Terms         []string `json:"terms"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetKeywordTrends calls get_keyword_trends: Get Google Trends interest for
// terms. Interest over time from Google Trends for up to five terms in the
// United States, each on a 0–100 scale relative to its own peak — so it shows
// when people care, not how many. Kept for existing callers; get_google_trends
// does the same for any market or worldwide, with regions, rising searches and
// news, YouTube, image or shopping search, and is the one to use. Answers 'is
// this seasonal' and 'is this growing'. Different from get_search_demand,
// which counts people in our own panel. A paid lookup on the client's account:
// quote it with estimate_cost.
func (c *Client) GetKeywordTrends(ctx context.Context, in GetKeywordTrendsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_keyword_trends", in)
}

// GetLandingPagesInput is the input to get_landing_pages.
type GetLandingPagesInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId"`
}

// GetLandingPages calls get_landing_pages: Which pages the tracked keywords
// land on. For a rank tracker's latest captures: each of the client's pages
// with how many tracked keywords it ranks for, how many in the top 3 and top
// 10, its average position and best keyword — and the keywords where two or
// more of its pages compete. Answers 'which pages carry our rankings' and
// 'where are we cannibalising'. Free.
func (c *Client) GetLandingPages(ctx context.Context, in GetLandingPagesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_landing_pages", in)
}

// GetLinkChecksInput is the input to get_link_checks.
type GetLinkChecksInput struct {
	ClientID string `json:"clientId"`
}

// GetLinkChecks calls get_link_checks: Do the linking pages still link?. What
// our own crawler last saw on each page said to link to the client: live (the
// link is there), lost (the page no longer links), unknown (the page could not
// be read), moved (the page redirected elsewhere and the link is not at the
// new address, with movedTo) or unverifiable_js (the page is a JavaScript app
// shell, so its links cannot be seen without a browser), with counts, and
// rendered: true where the last check read the page in a browser. Only lost is
// a removal. Read from the last check, never fetched now. Answers 'which links
// have we lost'. Run recheck_links to check again, or recheck_links with
// render: true to read the unverifiable_js pages in a browser. Free.
func (c *Client) GetLinkChecks(ctx context.Context, in GetLinkChecksInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_link_checks", in)
}

// GetLinkGapInput is the input to get_link_gap.
type GetLinkGapInput struct {
	ClientID      string   `json:"clientId"`
	Target        string   `json:"target,omitempty"` // Default: the client's own domain.
	Competitors   []string `json:"competitors"`
	Match         string   `json:"match,omitempty"`
	Limit         *int64   `json:"limit,omitempty"`
	Source        string   `json:"source,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetLinkGap calls get_link_gap: Link gap against competitors. The sites that
// link to competitors but not to the client — match any: at least one
// competitor (one read each); match all: every competitor (one read) — with
// each site's rank and how many links it sends each competitor. The outreach
// list Ahrefs calls Link Intersect, for any domains. Answers 'who links to
// them and not to us'. source index answers free from our own link index
// (page-level links and Common Crawl's domain graph; rank is null there); auto
// uses it when it holds every domain involved (50+ links each, or a complete
// graph pass) and otherwise asks DataForSEO; vendor always asks DataForSEO.
// source says which answered. Vendor reads: A paid lookup on the client's
// account unless anyone read the same thing this week (spend says which);
// quote it with estimate_cost.
func (c *Client) GetLinkGap(ctx context.Context, in GetLinkGapInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_link_gap", in)
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

// GetLinkVelocityInput is the input to get_link_velocity.
type GetLinkVelocityInput struct {
	ClientID     string   `json:"clientId"`
	Domains      []string `json:"domains,omitempty"` // Default: the client's site and its tracked competitors.
	Days         *int64   `json:"days,omitempty"`
	BaselineDays *int64   `json:"baselineDays,omitempty"`
	GraceDays    *int64   `json:"graceDays,omitempty"`
}

// GetLinkVelocity calls get_link_velocity: Referring domains gained and lost
// per day. New and lost referring domains per day for the client and its
// tracked competitors (or the domains given), from our own link index and the
// vendor links already synced, each against the domain's own baseline of the
// preceding weeks: spikes (a day of new referring domains far above its usual
// — a campaign, a viral piece, or somebody pointing a link farm at the site)
// and lost clusters (many domains lost on one day — a network deindexed, a
// partner's redesign, a broken migration), with example domains, and five red
// flags with the action each calls for (10× the usual new links in a week,
// half the links lost in a month, no new links for three months, every new
// link from one TLD, a spike from one country). Losses are only called after
// graceDays unseen, and not at all for a domain no index is still reading.
// Answers 'are we gaining links faster than usual', 'did we just lose a batch
// of links' and 'is a competitor running a link campaign'. Free.
func (c *Client) GetLinkVelocity(ctx context.Context, in GetLinkVelocityInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_link_velocity", in)
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
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// GetLlmResponsesInput is the input to get_llm_responses.
type GetLlmResponsesInput struct {
	ClientID      string         `json:"clientId"`
	Prompt        string         `json:"prompt"`
	Platforms     []string       `json:"platforms,omitempty"`
	WebSearch     *bool          `json:"webSearch,omitempty"`
	Country       string         `json:"country,omitempty"`       // Where the web search runs from (ChatGPT, Claude and Perplexity's Sonar models).
	Models        map[string]any `json:"models,omitempty"`        // A model per platform; defaults gpt-5-mini, claude-sonnet-4-6, gemini-2.5-flash, sonar.
	MaxCostMicros *int64         `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool          `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetLlmResponses calls get_llm_responses: Ask ChatGPT, Claude, Gemini and
// Perplexity through DataForSEO. Asks one question (up to 500 characters) live
// on the chosen models through DataForSEO — chat_gpt, claude, gemini,
// perplexity, and chat_gpt_app: the ChatGPT app itself as a person sees it,
// with the sources it cited and the searches it ran (about $0.004) — with web
// search on by default, and returns each answer's text, the pages it cited,
// the follow-up searches it ran (fanOutQueries, where the platform reports
// them), and which of the client's tracked brands it named and whether the
// client was named and cited. Needs no assistant keys of the agency's own;
// nothing is stored. Use it to explore a question before tracking it. Answers
// 'what do the big models say when asked this' and 'what does ChatGPT search
// for to answer it'. A paid lookup on the client's account unless anyone asked
// the same thing recently (served says which); quote it with estimate_cost.
// (one purchase per model; about $0.01–0.10 each with web search)
func (c *Client) GetLlmResponses(ctx context.Context, in GetLlmResponsesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_llm_responses", in)
}

// GetLocalCategoriesInput is the input to get_local_categories.
type GetLocalCategoriesInput struct {
	ClientID      string   `json:"clientId"`
	Lat           float64  `json:"lat"`
	Lng           float64  `json:"lng"`
	RadiusKm      *float64 `json:"radiusKm,omitempty"`
	Categories    []string `json:"categories,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetLocalCategories calls get_local_categories: Business categories in an
// area. The Google business categories used around a point, with how many
// businesses use each, and the attributes common among them — the categories a
// listing should name, and the gaps. Narrow it with categories to see what
// businesses in those also list. Answers 'which secondary categories do the
// dentists here use'. A paid lookup on the client's account unless anyone read
// the same thing recently; quote it with estimate_cost.
func (c *Client) GetLocalCategories(ctx context.Context, in GetLocalCategoriesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_local_categories", in)
}

// GetLocalSEOAuditInput is the input to get_local_seo_audit.
type GetLocalSEOAuditInput struct {
	ClientID           string   `json:"clientId"`
	LocationID         string   `json:"locationId"`
	UseListing         *bool    `json:"useListing,omitempty"`
	PostsTaskID        string   `json:"postsTaskId,omitempty"`        // A get_business_posts taskId for this listing; collecting it is free.
	CompetitorRadiusKm *float64 `json:"competitorRadiusKm,omitempty"` // Radius for competitor density; 0 skips the paid business search.
	Directories        *bool    `json:"directories,omitempty"`
	Duplicates         *bool    `json:"duplicates,omitempty"` // Search the listings within 1 km for duplicates at this address or phone (one paid business search).
	City               string   `json:"city,omitempty"`
	Service            string   `json:"service,omitempty"`
	Model              string   `json:"model,omitempty"`
	MaxCostMicros      *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun             *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetLocalSEOAudit calls get_local_seo_audit: Local SEO audit of a location.
// Scores one location's local SEO 0–100 (a heuristic, said so in the answer)
// from six weighted parts, each check full, partial or low with the reason, or
// null when not measured — then left out and the other weights rescaled: the
// Google Business Profile (25: primary category against the homepage, about
// four secondary categories, posts, photos, Q&A, hours), reviews (20: at least
// 10, the rating, the last three months, six months' monthly flow and gaps
// over three weeks, owner response rate, and a fake-review flag when two of
// six patterns match), the site (20: city and service in the title and H1,
// visible NAP, a tel: link, a map embed, location pages over 60% unique by the
// city-swap test, gates at 30 and 50 location pages), NAP and citations (15,
// the check_nap_consistency comparison; directories: true adds the directory
// citations), LocalBusiness schema (10: a subtype, geo to five decimals,
// priceRange under 100 characters, a unique @id per location with branchOf)
// and links (10); plus a rating spread that leans to five stars, review
// gating, service pages, location pages within three clicks, hours on the
// site, trust badges, the PostalAddress parts, the subtype that fits the
// vertical, areaServed, retired VehicleListing and SpecialAnnouncement markup
// and, with duplicates, other listings at the same address or phone.
// gbpChecklist scores the 25-field profile checklist (2/1/0 per field, null
// where DataForSEO returns nothing — verified status, Q&A, videos, products,
// photo dates — weighted by vertical, banded 90/75/50/25); grid reads back the
// latest grid scan for the main service (share of local voice, visibility
// points, the worst quadrant, a heatmap); mapsHealth combines profile 25, grid
// 20, reviews 20, NAP 15, schema 10 and rival standing 10. Detects storefront,
// service-area business (which skips the map and address checks) or hybrid,
// and the vertical; adds listing completeness 0–100 and, with
// competitorRadiusKm, the competitors per km² in the primary category. Reviews
// come from sync_local_reviews; posts from a get_business_posts task passed as
// postsTaskId. Answers 'how good is our local SEO' and 'what should we fix
// first for the map pack'. useListing reads the live listing and its Q&A, and
// competitorRadiusKm one business search, all paid: quote with estimate_cost.
// Free with useListing false and competitorRadiusKm 0.
func (c *Client) GetLocalSEOAudit(ctx context.Context, in GetLocalSEOAuditInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_local_seo_audit", in)
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
// ranking in the top 3 and top 10, a summary (share of local voice with its
// band — dominant 80+, strong 60+, moderate 40+, weak 20+, critical —
// visibility points 3 for the top three and 1 for fourth to tenth, each
// compass quadrant and the worst, and the grid as rows north to south for a
// heatmap, with a text version), and a competitor table: every business the
// grid saw, with its share of local voice (the share of points where it holds
// a top-3 slot), points seen, average position, rating and reviews — plus an
// analysis of the location's reviews (rating, volume, recency, response rate).
// Answers 'where in town do we show up for this search'. scanned is false when
// that term has never been scanned there: no grid, rather than a grid of
// blanks that would read as invisible everywhere. Free.
func (c *Client) GetLocalVisibility(ctx context.Context, in GetLocalVisibilityInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_local_visibility", in)
}

// GetLogOrphansInput is the input to get_log_orphans.
type GetLogOrphansInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetLogOrphans calls get_log_orphans: Pages crawlers find that the audit does
// not. Pages search and AI crawlers fetched successfully in the connected logs
// that the latest site crawl never reached — nothing on the site links to
// them, yet they are crawled (old campaigns, forgotten sections, pages only a
// sitemap or another site names) — with hits, which crawlers and when last
// seen. Answers 'which orphan pages is Google still crawling'. Free.
func (c *Client) GetLogOrphans(ctx context.Context, in GetLogOrphansInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_log_orphans", in)
}

// GetLongTailCandidatesInput is the input to get_long_tail_candidates.
type GetLongTailCandidatesInput struct {
	ClientID       string `json:"clientId"`
	Days           *int64 `json:"days,omitempty"` // Days of Search Console history, ending on the newest day Google has finished.
	MinImpressions *int64 `json:"minImpressions,omitempty"`
	Limit          *int64 `json:"limit,omitempty"`
}

// GetLongTailCandidates calls get_long_tail_candidates: Queries with
// impressions and no page of their own. Search Console queries the site is
// shown for where the page Google picked is about something else — its title
// and H1 (from the latest crawl) cover under half the query's words —
// clustered by shared words into write candidates, biggest first, each with
// its queries, impressions and the pages ranking now (where to link the new
// page from). Branded queries are left out. Answers 'what should we write
// next' and 'which searches have no page'. Free.
func (c *Client) GetLongTailCandidates(ctx context.Context, in GetLongTailCandidatesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_long_tail_candidates", in)
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

// GetMapsResultsInput is the input to get_maps_results.
type GetMapsResultsInput struct {
	ClientID      string   `json:"clientId"`
	Keyword       string   `json:"keyword"`
	Surface       string   `json:"surface,omitempty"`
	Zoom          *int64   `json:"zoom,omitempty"`
	Depth         *int64   `json:"depth,omitempty"`
	LocationCode  *int64   `json:"locationCode,omitempty"` // Default: the client's project country.
	LanguageCode  string   `json:"languageCode,omitempty"`
	Lat           *float64 `json:"lat,omitempty"`
	Lng           *float64 `json:"lng,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetMapsResults calls get_maps_results: Google Maps results from one spot.
// The Google Maps (or Local Finder) results for a keyword as searched from one
// point — or the client's country — with every business's rating, reviews,
// category, site and ids, and where the client sits. One lookup, where a grid
// scan is many. Answers 'who shows on the map for this from downtown'. A paid
// lookup on the client's account unless anyone read the same thing recently;
// quote it with estimate_cost.
func (c *Client) GetMapsResults(ctx context.Context, in GetMapsResultsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_maps_results", in)
}

// GetMatomoReportInput is the input to get_matomo_report.
type GetMatomoReportInput struct {
	ClientID string `json:"clientId"`
	Report   string `json:"report,omitempty"`
	Channel  string `json:"channel,omitempty"` // organic: visits from search engines; ai: from AI assistants (Matomo 5.3+); all: every channel.
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Compare  *bool  `json:"compare,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetMatomoReport calls get_matomo_report: Visits, pages and goals (Matomo).
// For a client that measures with Matomo instead of GA4, read live from their
// Matomo. report: overview (visits, unique visitors when the server processes
// them for ranges, actions, pageviews, bounce rate, average time, conversions,
// conversion rate, revenue); channels (visits, bounce rate, conversions and
// revenue by referrer channel: search, social, ai, direct, website, campaign —
// Matomo's own split); landing_pages (entrances, bounce rate, actions per
// entrance, conversions and revenue credited to the entry page); pages
// (pageviews, visits, average time, bounce and exit rate by page URL); goals
// (conversions, conversion rate and revenue per goal, and in total). channel
// limits overview, landing_pages, pages and goals to search engines (organic)
// or AI assistants (ai) through a Matomo segment; Matomo answers a segment
// only once it has processed it, and says so. Compares with the same number of
// days before unless compare is false; defaults to the 28 days ending
// yesterday. Rows carry key, label, current and previous (null where that
// period had no row). Figures Matomo did not send are null, never zero.
// Refuses with a sentence when Matomo is not connected (connect_integration
// with provider matomo). Answers 'is organic traffic growing' and 'which
// landing pages convert' for a Matomo site. Free.
func (c *Client) GetMatomoReport(ctx context.Context, in GetMatomoReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_matomo_report", in)
}

// GetMediaReportInput is the input to get_media_report.
type GetMediaReportInput struct {
	ClientID   string `json:"clientId"`
	CrawlID    string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	CheckSizes *bool  `json:"checkSizes,omitempty"`
}

// GetMediaReport calls get_media_report: Images and video. Images and video
// across a crawl: alt text coverage, formats (and the share in AVIF or WebP),
// pages whose first image — usually the Largest Contentful Paint element — is
// lazy-loaded, how many images the image sitemap lists, embedded video with
// and without VideoObject markup and video sitemap coverage, plus the image
// checks' page counts. checkSizes asks the server for up to 100 image sizes
// (HEAD requests to the client's own hosts and its image CDNs) and lists the
// heaviest, with each judged against its role's band (thumbnail: target under
// 50KB, warn over 100KB, critical over 200KB; content: 100/200/500KB; hero,
// the page's first image: 200/300/700KB), and reads up to 20 product images'
// XMP for an IPTC DigitalSourceType that declares them AI-generated
// (trainedAlgorithmicMedia, compositeSynthetic, algorithmicMedia,
// compositeWithTrainedAlgorithmicMedia), and grades each image's browser
// caching (cache: long at a max-age of a week or more or immutable, short
// under a week, none with no-store, no-cache or no lifetime at all). Always:
// image hosts and the CDN each shows (by name, or by response headers when
// sizes are checked), the share on a separate host, and how pages lazy-load
// (native, perfmatters, EWWW, a generic script loader or none; script loaders
// are not reported as missing lazy loading). Answers 'are our images slowing
// the site down', 'which images lack alt text' and 'can our videos show in
// video results'. Free.
func (c *Client) GetMediaReport(ctx context.Context, in GetMediaReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_media_report", in)
}

// GetNaverKeywordVolumesInput is the input to get_naver_keyword_volumes.
type GetNaverKeywordVolumesInput struct {
	ClientID       string   `json:"clientId"`
	Keywords       []string `json:"keywords"`
	IncludeRelated *bool    `json:"includeRelated,omitempty"`
}

// GetNaverKeywordVolumes calls get_naver_keyword_volumes: Naver's own search
// volume. Monthly searches on Naver, PC and mobile, from the keyword tool in
// the agency's Naver Search Ad account, for up to 50 terms, with Naver's
// related terms when includeRelated is true. The only source of Naver demand:
// no vendor resells it. Terms too rare are reported by Naver as under ten and
// flagged so, not given a number. Answers 'how often is this searched on
// Naver'. Free; needs the agency's Naver Search Ad credentials.
func (c *Client) GetNaverKeywordVolumes(ctx context.Context, in GetNaverKeywordVolumesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_naver_keyword_volumes", in)
}

// GetNextStepsInput is the input to get_next_steps.
type GetNextStepsInput struct {
	ClientID string `json:"clientId"`
}

// GetNextSteps calls get_next_steps: What to do next for a site. Answers 'what
// is left to do for this site, and what call does it'. The onboarding
// checklist with live status — done, todo, in progress, failed or blocked and
// why — each with the exact tool and arguments to call next; then the optional
// connections that unlock more tools (Bing Webmaster key, PageSpeed / Chrome
// UX Report key, alert delivery, rank tracker keywords, AI visibility
// schedule, analytics goals, crawl schedule), each with its status and call;
// then, once the site is set up, the weekly loop's calls with their arguments.
// Reads only; changes nothing. Free.
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

// GetOutboundDomainsInput is the input to get_outbound_domains.
type GetOutboundDomainsInput struct {
	ClientID string `json:"clientId"`
	Domain   string `json:"domain,omitempty"` // Any domain; default the client's own.
	Limit    *int64 `json:"limit,omitempty"`
}

// GetOutboundDomains calls get_outbound_domains: Domains a site links out to.
// The domains a site (the client's or any other) links out to, with how many
// links, how many of them dofollow and the dofollow share, and from how many
// of its pages. A site that already links to your competitors is an outreach
// prospect; one linking to dead or spammy domains is a risk. Read from our own
// crawl of the site; when we have not crawled it, its homepage and up to 15
// pages it links to are fetched now. Free. Answers 'who does this site link
// to' and 'which resources does this niche cite'.
func (c *Client) GetOutboundDomains(ctx context.Context, in GetOutboundDomainsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_outbound_domains", in)
}

// GetOutreachMailboxesInput is the input to get_outreach_mailboxes.
type GetOutreachMailboxesInput struct {
	ClientID string `json:"clientId"`
}

// GetOutreachMailboxes calls get_outreach_mailboxes: Mailboxes outreach is
// sent from. The Gmail and Microsoft 365 mailboxes connected for the client's
// outreach, with when replies were last checked and any error, and the links
// where a person connects another (they sign in themselves; nothing here can).
// Answers 'which mailbox do we send from' and 'is outreach email working'.
// Free.
func (c *Client) GetOutreachMailboxes(ctx context.Context, in GetOutreachMailboxesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_outreach_mailboxes", in)
}

// GetOutreachStatusInput is the input to get_outreach_status.
type GetOutreachStatusInput struct {
	ClientID string `json:"clientId"`
}

// GetOutreachStatus calls get_outreach_status: Where each outreach
// conversation stands. Every prospect the client has written to, with its
// stage, contact, when it replied, and each email of its sequence: sent,
// scheduled, failed or called off. Answers 'who have we emailed', 'who
// replied' and 'what is going out this week'. Free.
func (c *Client) GetOutreachStatus(ctx context.Context, in GetOutreachStatusInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_outreach_status", in)
}

// GetPageBacklinksInput is the input to get_page_backlinks.
type GetPageBacklinksInput struct {
	ClientID      string `json:"clientId"`
	URL           string `json:"url"`
	DofollowOnly  *bool  `json:"dofollowOnly,omitempty"`
	Limit         *int64 `json:"limit,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetPageBacklinks calls get_page_backlinks: Who links to one page. The sites
// linking to one exact page — one row per linking domain, live links only,
// best domains first — with the linking page, its title, the anchor, dofollow
// and the domain's rank. Run it on a competitor's dead page from
// get_competitor_broken_pages to get the people to write to (broken-link
// building), or on any page to see who cites it. Answers 'who links to that
// dead page' and 'who would link to ours instead'. A paid DataForSEO read
// unless anyone read the same page this week; quote it with estimate_cost.
func (c *Client) GetPageBacklinks(ctx context.Context, in GetPageBacklinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_page_backlinks", in)
}

// GetPageDetailsInput is the input to get_page_details.
type GetPageDetailsInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	URL      string `json:"url"`
}

// GetPageDetails calls get_page_details: Everything the crawl read on a page.
// One crawled page in full: its H2/H3 outline, word count, published and
// modified dates, structured data as written, custom extraction values, the
// addresses that redirect to it, what JavaScript changed when it was rendered,
// and its place in the link graph (inlinks, outlinks, click depth, internal
// PageRank, anchors). Answers 'what does the crawler see on this page'. Use
// get_page_issues for its issues. Free.
func (c *Client) GetPageDetails(ctx context.Context, in GetPageDetailsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_page_details", in)
}

// GetPageIndexingReasonsInput is the input to get_page_indexing_reasons.
type GetPageIndexingReasonsInput struct {
	ClientID string `json:"clientId"`
	Template string `json:"template,omitempty"` // A template as the report writes it, e.g. /products/*.
	Reason   string `json:"reason,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetPageIndexingReasons calls get_page_indexing_reasons: Why pages are not
// indexed. Search Console's Page indexing report, rebuilt (its API has none):
// every sitemap and crawled page with one reason it is or is not in Google —
// noindex, blocked by robots.txt, 404, server error, redirect, soft 404,
// Google chose another canonical, crawled or discovered but not indexed,
// unknown to Google — from URL Inspection (2,000 a day per property, worked
// through the sitemap a batch a day) and our crawl, which stands in, marked,
// for pages not inspected yet. Rolled up per template (path pattern) with the
// share indexed and the main reason, and a fix for each reason. Filter by
// template or reason. Answers 'why isn't this page indexed' and 'which
// templates are kept out of Google'. Free.
func (c *Client) GetPageIndexingReasons(ctx context.Context, in GetPageIndexingReasonsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_page_indexing_reasons", in)
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

// GetParameterReportInput is the input to get_parameter_report.
type GetParameterReportInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
}

// GetParameterReport calls get_parameter_report: URL parameters and crawl
// waste. Every query parameter the crawl met — sort orders, filters, tracking,
// pagination — with how many pages carry it, how many of those are
// canonicalised to another page (crawl waste), their issues, and whether to
// strip it from crawls (set_crawl_settings stripParams), keep it, or review
// it. Answers 'why is the crawl full of ?sort= pages'. Free.
func (c *Client) GetParameterReport(ctx context.Context, in GetParameterReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_parameter_report", in)
}

// GetParasiteRiskInput is the input to get_parasite_risk.
type GetParasiteRiskInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"`
}

// GetParasiteRisk calls get_parasite_risk: Sections that read like rented-out
// pages. Site reputation abuse risk per section (first path segment) from the
// crawl's text and links: sponsored or partner phrases per page (1.0 or more
// is high risk), commerce phrases per page (2.0 or more), outbound links with
// affiliate parameters such as tag=, aff_id= or ref_= per page (3.0 or more),
// commerce and affiliate together (high risk), and a section whose commerce
// rate is more than twice the site's mean (drift — a section that has turned
// into something else). Each section with its rates, signals, risk and the
// worst pages. A heuristic, labelled as one. Answers 'is any part of our site
// at risk under Google's site reputation abuse policy' and 'which section has
// been turned into coupon or affiliate pages'. Free.
func (c *Client) GetParasiteRisk(ctx context.Context, in GetParasiteRiskInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_parasite_risk", in)
}

// GetPersonalKeywordDifficultyInput is the input to get_personal_keyword_difficulty.
type GetPersonalKeywordDifficultyInput struct {
	ClientID     string   `json:"clientId"`
	Keywords     []string `json:"keywords"`
	LocationCode *int64   `json:"locationCode,omitempty"` // Google location code (2840 United States, 2826 United Kingdom). Default: the client's project.
	LanguageCode string   `json:"languageCode,omitempty"` // Language code (en, es). Default: the client's project.
}

// GetPersonalKeywordDifficulty calls get_personal_keyword_difficulty: How hard
// a keyword is for this site. Keyword difficulty adjusted for the client's own
// site: its authority against the first page's, how much of the topic it
// already ranks for, and whether it already ranks for the keyword — each
// adjustment returned with its reason, so the number can be argued with.
// Answers 'which of these can we actually win'. Free: read from stored first
// pages, link counts, authority and rankings; a factor with no data behind it
// is left out and said so.
func (c *Client) GetPersonalKeywordDifficulty(ctx context.Context, in GetPersonalKeywordDifficultyInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_personal_keyword_difficulty", in)
}

// GetPlatformReviewsInput is the input to get_platform_reviews.
type GetPlatformReviewsInput struct {
	ClientID      string `json:"clientId"`
	Platform      string `json:"platform"`
	Business      string `json:"business,omitempty"`
	TaskID        string `json:"taskId,omitempty"`
	Depth         *int64 `json:"depth,omitempty"`         // How many reviews to read; billed per 10 (Tripadvisor) or 20 (Trustpilot).
	LocationCode  *int64 `json:"locationCode,omitempty"`  // Default: the client's project country.
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetPlatformReviews calls get_platform_reviews: A business's Tripadvisor or
// Trustpilot reviews. Reviews of a business on Tripadvisor (business: its name
// and town, or a Tripadvisor URL path such as /Restaurant_Review-g…) or
// Trustpilot (business: its domain), newest first, with the platform's own
// rating and review count, the star spread over the reviews read, the average
// and the owner's response rate — for the cross-platform comparison beside
// Google's reviews (get_business_reviews, sync_local_reviews). Both are
// queued: the first call (with business) is the paid one and returns a taskId;
// calling again with the taskId collects it for free (status pending until
// ready). Answers 'how are we rated on Tripadvisor' and 'what does Trustpilot
// say about this competitor'. A paid lookup on the client's account unless
// anyone read the same thing recently; quote it with estimate_cost.
func (c *Client) GetPlatformReviews(ctx context.Context, in GetPlatformReviewsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_platform_reviews", in)
}

// GetPortfolioMetricsInput is the input to get_portfolio_metrics.
type GetPortfolioMetricsInput struct {
	PortfolioID string `json:"portfolioId"`
}

// GetPortfolioMetrics calls get_portfolio_metrics: A portfolio's figures.
// Traffic, rankings, issues and AI visibility across a portfolio, with each
// member's own row: Google clicks over the newest 28 days and the 28 before,
// tracked keywords in the top 3 and 10 and the average position, site health
// and critical issues from the newest crawl, and how often AI assistants named
// the brand. Each total says how many members it covers; a member without the
// data is left out, never counted as zero. A site member is read for that
// site; AI visibility is per client. Answers 'how are the dental practices
// doing'. Free.
func (c *Client) GetPortfolioMetrics(ctx context.Context, in GetPortfolioMetricsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_portfolio_metrics", in)
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

// GetPreloadReadinessInput is the input to get_preload_readiness.
type GetPreloadReadinessInput struct {
	ClientID string `json:"clientId"`
	URL      string `json:"url,omitempty"`
}

// GetPreloadReadiness calls get_preload_readiness: How ready a page is for
// instant navigation. A 0–100 heuristic from one page's markup and headers, 25
// points each: speculation rules (a <script type="speculationrules"> block or
// a Speculation-Rules header) for the next navigation; fetchpriority="high" on
// the image that is probably the Largest Contentful Paint; nothing that keeps
// it out of the back/forward cache (Cache-Control: no-store, an unload
// listener); and no deprecated <link rel="prerender">. With what to change for
// each part missed. A static reading, not a measurement: lab diagnostics come
// from page_lab. url defaults to the homepage and must be on the client's
// site. Answers 'can the next page load instantly' and 'why does Back reload
// the page'. Free.
func (c *Client) GetPreloadReadiness(ctx context.Context, in GetPreloadReadinessInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_preload_readiness", in)
}

// GetPriceListInput is the input to get_price_list.
type GetPriceListInput struct {
}

// GetPriceList calls get_price_list: What everything costs. Every paid data
// operation with its most-per-request price in micro-USD and how long an
// answer stays free once anyone has bought it (the shared archive), plus which
// tools spend. A tool's own quote for given arguments is estimate_cost, or the
// tool itself with dryRun: true. Answers 'how much does this cost' and 'what
// is cached for how long'. Free.
func (c *Client) GetPriceList(ctx context.Context, in GetPriceListInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_price_list", in)
}

// GetQueryIntentsInput is the input to get_query_intents.
type GetQueryIntentsInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"` // Days of Search Console history, ending on the newest day Google has finished.
	By       string `json:"by,omitempty"`
	Brand    string `json:"brand,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetQueryIntents calls get_query_intents: Search traffic by intent or topic.
// Google web clicks, impressions and position grouped by what each Search
// Console query asks for — informational, commercial, transactional or
// navigational (navigational: it names the client, from the brief's brand
// terms) — or by topic (queries sharing their most common words). Every query
// is tagged after each Search Console sync; query_data can also group or
// filter by intent, topic and branded. Answers 'how much of our traffic is
// people ready to buy' and 'which topics bring the clicks'. Free.
func (c *Client) GetQueryIntents(ctx context.Context, in GetQueryIntentsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_query_intents", in)
}

// GetQueryPageSwapsInput is the input to get_query_page_swaps.
type GetQueryPageSwapsInput struct {
	ClientID       string `json:"clientId"`
	Days           *int64 `json:"days,omitempty"`       // Length of each period, ending on the newest day Google has finished counting.
	Compare        string `json:"compare,omitempty"`    // previous: the same number of days just before. year: the same days a year earlier.
	SearchType     string `json:"searchType,omitempty"` // Search Console's search type: web (default), image, video, news, discover or googleNews. Discover and Google News report pages but no queries.
	Brand          string `json:"brand,omitempty"`      // non_branded is what SEO work moves; branded is people already looking for the client.
	MinImpressions *int64 `json:"minImpressions,omitempty"`
	Limit          *int64 `json:"limit,omitempty"`
}

// GetQueryPageSwaps calls get_query_page_swaps: Queries answered by a
// different page. Queries where Google now sends searchers to a different page
// of the site than in the previous period, with both pages and their clicks.
// Answers 'is Google confused about which page to rank' — a swap is usually
// two pages competing (merge them or fix internal links) or a new page taking
// over. Every Search Console query, not only tracked keywords (get_url_swaps
// covers those). Each swap names the registry's owner for the query and
// whether Google's current page is it. searchType picks web (default), image,
// video or news. Free.
func (c *Client) GetQueryPageSwaps(ctx context.Context, in GetQueryPageSwapsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_query_page_swaps", in)
}

// GetQuestionsInput is the input to get_questions.
type GetQuestionsInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
	Status   string `json:"status,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetQuestions calls get_questions: The questions people ask, from every
// source. Questions about the client's subject merged from People Also Ask and
// related searches on captured results pages, the client's AI visibility
// prompts, Reddit and Quora threads ranking for tracked keywords, Search
// Console queries phrased as questions, and the site's own search box (GA4) —
// phrasings of one question deduplicated — each with its sources, volume where
// a source measured it (impressions, site searches, or volume, kept apart) and
// whether a page of the site answers it: answered (a heading asks or states
// it), mentioned (the text has its words) or unanswered, from the latest
// crawl. sources says which could be read. Answers 'what should our FAQ cover'
// and 'which questions have no page'. Free.
func (c *Client) GetQuestions(ctx context.Context, in GetQuestionsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_questions", in)
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

// GetRankedKeywordsInput is the input to get_ranked_keywords.
type GetRankedKeywordsInput struct {
	ClientID      string `json:"clientId"`
	Engine        string `json:"engine,omitempty"`
	Target        string `json:"target,omitempty"` // A domain (competitor.com), a subfolder (competitor.com/blog/) or a page URL. Default: the client's own domain.
	Scope         string `json:"scope,omitempty"`  // auto reads a bare domain as the domain, a path ending in / as a subfolder, anything else as one URL.
	Paid          *bool  `json:"paid,omitempty"`
	ResultType    string `json:"resultType,omitempty"` // ai_overview_reference: keywords whose AI Overview cites the domain; featured_snippet; local_pack; paid (same as paid: true).
	Change        string `json:"change,omitempty"`
	MinPosition   *int64 `json:"minPosition,omitempty"`
	MaxPosition   *int64 `json:"maxPosition,omitempty"`
	MinVolume     *int64 `json:"minVolume,omitempty"`
	Intent        string `json:"intent,omitempty"`
	Contains      string `json:"contains,omitempty"` // Only keywords containing these words.
	Sort          string `json:"sort,omitempty"`
	Limit         *int64 `json:"limit,omitempty"`
	Offset        *int64 `json:"offset,omitempty"`
	LocationCode  *int64 `json:"locationCode,omitempty"`  // Default: the client's project (2840 United States).
	LanguageCode  string `json:"languageCode,omitempty"`  // Default: the client's project.
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetRankedKeywords calls get_ranked_keywords: What a domain ranks for. Every
// keyword a domain, subfolder or page ranks for in Google — position (and the
// one before), URL, estimated traffic, volume, CPC, difficulty, intent and
// SERP features — filtered and sorted, paginated. paid: true (or resultType
// paid) lists the keywords it buys ads on, with the ad; resultType
// ai_overview_reference lists the keywords whose AI Overview cites it,
// featured_snippet those where it holds the snippet. change new / up / down
// lists its recent ranking movements, and lost the keywords it no longer ranks
// for (with the position it last held). engine amazon reads the keywords an
// Amazon product ranks for (target its ASIN); engine app_store or google_play
// the keywords an app ranks for (target the numeric App Store id or the Play
// package name; US English only) — position, volume and the listing, without
// the Google-only fields, filters or change. Answers 'what does this
// competitor rank for', 'which keywords drive their traffic', 'what did they
// just start ranking for' and 'which searches find this product or app'. A
// paid lookup on the client's account unless anyone read the same thing
// recently (spend says which); quote it with estimate_cost.
func (c *Client) GetRankedKeywords(ctx context.Context, in GetRankedKeywordsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_ranked_keywords", in)
}

// GetRecipeInput is the input to get_recipe.
type GetRecipeInput struct {
	ClientID string `json:"clientId"`
	Name     string `json:"name,omitempty"`
}

// GetRecipe calls get_recipe: A named plan of calls. An ordered call plan for
// a common job, with arguments filled in from the client (its domain, rivals,
// a keyword it ranks for) and the ones left to fill named in fill. Recipes:
// weekly_review, new_article, technical_audit, site_migration,
// competitor_teardown, ai_visibility_baseline. A plan, not automation: make
// each call, read it, and skip a step the one before says is unnecessary. With
// no name, lists them. Answers 'what's the process for a technical audit'.
// Free.
func (c *Client) GetRecipe(ctx context.Context, in GetRecipeInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_recipe", in)
}

// GetReferringDomainsInput is the input to get_referring_domains.
type GetReferringDomainsInput struct {
	ClientID      string `json:"clientId"`
	Domain        string `json:"domain,omitempty"` // Any domain; default the client's own.
	Source        string `json:"source,omitempty"`
	Limit         *int64 `json:"limit,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetReferringDomains calls get_referring_domains: Referring domains. The
// domains linking to a domain (the client's or any competitor's), with how
// many links each sends, whether they follow, the linking domain's authority
// (0–100) and when the links were first and last seen. Answers 'who links to
// us', 'who links to them' and 'which sites are worth asking'. source index is
// free: our own link index (our crawls, client crawls' outlinks, Common Crawl,
// Bing). auto answers from the index when it holds 50+ links to the domain
// (or, for referring domains, a complete pass of Common Crawl's domain graph),
// and otherwise asks the vendor — a paid lookup on the client's account unless
// anyone bought it in the last week; coverage says which answered. vendor
// always asks the vendor.
func (c *Client) GetReferringDomains(ctx context.Context, in GetReferringDomainsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_referring_domains", in)
}

// GetReferringNetworksInput is the input to get_referring_networks.
type GetReferringNetworksInput struct {
	ClientID      string `json:"clientId"`
	Target        string `json:"target,omitempty"` // Default: the client's site.
	Type          string `json:"type,omitempty"`
	Limit         *int64 `json:"limit,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetReferringNetworks calls get_referring_networks: Which networks the links
// come from. The IP addresses or subnets (type ip or subnet) the client's
// links — or any target's — come from, from DataForSEO's referring networks:
// per network its backlinks, referring domains and main domains, rank, first
// seen and lost date, with the networks carrying links from five or more
// separate domains flagged. Many unrelated sites on one address is the shape
// of a private blog network; a shared host puts honest sites together too, so
// a flag means look at the sites, not disavow. Answers 'are our links from a
// link network' and 'how many of these sites share a server'. One paid
// DataForSEO request; quote it with estimate_cost.
func (c *Client) GetReferringNetworks(ctx context.Context, in GetReferringNetworksInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_referring_networks", in)
}

// GetRenderingReportInput is the input to get_rendering_report.
type GetRenderingReportInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	Limit    *int64 `json:"limit,omitempty"`
}

// GetRenderingReport calls get_rendering_report: What JavaScript changes.
// Pages the crawl rendered with JavaScript, compared with their raw HTML:
// words and links that exist only after JavaScript runs, and a title,
// canonical, meta robots, H1, description or structured data that JavaScript
// changes — what a search engine or an AI crawler that does not run scripts
// never sees. Answers 'does our content depend on JavaScript'. Turn rendering
// on with set_crawl_settings render. Free.
func (c *Client) GetRenderingReport(ctx context.Context, in GetRenderingReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_rendering_report", in)
}

// GetReportDataInput is the input to get_report_data.
type GetReportDataInput struct {
	ClientID  string `json:"clientId"`
	From      string `json:"from"` // The period's first day (UTC).
	To        string `json:"to"`   // The period's last day (UTC), inclusive.
	CompareTo string `json:"compareTo,omitempty"`
}

// GetReportData calls get_report_data: Every figure a monthly report needs.
// One read of a client's period, from..to inclusive, against the period before
// (compareTo previous_period, the default) or the same days a year earlier
// (previous_year): searchConsole (clicks, impressions, CTR and position in
// total, branded and non-branded, with the non-branded queries and pages that
// moved most), analytics (GA4 organic sessions, key events and revenue, or the
// first-party web analytics' visits), revenue (imported orders by currency,
// organic and AI channels, attribution coverage — as get_seo_revenue),
// rankings (each tracker's summary and movers between the period's ends),
// audit (the latest crawl in the period, its health, and new and fixed issues
// against the last crawl before it), backlinks (gained and lost), aiVisibility
// (mention rate against the comparison), annotations (the timeline, with
// Google's updates) and alerts fired. A section is null with unavailableReason
// when the client lacks that source. Answers 'pull October's numbers for the
// report'. Use it first; the single tools go deeper on one section. Free.
func (c *Client) GetReportData(ctx context.Context, in GetReportDataInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_report_data", in)
}

// GetRobotsHistoryInput is the input to get_robots_history.
type GetRobotsHistoryInput struct {
	ClientID string `json:"clientId"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetRobotsHistory calls get_robots_history: How robots.txt and the sitemaps
// changed. The robots.txt each crawl read, newest first, with the lines added
// and removed against the version before (a new Disallow, a dropped Sitemap
// line), and per crawl how many sitemaps and entries were listed. Answers 'who
// changed robots.txt and when', 'why did Google stop crawling /blog' and 'did
// the sitemap shrink'. Free.
func (c *Client) GetRobotsHistory(ctx context.Context, in GetRobotsHistoryInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_robots_history", in)
}

// GetRunStatusInput is the input to get_run_status.
type GetRunStatusInput struct {
	ClientID string `json:"clientId"`
	Job      string `json:"job"`
	RunID    string `json:"runId"`
}

// GetRunStatus calls get_run_status: Get the status of a run. How a run
// started with refresh_rankings (job 'rank'), refresh_domain_data (job
// 'domain') or recheck_links (job 'links', including a browser re-check with
// render) is going: running, done or failed, and — when it failed or captured
// only some keywords — the reason, written for an agency to act on (a missing
// vendor key, an unreachable site). Answers 'did that capture work'. A run id
// that is not this client's, or that has aged out, reads as unknown.
func (c *Client) GetRunStatus(ctx context.Context, in GetRunStatusInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_run_status", in)
}

// GetSearchAnomaliesInput is the input to get_search_anomalies.
type GetSearchAnomaliesInput struct {
	ClientID     string   `json:"clientId"`
	LookbackDays *int64   `json:"lookbackDays,omitempty"`
	Scopes       []string `json:"scopes,omitempty"`
}

// GetSearchAnomalies calls get_search_anomalies: Unusual days in search
// traffic, explained. Days when Google clicks fell or jumped well outside
// their normal range — for the whole site, each section (first path segment)
// and each query topic — against a robust same-weekday baseline (median and
// MAD over eight weeks), consecutive days joined into one episode. Each names
// the likely cause from the timeline with a confidence: a crawl regression
// (pages gone 404 or newly noindex), a deploy or crawl change, a Google
// update, a lost SERP feature, or a page swap (the ranking URL changed). Also
// lists what the daily check already raised. Answers 'why did traffic drop on
// Tuesday' and 'is anything wrong with search'. Free.
func (c *Client) GetSearchAnomalies(ctx context.Context, in GetSearchAnomaliesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_search_anomalies", in)
}

// GetSearchAppearanceInput is the input to get_search_appearance.
type GetSearchAppearanceInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`    // Length of each period, ending on the newest day Google has finished counting.
	Compare  string `json:"compare,omitempty"` // previous: the same number of days just before. year: the same days a year earlier.
}

// GetSearchAppearance calls get_search_appearance: Search traffic by result
// kind. Google web search clicks and impressions by search appearance — rich
// results, videos, FAQ and review snippets, AMP and the other result kinds
// Search Console reports — two periods compared. Answers 'are our rich results
// earning clicks' and 'did we lose a result type'. Collected from the day this
// was added, and back sixteen months once a connection's history has been
// filled. Free.
func (c *Client) GetSearchAppearance(ctx context.Context, in GetSearchAppearanceInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_search_appearance", in)
}

// GetSearchChangesInput is the input to get_search_changes.
type GetSearchChangesInput struct {
	ClientID       string `json:"clientId"`
	Days           *int64 `json:"days,omitempty"`       // Length of each period, ending on the newest day Google has finished counting.
	Compare        string `json:"compare,omitempty"`    // previous: the same number of days just before. year: the same days a year earlier.
	SearchType     string `json:"searchType,omitempty"` // Search Console's search type: web (default), image, video, news, discover or googleNews. Discover and Google News report pages but no queries.
	Dimension      string `json:"dimension,omitempty"`
	Brand          string `json:"brand,omitempty"` // non_branded is what SEO work moves; branded is people already looking for the client.
	Country        string `json:"country,omitempty"`
	Device         string `json:"device,omitempty"`
	MinImpressions *int64 `json:"minImpressions,omitempty"`
	Limit          *int64 `json:"limit,omitempty"`
}

// GetSearchChanges calls get_search_changes: What changed in search. Two
// periods of Google search compared, by query or by page: what is new (shown
// now, not at all before), what was lost, and the biggest gainers and losers
// by clicks. Answers 'which queries did we start ranking for', 'what did we
// lose this week' and 'which pages are growing'. brand narrows to branded or
// non-branded queries (non-branded is what SEO work moves); country and device
// segment it. Periods end on the newest day Google has finished counting.
// searchType reads image, video, news, Discover or Google News instead of web
// (Discover and Google News by page only). Free.
func (c *Client) GetSearchChanges(ctx context.Context, in GetSearchChangesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_search_changes", in)
}

// GetSearchConsoleHourlyInput is the input to get_search_console_hourly.
type GetSearchConsoleHourlyInput struct {
	ClientID   string `json:"clientId"`
	SearchType string `json:"searchType,omitempty"`
	Days       *int64 `json:"days,omitempty"`
}

// GetSearchConsoleHourly calls get_search_console_hourly: Search Console
// clicks and impressions by the hour. Clicks, impressions and position from
// the client's Google Search Console hour by hour over the last 1–10 days
// (Google keeps hours for ten days only), read live, for one search type.
// Hours are Pacific time, as Search Console keys them; the newest are partial
// while Google counts and are marked so. Answers 'did traffic dip right after
// the release' and 'is the core update hitting us today'. Free; returns a
// refusal when Search Console is not connected.
func (c *Client) GetSearchConsoleHourly(ctx context.Context, in GetSearchConsoleHourlyInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_search_console_hourly", in)
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

// GetSearchTimelineInput is the input to get_search_timeline.
type GetSearchTimelineInput struct {
	ClientID   string `json:"clientId"`
	SearchType string `json:"searchType,omitempty"` // Search Console's search type: web (default), image, video, news, discover or googleNews. Discover and Google News report pages but no queries.
	From       string `json:"from,omitempty"`       // Default: 90 days before to.
	To         string `json:"to,omitempty"`         // Default: the newest day archived.
}

// GetSearchTimeline calls get_search_timeline: Search traffic with what
// happened. Daily Google search clicks and impressions (and non-branded
// clicks) beside the client's timeline: deploys, publishes, campaigns and
// notes, and Google's algorithm updates. searchType reads image, video, news,
// Discover or Google News instead of web (non-branded clicks are null for the
// last two, which have no queries). Answers 'what caused this drop' and 'did
// the release help'. Free.
func (c *Client) GetSearchTimeline(ctx context.Context, in GetSearchTimelineInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_search_timeline", in)
}

// GetSecurityReportInput is the input to get_security_report.
type GetSecurityReportInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
}

// GetSecurityReport calls get_security_report: Security headers and
// accessibility. Security and accessibility across a crawl: how many pages
// send HSTS (and with a weak max-age), a Content-Security-Policy and
// X-Content-Type-Options: nosniff, pages with mixed content (HTTP images,
// scripts, styles or frames on HTTPS), the security and accessibility checks'
// page counts, and the audit's Lighthouse runs with their accessibility and
// best-practices scores and failed accessibility audits — colour contrast
// first. Answers 'are our security headers set' and 'what accessibility
// problems does the site have'. Free.
func (c *Client) GetSecurityReport(ctx context.Context, in GetSecurityReportInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_security_report", in)
}

// GetSEOChangesInput is the input to get_seo_changes.
type GetSEOChangesInput struct {
	ClientID       string   `json:"clientId"`
	CrawlID        string   `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	AgainstCrawlID string   `json:"againstCrawlId,omitempty"`
	Fields         []string `json:"fields,omitempty"`
	Severity       string   `json:"severity,omitempty"`
	Template       string   `json:"template,omitempty"` // A template from byTemplate, e.g. /blog/*, to see only its pages.
	Limit          *int64   `json:"limit,omitempty"`
}

// GetSEOChanges calls get_seo_changes: What changed on the fields search
// engines read. Compares each page's status, robots directives (meta and
// X-Robots-Tag), canonical, hreflang, title, H1 and structured data types
// between two crawls (default: the latest and the one before), counted by
// field and by URL template, with every change listed. Critical: noindex
// added, canonical changed, title dropped, 200 to 404/5xx, hreflang lost, H1
// rewritten (under half its words kept) — the changes that cost rankings
// within days, the first five of which each crawl also files as regression-*
// issues. Warning: meta description changed, Open Graph tags removed,
// structured data values changed (schema hash), a Core Web Vital p75 more than
// 20% worse in Chrome's field data, a template's Lighthouse performance down
// 10 points or more. Notice: H2 outline changed, page copy changed (body
// hash). The description, og, h2 and body fields compare only crawls that
// recorded them. fields, severity and template narrow it. Answers 'what did
// the last release break', 'did a template lose its canonicals' and 'why did
// these pages drop out'. Free.
func (c *Client) GetSEOChanges(ctx context.Context, in GetSEOChangesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_seo_changes", in)
}

// GetSEORevenueInput is the input to get_seo_revenue.
type GetSEORevenueInput struct {
	ClientID string `json:"clientId"`
	By       string `json:"by,omitempty"`
	Days     *int64 `json:"days,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetSEORevenue calls get_seo_revenue: Revenue by channel, page and query.
// Imported revenue (orders, revenue net of refunds, customers — per currency,
// never converted) by channel (organic search, AI assistants, paid, social,
// email, referral, direct, unattributed), landing page, Search Console query
// cluster, campaign, source, AI assistant or store, with how many orders were
// matched to their visit. Answers 'what is SEO worth in revenue', 'which pages
// and queries make money', 'what does ChatGPT send us in sales'. Free.
func (c *Client) GetSEORevenue(ctx context.Context, in GetSEORevenueInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_seo_revenue", in)
}

// GetSEOScorecardInput is the input to get_seo_scorecard.
type GetSEOScorecardInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
}

// GetSEOScorecard calls get_seo_scorecard: SEO scorecard with an action plan.
// One 0–100 score for the site from its latest crawl (or crawlId), as seven
// weighted categories: Technical 22 (crawl, indexing, security and URL rules),
// Content 23 (thin and duplicate content, averaged with E-E-A-T over a sample
// of top pages), On-page 20 (titles, descriptions, headings, links, Open
// Graph), Schema 10 (structured-data rules and the share of pages with valid
// markup), Performance 10 (Chrome UX Report pass 100 / fail 50, averaged with
// the audit's Lighthouse runs), AI readiness 10 (the top pages' answer
// readiness) and Images 5. Each category says how it was measured, or that it
// was not (then the others are rescaled, never scored as zero), with its worst
// findings. Also the business type detected from the homepage and URLs (SaaS,
// local, ecommerce, publisher, agency; stored on the client brief), the five
// worst critical issues, five quick wins (warnings confined to one URL
// template, so one fix clears them) and a four-phase plan: critical issues in
// week 1, warnings in weeks 2–3, content and authority in month 2, then
// monitoring. Findings accepted as intentional (accept_issue) count against
// nothing and are listed in accepted with their reason. A heuristic beside the
// audit's health score, which is unchanged. Answers 'how is the site doing
// overall' and 'what should we fix first'. Free.
func (c *Client) GetSEOScorecard(ctx context.Context, in GetSEOScorecardInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_seo_scorecard", in)
}

// GetSerpContentProfileInput is the input to get_serp_content_profile.
type GetSerpContentProfileInput struct {
	ClientID      string `json:"clientId"`
	Keyword       string `json:"keyword"`
	Depth         *int64 `json:"depth,omitempty"`         // How many organic results to read.
	LocationCode  *int64 `json:"locationCode,omitempty"`  // Google location code (2840 United States, 2826 United Kingdom, 2724 Spain). Default: the client's project.
	LanguageCode  string `json:"languageCode,omitempty"`  // Language code (en, es). Default: the client's project.
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetSerpContentProfile calls get_serp_content_profile: What the ranking pages
// have in common. Reads the top Google results for a keyword and each ranking
// page, and reports what they share: word count range, the outline themes most
// of them cover, the terms they use and how often, the questions they answer,
// tables, lists, images, video, schema types, reading ease and how recently
// they were updated — plus People Also Ask and related searches, and each
// page's own summary. Results that are not pages to write against — videos and
// Shorts, forum threads, social posts, app store listings — stay in pages with
// their reason in excluded, and are listed in excluded as {url, reason};
// targetPages counts the rest. Answers 'what does it take to rank for this'.
// The results page is a paid lookup on the client's account unless someone
// bought it in the last day (spend says which); the pages are fetched free.
func (c *Client) GetSerpContentProfile(ctx context.Context, in GetSerpContentProfileInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_serp_content_profile", in)
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
// do we own'. imagePresence covers image results across every keyword the
// client ranks for: how many ranking pages carry an image pack, how many show
// a client image, and the keywords where it ranks but has no image in the pack
// (best position first) — 'where should we be in image results'. Features are
// null for a capture whose source did not report them, which is not the same
// as a page with none.
func (c *Client) GetSerpFeatures(ctx context.Context, in GetSerpFeaturesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_serp_features", in)
}

// GetSerpHistoryInput is the input to get_serp_history.
type GetSerpHistoryInput struct {
	ClientID      string `json:"clientId"`
	Keyword       string `json:"keyword"`
	From          string `json:"from,omitempty"`
	To            string `json:"to,omitempty"`
	LocationCode  *int64 `json:"locationCode,omitempty"`  // Default: the client's project (2840 United States).
	LanguageCode  string `json:"languageCode,omitempty"`  // Default: the client's project.
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetSerpHistory calls get_serp_history: Who ranked for a keyword over time.
// Past Google results pages for any keyword (not only tracked ones): for each
// date, the top results and the features shown. Answers 'who used to rank for
// this', 'when did this competitor take over' and 'did an update reshuffle
// this results page'. A paid lookup on the client's account unless anyone read
// the same thing recently (spend says which); quote it with estimate_cost.
func (c *Client) GetSerpHistory(ctx context.Context, in GetSerpHistoryInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_serp_history", in)
}

// GetSerpIntentShiftsInput is the input to get_serp_intent_shifts.
type GetSerpIntentShiftsInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
}

// GetSerpIntentShifts calls get_serp_intent_shifts: Keywords whose results
// page changed character. Tracked keywords whose results page changed what it
// is for between its two latest captures: the mix moving (informational,
// commercial, local, plain — e.g. informational → commercial when shopping
// results arrive) or a watched feature appearing or going (AI Overview, video
// carousel, shopping, map pack, featured snippet, top stories). A page that no
// longer fits the results page's intent loses however good it is. Answers 'did
// Google change what it shows for our keywords' and 'where did AI Overviews
// appear'. Free.
func (c *Client) GetSerpIntentShifts(ctx context.Context, in GetSerpIntentShiftsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_serp_intent_shifts", in)
}

// GetSerpNewcomersInput is the input to get_serp_newcomers.
type GetSerpNewcomersInput struct {
	ClientID      string `json:"clientId"`
	Days          *int64 `json:"days,omitempty"`
	BaselineWeeks *int64 `json:"baselineWeeks,omitempty"`
}

// GetSerpNewcomers calls get_serp_newcomers: Who is new in the top 20. Domains
// that entered the top 20 for any of the client's tracked keywords in the last
// days (default 7) and were in none of their top 20s in the baselineWeeks
// before, with the pages they rank with, most keywords first. The client's own
// domain is left out. Answers 'who is new in our search results this week' and
// 'which competitor just started ranking'. Free.
func (c *Client) GetSerpNewcomers(ctx context.Context, in GetSerpNewcomersInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_serp_newcomers", in)
}

// GetSerpVolatilityInput is the input to get_serp_volatility.
type GetSerpVolatilityInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId,omitempty"` // Default: every tracker of the client.
	Days      *int64 `json:"days,omitempty"`
}

// GetSerpVolatility calls get_serp_volatility: How much the results are
// moving. How much page one moved between readings for each tracked keyword (0
// still, ~10 turned over), a daily index for the client's keywords beside an
// industry-wide one from every keyword in the archive, and Google's updates in
// the window. Answers 'is it us or is it Google' when rankings move. Free.
func (c *Client) GetSerpVolatility(ctx context.Context, in GetSerpVolatilityInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_serp_volatility", in)
}

// GetShareOfSearchInput is the input to get_share_of_search.
type GetShareOfSearchInput struct {
	ClientID string   `json:"clientId"`
	Topic    string   `json:"topic,omitempty"` // Keywords containing this text, and AI questions about it.
	Keywords []string `json:"keywords,omitempty"`
	Days     *int64   `json:"days,omitempty"`
}

// GetShareOfSearch calls get_share_of_search: The brand's share of each
// engine. For the client's tracked keywords (narrowed to a topic, or to given
// keywords): one number per engine and the brand's share of it — Google, Bing
// and YouTube (estimated click share from the latest rank captures, beside the
// total search volume), Google's AI Overviews (share of the citations they
// showed), AI Mode and the other AI assistants (share of brand mentions in the
// answers Seolyticize sampled). Read only from stored data, so free; an engine
// with no trackers, captures or sampled answers is listed under noData with
// why. Answers 'where are we visible beyond Google' and 'what share of AI
// answers name us'.
func (c *Client) GetShareOfSearch(ctx context.Context, in GetShareOfSearchInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_share_of_search", in)
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

// GetShoppingKeywordGapInput is the input to get_shopping_keyword_gap.
type GetShoppingKeywordGapInput struct {
	ClientID   string   `json:"clientId"`
	Keywords   []string `json:"keywords,omitempty"`
	Sellers    []string `json:"sellers,omitempty"`
	OrganicTop *int64   `json:"organicTop,omitempty"`
	Days       *int64   `json:"days,omitempty"`
	Limit      *int64   `json:"limit,omitempty"`
}

// GetShoppingKeywordGap calls get_shopping_keyword_gap: Organic rankings
// against Google Shopping presence. The client's ranked organic keywords (as
// last synced, top organicTop positions) and any keywords given, against
// whether its products showed in the Google Shopping pages captured by
// get_shopping_serp in the last days, in four groups, each with what to do:
// both (defend), organic_only (ranks but no product listed — list it in
// Merchant Center), shopping_only (a product listed but no organic ranking —
// build the page) and neither. Organic keywords never captured in Shopping are
// listed apart as unchecked, to queue with check_shopping_serp (paid). Matched
// on the seller names given (default the client's name) and the client's
// domain. Answers 'where do we rank but not sell in Shopping' and 'which
// Shopping keywords need a page'. Free.
func (c *Client) GetShoppingKeywordGap(ctx context.Context, in GetShoppingKeywordGapInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_shopping_keyword_gap", in)
}

// GetShoppingSerpInput is the input to get_shopping_serp.
type GetShoppingSerpInput struct {
	ClientID string           `json:"clientId"`
	Tasks    []map[string]any `json:"tasks"`
	Sellers  []string         `json:"sellers,omitempty"`
	Days     *int64           `json:"days,omitempty"`
}

// GetShoppingSerp calls get_shopping_serp: Google Shopping positions, now and
// over time. Collects the Shopping pages queued by check_shopping_serp, keeps
// them, and says where the client's products sit — matched on the seller names
// given (default the client's name) and the client's domain — with the top
// listings beside it and the market on each page (price min, max, median, P25
// and P75 with outliers over 2σ from the median, the top 10 sellers' share of
// listings, the share with free shipping and the rating distribution), and the
// client's position history for those keywords from every capture kept. A task
// still queued is listed as pending; ask again shortly. Answers 'where do we
// rank in Google Shopping' and 'who outranks us there'. Free: collecting a
// queued task costs nothing.
func (c *Client) GetShoppingSerp(ctx context.Context, in GetShoppingSerpInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_shopping_serp", in)
}

// GetSiteConventionsInput is the input to get_site_conventions.
type GetSiteConventionsInput struct {
	ClientID string `json:"clientId"`
	URL      string `json:"url,omitempty"`
	Samples  *int64 `json:"samples,omitempty"`
}

// GetSiteConventions calls get_site_conventions: Site-wide conventions: HSTS
// preload, slashes, rendering, share image. Four site-wide checks from a
// handful of live requests: whether the base domain qualifies for the HSTS
// preload list (max-age of a year or more, includeSubDomains, preload, http://
// redirecting to https://; list membership itself is not checked); the
// trailing-slash convention, from up to samples crawled paths each asked for
// with and without the slash (one form redirecting to the other is consistent;
// both answering 200 is two addresses per page unless both name one canonical;
// some paths redirecting each way is mixed); dynamic rendering, suspected when
// Googlebot is served prerendered content and a browser an app shell with a
// fraction of the words (Google calls it a workaround, not a recommended
// solution); and the page's og:image, which should answer 200 and be at least
// 1200×630. url picks the page for the last two (default: the homepage).
// Answers 'are we ready for HSTS preload', 'do we have trailing-slash
// duplicates' and 'is our share image big enough'. Free.
func (c *Client) GetSiteConventions(ctx context.Context, in GetSiteConventionsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_site_conventions", in)
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

// GetSnippetTargetsInput is the input to get_snippet_targets.
type GetSnippetTargetsInput struct {
	ClientID      string `json:"clientId"`
	Limit         *int64 `json:"limit,omitempty"`
	Draft         *bool  `json:"draft,omitempty"`
	ReadMissing   *bool  `json:"readMissing,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetSnippetTargets calls get_snippet_targets: Featured snippets within reach.
// Queries where Google shows a featured snippet or People Also Ask and the
// client ranks 2–10 (Search Console, last 28 days), from the stored results
// pages (readMissing: true buys the ones not stored): who holds the snippet,
// the format to win it — a paragraph of about 45 words, a list or a table —
// and, with draft: true, a draft in that format from one metered Claude call
// for the batch. Answers 'which snippets can we steal' and 'write the answer
// box for this query'. Paid only for drafts and missing results pages; quote
// it with estimate_cost.
func (c *Client) GetSnippetTargets(ctx context.Context, in GetSnippetTargetsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_snippet_targets", in)
}

// GetSubdomainTrafficInput is the input to get_subdomain_traffic.
type GetSubdomainTrafficInput struct {
	ClientID      string `json:"clientId"` // The client whose account and monthly cap pay for the lookup
	Domain        string `json:"domain"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// GetSxoGapInput is the input to get_sxo_gap.
type GetSxoGapInput struct {
	ClientID      string `json:"clientId"`
	URL           string `json:"url"`
	Keyword       string `json:"keyword,omitempty"`
	Personas      *bool  `json:"personas,omitempty"`
	Wireframe     *bool  `json:"wireframe,omitempty"`
	Deep          *bool  `json:"deep,omitempty"` // Also render the page on a phone and score its E-E-A-T, and read UX and authority from those (free, slower).
	LocationCode  *int64 `json:"locationCode,omitempty"`
	LanguageCode  string `json:"languageCode,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetSxoGap calls get_sxo_gap: Is this the kind of page Google ranks here.
// Classifies the client's page and the top 10 results for its keyword into
// claude-seo's eight types (tool, local, comparison, product, landing,
// service, hybrid, else blog, in that order of precedence) and says how
// strongly Google agrees on one type (over 60% strong, 40–60% mixed, under 40%
// fragmented) and how bad the page's mismatch is, from an explicit page ×
// results table (a blog where products, tools, local pages or landing pages
// rank is critical; a blog against comparisons high; a landing page where
// tools rank high), and which of the expected type's required elements the
// page has. A 0–100 gap score against the ranking pages: page type, depth, UX,
// schema, media and authority 15 each, freshness 10 (null parts drop out).
// keyword defaults to what the title and H1 share, else the page's top Search
// Console query. wireframe: true adds the page's sections now, the target
// outline and the target layout for the page type the results want (landing,
// blog, product, comparison, service, local, tool or hybrid; free); deep: true
// reads UX from the page rendered on a phone (H1 and CTA above the fold, tap
// targets, font size, overlaps, overlays) and authority from its E-E-A-T score
// beside domain authority, and rescores (free, slower); personas: true adds
// 4–7 searchers from People Also Ask, ads and related searches, each a full
// card (role, goal, emotional state, journey stage, key questions, intent)
// scored on relevance, clarity, trust and action (25 each) and rated 80
// excellent, 60 good, 40 needs work, below a critical mismatch, a persona
// score weighted by each intent's share of the results, and 3–5 user stories
// ('As a …, I want to …, because …, but I'm blocked by …') with their
// evidence, from one metered Claude call. A heuristic, and it says so. Answers
// 'why doesn't this page rank' and 'what kind of page does Google want for
// this'. The results page is a paid lookup unless bought in the last day;
// quote with estimate_cost.
func (c *Client) GetSxoGap(ctx context.Context, in GetSxoGapInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_sxo_gap", in)
}

// GetTopPagesInput is the input to get_top_pages.
type GetTopPagesInput struct {
	ClientID      string `json:"clientId"`
	Target        string `json:"target,omitempty"` // A domain (competitor.com), a subfolder (competitor.com/blog/) or a page URL. Default: the client's own domain.
	Scope         string `json:"scope,omitempty"`  // auto reads a bare domain as the domain, a path ending in / as a subfolder, anything else as one URL.
	Limit         *int64 `json:"limit,omitempty"`
	LocationCode  *int64 `json:"locationCode,omitempty"`  // Default: the client's project (2840 United States).
	LanguageCode  string `json:"languageCode,omitempty"`  // Default: the client's project.
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetTopPages calls get_top_pages: A domain's pages by traffic. A domain's (or
// subfolder's) pages ranked by estimated organic traffic, with how many
// keywords each ranks for, how many in the top 3 and top 10, keywords won and
// lost, and traffic value. Answers 'which pages bring them their traffic' and
// 'what content works in this niche'. A paid lookup on the client's account
// unless anyone read the same thing recently (spend says which); quote it with
// estimate_cost.
func (c *Client) GetTopPages(ctx context.Context, in GetTopPagesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_top_pages", in)
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

// GetTopicMapInput is the input to get_topic_map.
type GetTopicMapInput struct {
	ClientID      string   `json:"clientId"`
	Seed          string   `json:"seed,omitempty"`
	Days          *int64   `json:"days,omitempty"`
	ExtraKeywords []string `json:"extraKeywords,omitempty"`
	Method        string   `json:"method,omitempty"`
	Limit         *int64   `json:"limit,omitempty"` // How many Search Console queries to map, most impressions first.
	MinShared     *int64   `json:"minShared,omitempty"`
	LocationCode  *int64   `json:"locationCode,omitempty"`  // Default: the client's project.
	LanguageCode  string   `json:"languageCode,omitempty"`  // Default: the client's project.
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetTopicMap calls get_topic_map: Map topics to pages. The client's topics
// from Search Console (non-branded queries, optionally containing a seed) plus
// any extra keywords, grouped into clusters and pillars, each mapped to the
// pages Google shows for it: no_page (write it), one_page, or several_pages
// (impressions split between pages — merge or fix internal links). Answers
// 'what should we write next', 'which topics have no page' and 'where are our
// pages competing'. method serp_overlap clusters by results-page overlap (one
// paid lookup per keyword unless bought in the last day); search_console
// clusters by the page each query lands on, free.
func (c *Client) GetTopicMap(ctx context.Context, in GetTopicMapInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_topic_map", in)
}

// GetTopicalAuthorityInput is the input to get_topical_authority.
type GetTopicalAuthorityInput struct {
	ClientID      string   `json:"clientId"`
	Days          *int64   `json:"days,omitempty"`
	Seed          string   `json:"seed,omitempty"`
	ExtraKeywords []string `json:"extraKeywords,omitempty"`
	Limit         *int64   `json:"limit,omitempty"`
}

// GetTopicalAuthority calls get_topical_authority: How strong the site is on
// each topic. A 0–100 topical authority score per topic (the topic map's
// pillars, from Search Console): coverage (the share of the topic's query
// clusters with a page ranking in the top 20) × rankings (impression-weighted)
// × internal links (links into each covering page, from the latest crawl) —
// each part given — and the next three pages that would raise it most: write
// (a cluster nothing covers), improve (a page beyond the first page) or link
// (a weakly linked page). extraKeywords adds topics the site does not rank for
// yet. Answers 'where are we an authority' and 'what should we write next to
// own this topic'. Free.
func (c *Client) GetTopicalAuthority(ctx context.Context, in GetTopicalAuthorityInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_topical_authority", in)
}

// GetToxicLinkRiskInput is the input to get_toxic_link_risk.
type GetToxicLinkRiskInput struct {
	ClientID        string         `json:"clientId"`
	MoneyTerms      []string       `json:"moneyTerms,omitempty"`
	SpamScores      map[string]any `json:"spamScores,omitempty"`      // Domain → DataForSEO spam score, from get_backlink_spam_scores.
	FetchSpamScores *bool          `json:"fetchSpamScores,omitempty"` // True to buy the vendor's spam score for every linking site inside this call (one paid request of up to 1,000 sites, the price of get_backlink_spam_scores; sites
	MinVerdict      string         `json:"minVerdict,omitempty"`
	Limit           *int64         `json:"limit,omitempty"`
	MaxCostMicros   *int64         `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun          *bool          `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetToxicLinkRisk calls get_toxic_link_risk: Which linking sites are a risk,
// and why. Judges each site linking to the client from the evidence a person
// can check: the vendor's spam score (fetchSpamScores: true buys it for every
// linking site inside this call — one paid DataForSEO request of up to 1,000
// sites, the price of get_backlink_spam_scores, with sites scored in the last
// 30 days reused free; or pass spamScores you already have), sitewide links
// (the link on most of the pages we hold of the site: a footer or blogroll),
// link-seller neighbourhoods (a site linking out to hundreds of domains from
// few pages), commercial exact-match anchors (pass moneyTerms; brand terms are
// excluded), the vendor links' own authority and outbound signals, and
// claude-seo's thirty toxic patterns where the data allows (a page with
// 10,000+ or 50+ outbound links, a domain registered under 30 days ago with
// 100+ outbound links or re-registered after it lapsed — from RDAP for the
// sites linking out most — pharma or casino injections, sitewide links and
// exact-match anchors from unrelated sites, foreign-language and parked sites,
// article directories, forum profiles, press release wires, coupon sites,
// comment links, unrelated niches when moneyTerms say what the client does),
// each pattern reported flagged, clear or unmeasured with why. disavow says
// whether a file is worth making (toxic over 10% of the sites: yes; under 2%:
// no). A verdict (healthy, review, toxic) and every reason with its detail;
// whether the last saved disavow file already names the site, and the toxic
// sites that appeared since it. Answers 'which of our links are toxic and why'
// and 'what should go in the disavow file now'. 'review' means look, not
// disavow. Free without fetchSpamScores; with it, quote it with estimate_cost
// or dryRun first.
func (c *Client) GetToxicLinkRisk(ctx context.Context, in GetToxicLinkRiskInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_toxic_link_risk", in)
}

// GetTrackerAIModeInput is the input to get_tracker_ai_mode.
type GetTrackerAIModeInput struct {
	ClientID      string `json:"clientId"`
	TrackerID     string `json:"trackerId"`
	Limit         *int64 `json:"limit,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetTrackerAIMode calls get_tracker_ai_mode: AI Mode beside each ranking.
// Reads Google AI Mode for a tracker's keywords (up to limit) and puts beside
// each latest position whether AI Mode cites the client and which domains it
// cites — the gap between ranking and being the answer. One paid AI Mode read
// per keyword: quote it with estimate_cost first. Answers 'we rank for these,
// but does AI Mode cite us'.
func (c *Client) GetTrackerAIMode(ctx context.Context, in GetTrackerAIModeInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_tracker_ai_mode", in)
}

// GetTranslationGapInput is the input to get_translation_gap.
type GetTranslationGapInput struct {
	ClientID string `json:"clientId"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	Examples *int64 `json:"examples,omitempty"`
}

// GetTranslationGap calls get_translation_gap: Pages missing in a language.
// Which pages exist in one language and not another: each indexable page's
// language (its hreflang, its /es/-style path prefix or its html lang),
// versions matched by hreflang or by the same path without the prefix, and for
// each language its coverage of the main language with the pages it lacks.
// Answers 'which pages have no Spanish version' and 'how complete is our /es
// site'. Free.
func (c *Client) GetTranslationGap(ctx context.Context, in GetTranslationGapInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_translation_gap", in)
}

// GetTrustSignalsInput is the input to get_trust_signals.
type GetTrustSignalsInput struct {
	ClientID string `json:"clientId"`
}

// GetTrustSignals calls get_trust_signals: Who stands behind the site. Checks
// the signals of who is behind a site that Google's quality raters and AI
// assistants weigh (E-E-A-T): an about page, a contact page with a way to
// reach a person, a privacy policy and terms, Organization structured data on
// the homepage, named authors and publish dates on articles, and HTTPS. Each
// pass, fail or unknown, with what was found. Answers 'does our site show who
// we are' and 'what trust signals are missing'. Free.
func (c *Client) GetTrustSignals(ctx context.Context, in GetTrustSignalsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_trust_signals", in)
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

// GetVendorReconciliationInput is the input to get_vendor_reconciliation.
type GetVendorReconciliationInput struct {
}

// GetVendorReconciliation calls get_vendor_reconciliation: Reconcile vendor
// spend against the ledger. Sets the agency's DataForSEO account (deposited,
// balance, spent; today's spend per API) beside what Seolyticize's cost ledger
// recorded on the same account, with the difference, the ledger by operation,
// and what can explain a gap: spend before the ledger began, the same login
// used outside Seolyticize, other agencies sharing it. Answers 'why did our
// DataForSEO balance fall more than get_costs shows'. agreesWithinCents is
// true when the two agree to a cent. Reads the account from DataForSEO's free
// user_data endpoint. On the platform's shared key, only a platform admin may
// reconcile. Amounts are US dollars. Free.
func (c *Client) GetVendorReconciliation(ctx context.Context, in GetVendorReconciliationInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_vendor_reconciliation", in)
}

// GetVitalsHistoryInput is the input to get_vitals_history.
type GetVitalsHistoryInput struct {
	ClientID   string `json:"clientId"`
	URL        string `json:"url,omitempty"`
	FormFactor string `json:"formFactor,omitempty"`
}

// GetVitalsHistory calls get_vitals_history: Core Web Vitals over 25 weeks.
// Core Web Vitals as Chrome measured them, week by week for the last 25 weeks,
// from the Chrome UX Report History API: p75 LCP, INP, CLS, TTFB and FCP (good
// at 1.8 s or less, poor over 3 s) for the client's origin, or one URL (url)
// when it has enough traffic, by form factor (phone, desktop or all).
// Regressions name a metric that got worse by more than 10% over the last four
// weeks or crossed into a worse rating; trends compare the first four weeks
// with the last four (within ±5% is stable). lcpBreakdown splits the newest
// image LCP into Chrome's four subparts (time to first byte, resource load
// delay, load duration, render delay) and names any at 40% of p75 LCP or more,
// with what to fix. Also the audits' Lighthouse runs per template, for the lab
// side. Without a Chrome UX Report key it says so and shows what is stored;
// nothing is estimated. Answers 'did our Core Web Vitals get worse', 'when did
// LCP regress' and 'are we passing on mobile'. Free.
func (c *Client) GetVitalsHistory(ctx context.Context, in GetVitalsHistoryInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_vitals_history", in)
}

// GetVolumeCalibrationInput is the input to get_volume_calibration.
type GetVolumeCalibrationInput struct {
	ClientID string `json:"clientId"`
}

// GetVolumeCalibration calls get_volume_calibration: How far vendor volumes
// are off for this site. The correction factor between the vendor's search
// volumes and what Search Console counted, per market: for queries where the
// site held positions 1–3 (where impressions are close to searches), the
// median of impressions over volume, with how many terms it rests on and their
// spread. Keyword tools add calibratedVolume next to volume where a factor
// exists. Answers 'can we trust these search volumes' and 'why does Search
// Console show more searches than the keyword tool'. Free.
func (c *Client) GetVolumeCalibration(ctx context.Context, in GetVolumeCalibrationInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_volume_calibration", in)
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
// (ISO codes; null means unknown), devices, browsers, os, custom events (event
// count), channels (visits by channel: AI, Organic Search, Paid Search,
// Social, Email, Referral, Campaign, Direct), ai_sources (visits from each AI
// assistant: ChatGPT, Perplexity, Gemini, Copilot, Claude, Google AI Mode and
// others) or ai_landing_pages (the pages AI assistants send visits to, detail
// naming the assistant). Each row has key, visitors and count; counts says
// what count counts, and totalKeys how many rows exist before the limit.
// Answers 'which pages and sources drive visits'. Refuses with a sentence when
// web analytics has not been set up for the client. Free.
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

// GetYandexDiagnosticsInput is the input to get_yandex_diagnostics.
type GetYandexDiagnosticsInput struct {
	ClientID string `json:"clientId"`
}

// GetYandexDiagnostics calls get_yandex_diagnostics: What Yandex says is wrong
// with the site. Yandex Webmaster's site diagnostics: each check (robots.txt
// blocking the site, DNS or SSL errors, a slow server, missing or broken
// sitemaps, duplicate pages, missing titles, not mobile friendly...) with its
// severity and whether the problem is present now, worst first. Answers 'does
// Yandex see anything broken'. Free.
func (c *Client) GetYandexDiagnostics(ctx context.Context, in GetYandexDiagnosticsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_yandex_diagnostics", in)
}

// GetYandexExternalLinksInput is the input to get_yandex_external_links.
type GetYandexExternalLinksInput struct {
	ClientID string `json:"clientId"`
	Offset   *int64 `json:"offset,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetYandexExternalLinks calls get_yandex_external_links: Links Yandex knows
// point at the site. Yandex's count of external links to the client's host
// over time, and a page of the links themselves (source page, the page of ours
// it points to, when Yandex found it). Answers 'is our link count growing in
// Yandex' and 'who links to us, per Yandex'. Free.
func (c *Client) GetYandexExternalLinks(ctx context.Context, in GetYandexExternalLinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_yandex_external_links", in)
}

// GetYandexImportantPagesInput is the input to get_yandex_important_pages.
type GetYandexImportantPagesInput struct {
	ClientID string `json:"clientId"`
}

// GetYandexImportantPages calls get_yandex_important_pages: The pages Yandex
// is watching for the site. Yandex Webmaster's important-pages monitoring:
// each page the site owner chose to watch, with its last crawl status and HTTP
// code, whether it is searchable, its title and description in Yandex, why it
// is excluded if it is, and what changed since the previous check. Answers
// 'are our key pages still in Yandex'. Pages are chosen in Yandex Webmaster
// itself; the API only reads them. Free.
func (c *Client) GetYandexImportantPages(ctx context.Context, in GetYandexImportantPagesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_yandex_important_pages", in)
}

// GetYandexIndexingInput is the input to get_yandex_indexing.
type GetYandexIndexingInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from"`
	To       string `json:"to"`
	Offset   *int64 `json:"offset,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// GetYandexIndexing calls get_yandex_indexing: Which pages Yandex has in
// search. Yandex Webmaster's indexing views for the client's verified host:
// pages in search by day and a sample of them, pages that appeared in or were
// removed from search by day and a sample with Yandex's reason for each
// removal (duplicate, not canonical, noindex, HTTP error, low quality...)
// counted by reason, and a sample of what the robot fetched with the status it
// got. Answers 'which pages has Yandex dropped, and why'. Samples are paged
// (limit up to 100). Free.
func (c *Client) GetYandexIndexing(ctx context.Context, in GetYandexIndexingInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_yandex_indexing", in)
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

// GetYandexQueryHistoryInput is the input to get_yandex_query_history.
type GetYandexQueryHistoryInput struct {
	ClientID string `json:"clientId"`
	From     string `json:"from"`
	To       string `json:"to"`
	Device   string `json:"device,omitempty"`
	QueryID  string `json:"queryId,omitempty"`
}

// GetYandexQueryHistory calls get_yandex_query_history: Yandex search traffic
// by day. Yandex Webmaster's search query history: shows, clicks and average
// show and click position by day for the whole site, or for one query
// (queryId, from get_yandex_performance's queries), optionally for one device.
// Answers 'how has our Yandex traffic moved' and 'how is this query trending
// in Yandex'. Free.
func (c *Client) GetYandexQueryHistory(ctx context.Context, in GetYandexQueryHistoryInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_yandex_query_history", in)
}

// GetYoutubeChannelRankingsInput is the input to get_youtube_channel_rankings.
type GetYoutubeChannelRankingsInput struct {
	ClientID      string   `json:"clientId"`
	Channel       string   `json:"channel"` // A channel id (UC…) or the channel's name.
	Keywords      []string `json:"keywords"`
	Depth         *int64   `json:"depth,omitempty"`
	LocationCode  *int64   `json:"locationCode,omitempty"`
	LanguageCode  string   `json:"languageCode,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// GetYoutubeChannelRankings calls get_youtube_channel_rankings: Which of a
// channel's videos rank, for which phrases. For up to 20 phrases, reads
// YouTube's results and lists every video from the channel (matched by channel
// id, or by name) with its position and views, and the phrases where the
// channel does not appear. Works for the client's channel or a competitor's.
// For tracking over time, a rank tracker on the youtube engine does the same
// daily. Answers 'where do our videos rank' and 'which of their videos rank
// for our topics'. One results page per phrase: A paid lookup on the client's
// account unless anyone read the same thing today (spend says which); quote it
// with estimate_cost.
func (c *Client) GetYoutubeChannelRankings(ctx context.Context, in GetYoutubeChannelRankingsInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_youtube_channel_rankings", in)
}

// GetZeroClickEstimatesInput is the input to get_zero_click_estimates.
type GetZeroClickEstimatesInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"` // Days of Search Console history, ending on the newest day Google has finished.
	Limit    *int64 `json:"limit,omitempty"`
}

// GetZeroClickEstimates calls get_zero_click_estimates: How many searches end
// without a click. For each Search Console query that a tracked keyword's
// latest capture describes, the estimated share of searches that end with no
// organic click, from the features on its results page (AI Overview, featured
// snippet, map pack, shopping, knowledge panel, video…) — measured on this
// site's own click-through curve where it can be, a published share otherwise
// (source says which) — and how many of the query's searches that is. Answers
// 'how much demand is the AI Overview eating' and 'which keywords are worth
// less than their volume says'. Free.
func (c *Client) GetZeroClickEstimates(ctx context.Context, in GetZeroClickEstimatesInput) (json.RawMessage, error) {
	return c.Call(ctx, "get_zero_click_estimates", in)
}

// ImportRevenueNowInput is the input to import_revenue_now.
type ImportRevenueNowInput struct {
	ClientID string `json:"clientId"`
	Provider string `json:"provider,omitempty"`
}

// ImportRevenueNow calls import_revenue_now: Import orders now. Imports what
// changed in the client's connected stores since the last import (the nightly
// import does this on its own), attributing each order to the visit it came
// from. Answers 'pull in today's orders'. Free.
func (c *Client) ImportRevenueNow(ctx context.Context, in ImportRevenueNowInput) (json.RawMessage, error) {
	return c.Call(ctx, "import_revenue_now", in)
}

// InspectBingURLInput is the input to inspect_bing_url.
type InspectBingURLInput struct {
	ClientID string `json:"clientId"`
	URL      string `json:"url"`
}

// InspectBingURL calls inspect_bing_url: Ask Bing what it holds for one page.
// Bing's index details for one address on the client's site (GetUrlInfo): when
// Bing discovered and last crawled it, the HTTP status and size it got, links
// pointing at it and, for a directory, how many pages sit under it — plus its
// Bing clicks and impressions. Answers 'is this page in Bing, and when did
// Bing last see it'. Free; no daily allowance.
func (c *Client) InspectBingURL(ctx context.Context, in InspectBingURLInput) (json.RawMessage, error) {
	return c.Call(ctx, "inspect_bing_url", in)
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

// InviteTeamMemberInput is the input to invite_team_member.
type InviteTeamMemberInput struct {
	Email     string   `json:"email"`
	Role      string   `json:"role"`
	ClientIDs []string `json:"clientIds,omitempty"`
}

// InviteTeamMember calls invite_team_member: Invite someone to the team by
// email. Emails an invitation to join the agency with a role — admin, manager,
// analyst or client_viewer — optionally limited to some clients. Unlike
// add_team_member the person need not have an account yet: they sign up or
// sign in, as the invited address, and the link makes them a member. The link
// lasts 7 days and works once. Answers 'invite sam@agency.com as an analyst on
// Acme'. Returns the invite, whether the email was sent, and the link itself
// for passing on by hand. Refused for an existing member, an address with a
// live invite (use resend_team_invite), or a client outside the caller's
// reach.
func (c *Client) InviteTeamMember(ctx context.Context, in InviteTeamMemberInput) (json.RawMessage, error) {
	return c.Call(ctx, "invite_team_member", in)
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

// IssueServerCollectKeyInput is the input to issue_server_collect_key.
type IssueServerCollectKeyInput struct {
	ClientID string `json:"clientId"`
}

// IssueServerCollectKey calls issue_server_collect_key: Key for server-side
// collection. Issues the secret key a server-side collector
// (@seolyticize/next/server, the Python and PHP collectors, or a POST to
// /api/collect with Authorization: Bearer) reports with, so it may send each
// visitor's own address and user agent (body fields ip and ua; howToUse lists
// the rest) — which is what lets it count people, and the AI assistants
// (ChatGPT-User, Claude-User, Perplexity-User) that never run the browser tag.
// Shown once; issuing another revokes the last. Answers 'track visits from our
// server' and 'count AI agents reading our pages'. Free.
func (c *Client) IssueServerCollectKey(ctx context.Context, in IssueServerCollectKeyInput) (json.RawMessage, error) {
	return c.Call(ctx, "issue_server_collect_key", in)
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
	DryRun        *bool          `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// ListAcceptedIssuesInput is the input to list_accepted_issues.
type ListAcceptedIssuesInput struct {
	ClientID      string `json:"clientId"`
	IncludeLapsed *bool  `json:"includeLapsed,omitempty"`
}

// ListAcceptedIssues calls list_accepted_issues: List the issues accepted as
// intentional. The client's accepted audit findings: rule, template or
// addresses, the reason, until and who accepted it. includeLapsed adds the
// ones whose until has passed, which count again. Answers 'what have we
// decided to ignore, and why'. Free.
func (c *Client) ListAcceptedIssues(ctx context.Context, in ListAcceptedIssuesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_accepted_issues", in)
}

// ListAlertRulesInput is the input to list_alert_rules.
type ListAlertRulesInput struct {
	ClientID string `json:"clientId"`
}

// ListAlertRules calls list_alert_rules: List alert rules. What a client's
// alerts watch for, at what threshold, and whether each is paused and can be
// delivered (hasDeliveryTarget false: it fires into nothing; undeliverable
// lists those rules with the fix). Lists the rules; list_recent_alerts lists
// what they fired. add_default_alerts adds the starting four.
func (c *Client) ListAlertRules(ctx context.Context, in ListAlertRulesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_alert_rules", in)
}

// ListAnnotationsInput is the input to list_annotations.
type ListAnnotationsInput struct {
	ClientID      string `json:"clientId"`
	From          string `json:"from,omitempty"` // Default: a year before to.
	To            string `json:"to,omitempty"`   // Default: today.
	IncludeGoogle *bool  `json:"includeGoogle,omitempty"`
}

// ListAnnotations calls list_annotations: List the client's timeline. The
// client's timeline notes between two days, oldest first, with Google's
// algorithm updates when includeGoogle is true. Answers 'what changed around
// then'. Free.
func (c *Client) ListAnnotations(ctx context.Context, in ListAnnotationsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_annotations", in)
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

// ListClientFactsInput is the input to list_client_facts.
type ListClientFactsInput struct {
	ClientID string `json:"clientId"`
}

// ListClientFacts calls list_client_facts: List the client's canonical facts.
// The client's registered facts by category: key, value, terms and source.
// Answers 'what do we hold as true about this client'. Free.
func (c *Client) ListClientFacts(ctx context.Context, in ListClientFactsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_client_facts", in)
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

// ListContentPlanInput is the input to list_content_plan.
type ListContentPlanInput struct {
	ClientID string   `json:"clientId"`
	From     string   `json:"from,omitempty"`
	To       string   `json:"to,omitempty"`
	Status   []string `json:"status,omitempty"`
	PlanID   string   `json:"planId,omitempty"`
}

// ListContentPlan calls list_content_plan: List the editorial calendar. The
// client's planned content, month by month: keyword, kind (new or refresh),
// target URL, status, brief arguments, internal links and estimated value.
// Filter by month range, status or one plan. Answers 'what is planned for next
// month'. Free.
func (c *Client) ListContentPlan(ctx context.Context, in ListContentPlanInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_content_plan", in)
}

// ListConversionGoalsInput is the input to list_conversion_goals.
type ListConversionGoalsInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
}

// ListConversionGoals calls list_conversion_goals: The events a site sends.
// The custom events the client's site sends to the cookieless analytics,
// most-sent first, with how many carry a value — the goals get_conversions can
// count. Answers 'what are we tracking'. Says how to send one when there are
// none. Free.
func (c *Client) ListConversionGoals(ctx context.Context, in ListConversionGoalsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_conversion_goals", in)
}

// ListCrawlIssuesInput is the input to list_crawl_issues.
type ListCrawlIssuesInput struct {
	ClientID      string `json:"clientId"`
	CrawlID       string `json:"crawlId,omitempty"`
	IndexableOnly *bool  `json:"indexableOnly,omitempty"`
}

// ListCrawlIssues calls list_crawl_issues: List site audit issues. Technical
// issues one crawl found, by rule: its readable label, severity, category, how
// many pages it affects, and howToFix — concrete steps to fix it and check the
// fix. Lists what to fix. Use crawlId 'latest' for the newest completed crawl.
// indexableOnly (default true) leaves out findings on pages search engines
// will not index as themselves — noindexed, not answering 200, or
// canonicalised elsewhere — except the findings that explain why; pass false
// to see everything.
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

// ListDataSyncsInput is the input to list_data_syncs.
type ListDataSyncsInput struct {
	ClientID string `json:"clientId"`
}

// ListDataSyncs calls list_data_syncs: List data syncs. Lists every scheduled
// data sync for the client: what it queries, where it writes, how often, when
// it last ran, how many rows and any error. Free.
func (c *Client) ListDataSyncs(ctx context.Context, in ListDataSyncsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_data_syncs", in)
}

// ListDigestSchedulesInput is the input to list_digest_schedules.
type ListDigestSchedulesInput struct {
	ClientID string `json:"clientId"`
}

// ListDigestSchedules calls list_digest_schedules: List digest schedules.
// Lists where and when the client's weekly digest is posted, and when it last
// went. Free.
func (c *Client) ListDigestSchedules(ctx context.Context, in ListDigestSchedulesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_digest_schedules", in)
}

// ListFunnelsInput is the input to list_funnels.
type ListFunnelsInput struct {
	ClientID string `json:"clientId"`
}

// ListFunnels calls list_funnels: List funnels. Lists the funnels defined for
// a client, with their steps. Free.
func (c *Client) ListFunnels(ctx context.Context, in ListFunnelsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_funnels", in)
}

// ListGoalsInput is the input to list_goals.
type ListGoalsInput struct {
	ClientID string `json:"clientId"`
}

// ListGoals calls list_goals: List conversion goals. The goals defined for a
// client's cookieless analytics, with what each matches and its value. Lists
// what get_goal_report counts; list_conversion_goals shows the custom events
// being sent. Free.
func (c *Client) ListGoals(ctx context.Context, in ListGoalsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_goals", in)
}

// ListIntegrationsInput is the input to list_integrations.
type ListIntegrationsInput struct {
	ClientID string `json:"clientId"`
}

// ListIntegrations calls list_integrations: A client's connected tools. Lists
// every outside service connected for the client — stores, sheets, Airtable,
// BigQuery, chats, trackers — with how it was connected, the key's last
// characters, when it last worked and its last error. Never the secret. Free.
func (c *Client) ListIntegrations(ctx context.Context, in ListIntegrationsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_integrations", in)
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
	ClientID      string `json:"clientId"`
	RuleID        string `json:"ruleId"`
	CrawlID       string `json:"crawlId,omitempty"`
	Limit         *int64 `json:"limit,omitempty"`
	Offset        *int64 `json:"offset,omitempty"`
	IndexableOnly *bool  `json:"indexableOnly,omitempty"`
}

// ListIssuePages calls list_issue_pages: List the pages an issue is on. Lists
// the pages where one crawl issue was found, with the message for each page,
// so they can be fixed. Answers 'which pages have missing titles' or 'where
// are the broken links'. ruleId comes from get_audit_summary or
// list_crawl_issues; crawlId defaults to the latest completed crawl. Paged:
// limit (up to 500) and offset; total is the whole count. indexableOnly
// (default true) leaves out findings on pages search engines will not index as
// themselves — noindexed, not answering 200, or canonicalised elsewhere —
// except the findings that explain why; pass false to see everything. Each
// page carries its path owner (owner, from set_path_owners) and, when
// accept_issue covers it, accepted with the reason. Free.
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

// ListKeywordOwnersInput is the input to list_keyword_owners.
type ListKeywordOwnersInput struct {
	ClientID string `json:"clientId"`
	URL      string `json:"url,omitempty"`
}

// ListKeywordOwners calls list_keyword_owners: Which page owns which keyword.
// The client's keyword-to-page registry: each keyword (with its cluster) and
// the page that owns it, optionally for one page. Answers 'which keywords is
// /blog/x meant to rank for' and 'who owns this query'. Free.
func (c *Client) ListKeywordOwners(ctx context.Context, in ListKeywordOwnersInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_keyword_owners", in)
}

// ListLeadWidgetsInput is the input to list_lead_widgets.
type ListLeadWidgetsInput struct {
}

// ListLeadWidgets calls list_lead_widgets: List lead-gen widgets. The agency's
// audit forms, each with its embed code, allowed sites, who is notified and
// whether it is switched on. Answers 'which audit forms do we have on our
// sites'. Free.
func (c *Client) ListLeadWidgets(ctx context.Context, in ListLeadWidgetsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_lead_widgets", in)
}

// ListLeadsInput is the input to list_leads.
type ListLeadsInput struct {
	WidgetID string `json:"widgetId,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
}

// ListLeads calls list_leads: List leads from audit widgets. People who asked
// for an audit through the agency's widgets, newest first: email, the page
// they checked, where the form was, the audit they were shown (counts, page
// title and findings), and whether the agency was emailed (notifyError says
// why not). Answers 'who filled in our audit form'. Free.
func (c *Client) ListLeads(ctx context.Context, in ListLeadsInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_leads", in)
}

// ListLinkProspectsInput is the input to list_link_prospects.
type ListLinkProspectsInput struct {
	ClientID string `json:"clientId"`
}

// ListLinkProspects calls list_link_prospects: List a client's link prospects.
// The client's outreach list: sites worth asking for a link, each with its
// stage (identified, contacted, replied, won, rejected), why it was added,
// when it last moved, the contact to write to and, for won links, whether the
// last link check found the link live on its page — most recently moved first.
// Answers 'where is our link outreach' and 'are the links we won still up'.
// Free.
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

// ListLogSourcesInput is the input to list_log_sources.
type ListLogSourcesInput struct {
	ClientID string `json:"clientId"`
}

// ListLogSources calls list_log_sources: Connected log feeds. The client's
// connected log feeds: kind, when each last delivered, how many crawler
// requests it has stored and its last error (an S3 pull that was refused,
// say). Ingest URLs and keys are never shown again. Answers 'are our logs
// still arriving'. Free.
func (c *Client) ListLogSources(ctx context.Context, in ListLogSourcesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_log_sources", in)
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

// ListOutreachSequencesInput is the input to list_outreach_sequences.
type ListOutreachSequencesInput struct {
	ClientID string `json:"clientId"`
}

// ListOutreachSequences calls list_outreach_sequences: The client's outreach
// sequences. The sequences saved for the client, newest first, with their
// steps. Answers 'which sequences do we have'. Free.
func (c *Client) ListOutreachSequences(ctx context.Context, in ListOutreachSequencesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_outreach_sequences", in)
}

// ListPathOwnersInput is the input to list_path_owners.
type ListPathOwnersInput struct {
	ClientID string `json:"clientId"`
}

// ListPathOwners calls list_path_owners: List who owns which part of the site.
// The client's path owners from set_path_owners: prefix, owner and contact,
// shortest prefix first. Answers 'who fixes issues on /docs'. Free.
func (c *Client) ListPathOwners(ctx context.Context, in ListPathOwnersInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_path_owners", in)
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

// ListPortfoliosInput is the input to list_portfolios.
type ListPortfoliosInput struct {
}

// ListPortfolios calls list_portfolios: List portfolios. The agency's
// portfolios — named groups of clients or single sites — with their members. A
// viewer limited to some clients sees only their members, with a count of the
// rest. Answers 'which client groups do we have'. Free.
func (c *Client) ListPortfolios(ctx context.Context, in ListPortfoliosInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_portfolios", in)
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

// ListReportLinksInput is the input to list_report_links.
type ListReportLinksInput struct {
	ClientID string `json:"clientId,omitempty"`
}

// ListReportLinks calls list_report_links: List live report links. The share
// links made with create_report_link that still open, newest first, for one
// client or (without clientId) every client the caller can see: each with
// linkId, the report, its period, when it was made and by whom, when it
// expires, and when it was last opened (lastViewedAt, null if never) with a
// view count. Answers 'which report links are out there' and 'has the client
// opened the report'. The address itself is not returned: it is a secret shown
// once and never stored, so make a new link to send it again. Free.
func (c *Client) ListReportLinks(ctx context.Context, in ListReportLinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_report_links", in)
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

// ListTaskLinksInput is the input to list_task_links.
type ListTaskLinksInput struct {
	ClientID string `json:"clientId"`
}

// ListTaskLinks calls list_task_links: Tasks mirrored in trackers. Lists which
// of the client's tasks are mirrored as issues, where (PROJ-12, ENG-34, #56,
// with links), and the tracker's state as last read. Free.
func (c *Client) ListTaskLinks(ctx context.Context, in ListTaskLinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_task_links", in)
}

// ListTasksInput is the input to list_tasks.
type ListTasksInput struct {
	ClientID       string   `json:"clientId"`
	Statuses       []string `json:"statuses,omitempty"`
	AssigneeUserID string   `json:"assigneeUserId,omitempty"`
	UnassignedOnly *bool    `json:"unassignedOnly,omitempty"`
	Limit          *int64   `json:"limit,omitempty"`
}

// ListTasks calls list_tasks: List a client's tasks. A client's tasks, open
// ones first and soonest due first: title, status, assignee, due date, whether
// overdue, and the issue, page or keyword each is about. Answers 'what is left
// to do for Acme' and 'what is Sam working on'. Filter by status and assignee.
// Free.
func (c *Client) ListTasks(ctx context.Context, in ListTasksInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_tasks", in)
}

// ListTeamInvitesInput is the input to list_team_invites.
type ListTeamInvitesInput struct {
}

// ListTeamInvites calls list_team_invites: List pending invites. Invites not
// yet accepted or withdrawn, newest first, with role, client scope, when each
// expires and how often it was sent. Expired ones are listed so they can be
// resent. Answers 'who have we invited who has not joined'.
func (c *Client) ListTeamInvites(ctx context.Context, in ListTeamInvitesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_team_invites", in)
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

// ListWebhookDeliveriesInput is the input to list_webhook_deliveries.
type ListWebhookDeliveriesInput struct {
	WebhookID string `json:"webhookId"`
	Limit     *int64 `json:"limit,omitempty"`
}

// ListWebhookDeliveries calls list_webhook_deliveries: A webhook's deliveries.
// A webhook's recent deliveries, newest first: event, status, attempts, the
// receiver's last HTTP status and error. Answers 'why isn't my webhook
// receiving anything'. Free.
func (c *Client) ListWebhookDeliveries(ctx context.Context, in ListWebhookDeliveriesInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_webhook_deliveries", in)
}

// ListWebhooksInput is the input to list_webhooks.
type ListWebhooksInput struct {
}

// ListWebhooks calls list_webhooks: List webhooks. The agency's webhooks: URL,
// events, client, and whether one was switched off after failing for too long.
// Answers 'what webhooks do we have'. Free.
func (c *Client) ListWebhooks(ctx context.Context, in ListWebhooksInput) (json.RawMessage, error) {
	return c.Call(ctx, "list_webhooks", in)
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

// MineAppReviewsInput is the input to mine_app_reviews.
type MineAppReviewsInput struct {
	ClientID      string `json:"clientId"`
	Store         string `json:"store"`
	AppID         string `json:"appId"` // App Store: the numeric id (id123456789 or 123456789). Google Play: the package name.
	LocationCode  *int64 `json:"locationCode,omitempty"`
	LanguageCode  string `json:"languageCode,omitempty"`
	Limit         *int64 `json:"limit,omitempty"`
	TaskID        string `json:"taskId,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// MineAppReviews calls mine_app_reviews: What an app's reviews keep saying.
// Reads an app's recent reviews and mines them: the rating split, themes
// (crashes and bugs, performance, sign-in, ads, pricing and subscriptions,
// support, design, features) with how often each is raised and how negative,
// the phrases most repeated in 1–2 star reviews and in 4–5 star ones, and the
// feature requests. Works on a competitor's app as well as the client's. App
// Store reviews come free from Apple's public reviews feed (most recent 500);
// Google Play reviews through DataForSEO's App Data, queued: pending with a
// taskId to pass again when not back within half a minute. Answers 'what do
// users complain about', 'what do they love' and 'what are they asking for'.
// App Store: free. Google Play: A paid lookup on the client's account (spend
// says how much); quote it with estimate_cost.
func (c *Client) MineAppReviews(ctx context.Context, in MineAppReviewsInput) (json.RawMessage, error) {
	return c.Call(ctx, "mine_app_reviews", in)
}

// NotifyGoogleIndexingInput is the input to notify_google_indexing.
type NotifyGoogleIndexingInput struct {
	ClientID string   `json:"clientId"`
	URLs     []string `json:"urls"`
	Type     string   `json:"type,omitempty"`
	DryRun   *bool    `json:"dryRun,omitempty"` // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
}

// NotifyGoogleIndexing calls notify_google_indexing: Tell Google a job posting
// or livestream changed. Sends Google's Indexing API a notification for each
// address: URL_UPDATED for a new or changed page, URL_DELETED for one taken
// down. Google accepts it only for pages with JobPosting markup or a
// BroadcastEvent inside a VideoObject, so each page is checked first — read
// live for an update, from the latest crawl for a removal — and any other page
// is refused with the reason (use submit_indexnow or a sitemap for those). At
// most 200 a day per agency. Needs Search Console connected as the agency's
// Google account (verify_site) with that account an Owner of the property; a
// Google sign-in connection cannot send these. dryRun: true checks every page
// the same way but sends nothing and spends none of the daily allowance,
// returning wouldSubmit and refused. Answers 'get this job posting indexed
// now' and 'tell Google the vacancy is filled'.
func (c *Client) NotifyGoogleIndexing(ctx context.Context, in NotifyGoogleIndexingInput) (json.RawMessage, error) {
	return c.Call(ctx, "notify_google_indexing", in)
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
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// PageLab calls page_lab: How a page scores in one synthetic run. Lighthouse
// performance, SEO, accessibility and best-practices scores for one address,
// from PageSpeed Insights. Answers 'what is slow or wrong on this page' with a
// reproducible lab run, including for pages too quiet for field data.
// diagnostics says where the time went — the LCP element and its subparts
// (TTFB, resource load delay, resource load duration, element render delay),
// DOM size, long tasks, main-thread work, script execution, render-blocking
// requests, unused JavaScript and CSS, third parties — and fixes ranks the
// findings by estimated saving. This is not what real visitors experienced —
// use page_vitals for that, and say which is which when reporting both. There
// is no interaction metric here: a synthetic run interacts with nothing.
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

// PlanContentInput is the input to plan_content.
type PlanContentInput struct {
	ClientID         string   `json:"clientId"`
	Months           *int64   `json:"months,omitempty"`
	Cadence          string   `json:"cadence,omitempty"`
	StartMonth       string   `json:"startMonth,omitempty"` // Default: next month.
	Seed             string   `json:"seed,omitempty"`
	ExtraKeywords    []string `json:"extraKeywords,omitempty"`
	Trends           *bool    `json:"trends,omitempty"`
	IncludeRefreshes *bool    `json:"includeRefreshes,omitempty"`
	MaxCostMicros    *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun           *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// PlanContent calls plan_content: Plan the editorial calendar. Builds and
// stores an editorial calendar for the coming months at a cadence (weekly ≈ 4
// a month, fortnightly 2, monthly 1, or a number a month) from the topic map
// (Search Console clusters with no strong page: write; clusters on page two:
// refresh), extraKeywords not yet covered, the keyword registry (a keyword
// another page owns becomes a refresh of that page, never a second page), and
// seasonality (stored monthly volumes; trends: true reads Google Trends for
// the top items, paid): seasonal pieces land two months before their peak, the
// rest in order of value. Each item carries its keyword and cluster,
// ready-made build_content_brief arguments, target URL, internal links to add
// (to and from), search volume, estimated monthly clicks at position 3 and
// their value (clicks × CPC). hubs lays each pillar with three or more
// clusters out as a hub — one pillar (2,500–4,000 words), 2–5 clusters of 2–4
// spokes (1,200–1,800 words), templates by intent, a link matrix (spoke ↔
// pillar mandatory, 2–3 spoke-to-spoke, at most one cross-cluster) — and
// scores the live site's links against it: no orphans, pillar links both ways,
// 80% of planned links present. Answers 'what should we publish over the next
// quarter'. Read back with list_content_plan; change items with
// update_content_plan_item.
func (c *Client) PlanContent(ctx context.Context, in PlanContentInput) (json.RawMessage, error) {
	return c.Call(ctx, "plan_content", in)
}

// PlanRedirectMapInput is the input to plan_redirect_map.
type PlanRedirectMapInput struct {
	ClientID      string   `json:"clientId"`
	NewCrawlID    string   `json:"newCrawlId,omitempty"`
	OldURLs       []string `json:"oldUrls,omitempty"`
	OldSitemapURL string   `json:"oldSitemapUrl,omitempty"`
	OldCrawlID    string   `json:"oldCrawlId,omitempty"`
	Format        string   `json:"format,omitempty"`
	MinConfidence *float64 `json:"minConfidence,omitempty"`
	Limit         *int64   `json:"limit,omitempty"`
}

// PlanRedirectMap calls plan_redirect_map: Redirect map for a migration.
// Proposes a 1:1 301 redirect for each old address of a site migration, onto
// the page of the new site's crawl (newCrawlId, default latest) that most
// plausibly replaced it, with a confidence (0–1) and the evidence: the new
// site already redirects it, the same path (also ignoring case, slashes and
// .html/.php), the same text, or similar path words, title and text; failing
// those the nearest surviving section, and only then the homepage (flagged:
// Google treats mass homepage redirects as soft 404s). Give exactly one old
// source: oldUrls (a list), oldSitemapUrl (the old site's sitemap) or
// oldCrawlId (an earlier crawl of this client — best, since titles and text
// are compared too). format csv, htaccess (exact-match RewriteRules) or nginx
// (a map block) returns the whole map as a file in content. Answers 'where
// should each old URL redirect'. Free.
func (c *Client) PlanRedirectMap(ctx context.Context, in PlanRedirectMapInput) (json.RawMessage, error) {
	return c.Call(ctx, "plan_redirect_map", in)
}

// PostReviewReplyInput is the input to post_review_reply.
type PostReviewReplyInput struct {
	ClientID   string `json:"clientId"`
	LocationID string `json:"locationId"`
	ReviewID   string `json:"reviewId"`
	Reply      string `json:"reply"`
}

// PostReviewReply calls post_review_reply: Post a reply to a Google review.
// Posts the owner's reply to one Google review — or replaces the reply already
// there — through the client's connected Business Profile, and records it as
// answered. Needs the listing connected under Local (the connection asks
// Google for permission to manage it); reviewId is from draft_review_replies
// or get_local_visibility. Answers 'post that reply'. Free.
func (c *Client) PostReviewReply(ctx context.Context, in PostReviewReplyInput) (json.RawMessage, error) {
	return c.Call(ctx, "post_review_reply", in)
}

// PreviewWeeklyDigestInput is the input to preview_weekly_digest.
type PreviewWeeklyDigestInput struct {
	ClientID string `json:"clientId"`
}

// PreviewWeeklyDigest calls preview_weekly_digest: This week's digest. The
// weekly digest for the client as it would be posted, for the last complete
// Monday-to-Sunday week (week; the title names its days): what changed in
// Search Console, revenue and AI assistant visits; why it probably changed
// (demand, rankings, or particular pages); and the top five actions — from
// striking distance, low click-through, decaying pages and the latest audit —
// with the extra monthly clicks each is estimated to bring and, where revenue
// or goal values allow, what they are worth. Audit findings accepted as
// intentional (accept_issue) are left out and counted in basis; with path
// owners set, each audit action names its owner and byOwner groups them.
// Answers 'what should we do this week and why'. Free.
func (c *Client) PreviewWeeklyDigest(ctx context.Context, in PreviewWeeklyDigestInput) (json.RawMessage, error) {
	return c.Call(ctx, "preview_weekly_digest", in)
}

// ProposeKeywordOwnersInput is the input to propose_keyword_owners.
type ProposeKeywordOwnersInput struct {
	ClientID string `json:"clientId"`
	Days     *int64 `json:"days,omitempty"`
	Limit    *int64 `json:"limit,omitempty"`
	Seed     string `json:"seed,omitempty"`
}

// ProposeKeywordOwners calls propose_keyword_owners: Propose an owner page for
// each query cluster. Groups the client's non-branded Search Console queries
// by the page they land on most and proposes, for each cluster, the page
// Search Console sends the most clicks to (impressions when none has clicks):
// its share of the cluster's clicks, the other pages competing, whether it is
// already registered (and to which page), and confidence clear, split (the
// clicks are divided: cannibalization) or thin. Nothing is saved: accept with
// set_keyword_owners. Answers 'which page should own each topic'. Free.
func (c *Client) ProposeKeywordOwners(ctx context.Context, in ProposeKeywordOwnersInput) (json.RawMessage, error) {
	return c.Call(ctx, "propose_keyword_owners", in)
}

// PushTasksToTrackerInput is the input to push_tasks_to_tracker.
type PushTasksToTrackerInput struct {
	ClientID string   `json:"clientId"`
	Provider string   `json:"provider"`
	TaskIDs  []string `json:"taskIds,omitempty"`
	Statuses []string `json:"statuses,omitempty"`
	Limit    *int64   `json:"limit,omitempty"`
	DryRun   *bool    `json:"dryRun,omitempty"` // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
}

// PushTasksToTracker calls push_tasks_to_tracker: Mirror tasks into Jira,
// Linear or GitHub. Creates an issue in the client's connected Jira project,
// Linear team or GitHub repository for each task (the chosen taskIds, or every
// task of the chosen statuses — default the open ones), with the task's notes
// and what it is about, and links the two so status syncs both ways. Tasks
// already mirrored there are skipped. dryRun: true creates nothing and lists
// wouldCreate (each issue's title and description) and alreadyLinked. Pairs
// with create_tasks_from_audit. Answers 'put the audit into our Jira'. Connect
// the tracker first with connect_integration. Free.
func (c *Client) PushTasksToTracker(ctx context.Context, in PushTasksToTrackerInput) (json.RawMessage, error) {
	return c.Call(ctx, "push_tasks_to_tracker", in)
}

// QueryDataInput is the input to query_data.
type QueryDataInput struct {
	ClientID   string           `json:"clientId"`
	Dataset    string           `json:"dataset"`
	Dimensions []string         `json:"dimensions,omitempty"`
	Metrics    []string         `json:"metrics,omitempty"`
	Columns    []string         `json:"columns,omitempty"`
	Filters    []map[string]any `json:"filters,omitempty"`
	From       string           `json:"from,omitempty"`
	To         string           `json:"to,omitempty"`
	OrderBy    []map[string]any `json:"orderBy,omitempty"` // Sort by a selected dimension or metric: [{ "column": "sum(clicks)", "desc": true }].
	CrawlID    string           `json:"crawlId,omitempty"` // For crawl datasets: a crawl id or latest.
	Limit      *int64           `json:"limit,omitempty"`
}

// QueryData calls query_data: Query the stored data. Asks any question of a
// client's stored data the other tools do not answer: pick a dataset
// (describe_datasets), group by dimensions, compute metrics (count, sum(col),
// avg(col), min, max, uniq, plus ctr and position for Search Console, visitors
// and visits for analytics), filter (eq, neq, contains, starts_with, gt, lt,
// in, regex), order and limit — or list rows with columns. Search Console can
// also be grouped or filtered by each query's intent, topic and branded flag.
// Answers 'which queries containing X lost clicks on mobile in Spain', 'how
// many 404s per template', 'how are commercial queries trending'. Read-only
// and always limited to the client's own data. Up to 10,000 rows; export_data
// for more. Free.
func (c *Client) QueryData(ctx context.Context, in QueryDataInput) (json.RawMessage, error) {
	return c.Call(ctx, "query_data", in)
}

// QueryKeywordListInput is the input to query_keyword_list.
type QueryKeywordListInput struct {
	ListID        string   `json:"listId"`
	Text          string   `json:"text,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Untagged      *bool    `json:"untagged,omitempty"`
	Intents       []string `json:"intents,omitempty"`
	MinVolume     *int64   `json:"minVolume,omitempty"`
	MaxVolume     *int64   `json:"maxVolume,omitempty"`
	MinDifficulty *int64   `json:"minDifficulty,omitempty"`
	MaxDifficulty *int64   `json:"maxDifficulty,omitempty"`
	Sort          string   `json:"sort,omitempty"`
	Order         string   `json:"order,omitempty"` // Default: A to Z for term, oldest first for added, biggest first otherwise.
	Limit         *int64   `json:"limit,omitempty"`
	Offset        *int64   `json:"offset,omitempty"`
	Format        string   `json:"format,omitempty"`
}

// QueryKeywordList calls query_keyword_list: Filter and sort a keyword list.
// One list's keywords with the metrics last read for each — volume, keyword
// difficulty, CPC (micro-USD), competition, intent, SERP features and when
// they were read — plus tags and notes, filtered by text, tags (all must
// match), untagged, intent and volume or difficulty ranges, and sorted by
// term, volume, difficulty, cpc, added or refreshed. Keywords with no metrics
// sort last. format: csv also returns the page as CSV for a spreadsheet.
// Answers 'which keywords on that list are easy and over 1,000 searches' and
// 'export that list'. Metrics are null until refresh_keyword_list_metrics has
// run. Free.
func (c *Client) QueryKeywordList(ctx context.Context, in QueryKeywordListInput) (json.RawMessage, error) {
	return c.Call(ctx, "query_keyword_list", in)
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
	ClientID   string   `json:"clientId"`
	Render     *bool    `json:"render,omitempty"`     // True to read the unverifiable_js pages (or sourceUrls) in a real browser, up to 25.
	SourceURLs []string `json:"sourceUrls,omitempty"` // With render: the linking pages to read, from get_link_checks. Default: those marked unverifiable_js.
}

// RecheckLinks calls recheck_links: Re-check the pages that link to a client.
// Sends our crawler to the pages said to link to the client to see whether
// each still does — slowly and politely, so it takes a while; read the result
// with get_link_checks. A second request within the hour lands on the check
// already running. render: true instead reads, in a real browser, the pages
// the last check marked unverifiable_js (app shells whose links JavaScript
// draws) — or the sourceUrls given, if the link checks know them — up to 25 a
// run, and re-judges each link as live, lost or moved (a page that draws
// nothing even in the browser stays unverifiable_js, and one that will not
// load keeps its verdict); it runs in the background, so poll get_run_status
// with job 'links' and the runId, then read get_link_checks (rendered: true on
// those pages). No vendor charge, but it visits other people's sites on the
// client's behalf, so it needs client:manage. Refused when the client has no
// rank tracker yet (HTML) or no page is waiting on a browser (render:
// nothing_to_render). Answers 'check our links again' and 'check the
// JavaScript pages properly'.
func (c *Client) RecheckLinks(ctx context.Context, in RecheckLinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "recheck_links", in)
}

// RecordOrderInput is the input to record_order.
type RecordOrderInput struct {
	ClientID    string `json:"clientId"`
	OrderID     string `json:"orderId"`
	AmountMinor int64  `json:"amountMinor"`
	Currency    string `json:"currency"`
	At          string `json:"at"`
	Visit       string `json:"visit,omitempty"`
	RefundOf    string `json:"refundOf,omitempty"`
	CustomerID  string `json:"customerId,omitempty"`
}

// RecordOrder calls record_order: Record an order (no store connected).
// Records one order for a business not on Stripe, Shopify or Paddle, so
// get_seo_revenue counts it: orderId (the site's own; recording it again
// replaces it, a replay is answered 'unchanged'), amountMinor in the
// currency's minor unit (4999 = 49.99 USD), currency, at (when paid) and, when
// the site kept one, visit — the visit token /api/collect gave its server
// collector while the buyer was on the site, which credits the order to that
// visit's landing page, source and channel. A refund is an order of its own
// with refundOf naming the order it refunds. Answers 'log this crypto payment
// against the visit it came from'. Free.
func (c *Client) RecordOrder(ctx context.Context, in RecordOrderInput) (json.RawMessage, error) {
	return c.Call(ctx, "record_order", in)
}

// RecordOrdersInput is the input to record_orders.
type RecordOrdersInput struct {
	ClientID string           `json:"clientId"`
	Orders   []map[string]any `json:"orders"`
}

// RecordOrders calls record_orders: Record orders in a batch. record_order for
// up to 500 orders in one call, for a server pushing its orders (POST
// /api/v1/tools/record_orders with an Idempotency-Key). Each is answered on
// its own — recorded, updated, unchanged or refused with the reason — and one
// refused order does not stop the rest. Refunds may refer to orders earlier in
// the same batch. Answers 'send last night's orders in one request'. Free.
func (c *Client) RecordOrders(ctx context.Context, in RecordOrdersInput) (json.RawMessage, error) {
	return c.Call(ctx, "record_orders", in)
}

// RefreshDomainDataInput is the input to refresh_domain_data.
type RefreshDomainDataInput struct {
	ClientID      string `json:"clientId"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// RefreshKeywordListMetricsInput is the input to refresh_keyword_list_metrics.
type RefreshKeywordListMetricsInput struct {
	ClientID          string `json:"clientId"`
	ListID            string `json:"listId"`
	OnlyOlderThanDays *int64 `json:"onlyOlderThanDays,omitempty"`
	MaxCostMicros     *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun            *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// RefreshKeywordListMetrics calls refresh_keyword_list_metrics: Refresh a
// keyword list's metrics. Re-reads volume, keyword difficulty, CPC,
// competition, intent and SERP features for every Google keyword on a list
// from DataForSEO Labs, in bulk (one read per market per 700 keywords), and
// stores them on the list. Answers 'update the numbers on that list'.
// onlyOlderThanDays skips keywords refreshed more recently, so a big list is
// not bought twice in a week. Keywords for other engines, or markets the
// research index does not cover, are skipped with the reason; terms the index
// has nothing for are listed in noData and their old figures cleared. clientId
// is the account billed: a client's list is refreshed under its own client, an
// agency list under any client the caller names. A paid lookup unless anyone
// read the same thing recently; quote it with estimate_cost.
func (c *Client) RefreshKeywordListMetrics(ctx context.Context, in RefreshKeywordListMetricsInput) (json.RawMessage, error) {
	return c.Call(ctx, "refresh_keyword_list_metrics", in)
}

// RefreshRankingsInput is the input to refresh_rankings.
type RefreshRankingsInput struct {
	ClientID      string `json:"clientId"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// RefreshRankings calls refresh_rankings: Capture rankings now. Captures the
// client's rank tracker now instead of at its next scheduled run: every
// tracked keyword on every device the tracker reads, each a paid lookup.
// Answers 'get me today's positions'. Quote it with estimate_cost first; with
// maxCostMicros the run is refused before it starts when its estimate is over
// the limit. Runs in the background: poll get_run_status with job 'rank', then
// read get_rank_history or get_rank_summary. A second request within the same
// minute lands on the run already going rather than paying twice.
func (c *Client) RefreshRankings(ctx context.Context, in RefreshRankingsInput) (json.RawMessage, error) {
	return c.Call(ctx, "refresh_rankings", in)
}

// RemoveAnnotationsInput is the input to remove_annotations.
type RemoveAnnotationsInput struct {
	ClientID string   `json:"clientId"`
	IDs      []string `json:"ids"`
}

// RemoveAnnotations calls remove_annotations: Remove timeline notes. Removes
// the client's timeline notes by id; unknown ids are ignored. Answers 'take
// that note off'. Google's updates cannot be removed. Free.
func (c *Client) RemoveAnnotations(ctx context.Context, in RemoveAnnotationsInput) (json.RawMessage, error) {
	return c.Call(ctx, "remove_annotations", in)
}

// RemoveClientFactsInput is the input to remove_client_facts.
type RemoveClientFactsInput struct {
	ClientID string   `json:"clientId"`
	Keys     []string `json:"keys"`
}

// RemoveClientFacts calls remove_client_facts: Remove facts from the registry.
// Removes the client's facts by key; unknown keys are ignored. Answers 'forget
// the old enterprise price'. Free.
func (c *Client) RemoveClientFacts(ctx context.Context, in RemoveClientFactsInput) (json.RawMessage, error) {
	return c.Call(ctx, "remove_client_facts", in)
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

// RemoveKeywordOwnersInput is the input to remove_keyword_owners.
type RemoveKeywordOwnersInput struct {
	ClientID string   `json:"clientId"`
	Keywords []string `json:"keywords"`
}

// RemoveKeywordOwners calls remove_keyword_owners: Release keywords from the
// registry. Removes the registry's claims on the given keywords; unknown
// keywords are ignored. Answers 'that keyword no longer belongs to that page'.
// Free.
func (c *Client) RemoveKeywordOwners(ctx context.Context, in RemoveKeywordOwnersInput) (json.RawMessage, error) {
	return c.Call(ctx, "remove_keyword_owners", in)
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

// ReplayWebhooksInput is the input to replay_webhooks.
type ReplayWebhooksInput struct {
	WebhookID  string `json:"webhookId"`
	Since      string `json:"since,omitempty"`
	DeliveryID string `json:"deliveryId,omitempty"`
}

// ReplayWebhooks calls replay_webhooks: Send events again. Queues events for a
// webhook again: one delivery (deliveryId), or every subscribed event since an
// ISO time (default: the last 24 hours). For a receiver that was down. Answers
// 'resend what we missed'. Free.
func (c *Client) ReplayWebhooks(ctx context.Context, in ReplayWebhooksInput) (json.RawMessage, error) {
	return c.Call(ctx, "replay_webhooks", in)
}

// ResearchTopicInput is the input to research_topic.
type ResearchTopicInput struct {
	ClientID      string `json:"clientId"`
	Seed          string `json:"seed"`
	Limit         *int64 `json:"limit,omitempty"`
	Serp          *bool  `json:"serp,omitempty"`          // Also read the results page for People Also Ask and headlines.
	LocationCode  *int64 `json:"locationCode,omitempty"`  // Google location code (2840 United States, 2826 United Kingdom). Default: the client's project.
	LanguageCode  string `json:"languageCode,omitempty"`  // Language code (en, es). Default: the client's project.
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// ResearchTopic calls research_topic: Subtopics, questions and headlines
// around a seed. Maps a topic: related keywords and keyword suggestions for a
// seed, with volume, difficulty and search intent, grouped into subtopics by
// the word they add to the seed; the questions people ask (People Also Ask and
// question keywords); and the headlines ranking now. Answers 'what should we
// write about around this topic'. Paid DataForSEO Labs reads and one results
// page on the client's account, each free when anyone read it lately.
func (c *Client) ResearchTopic(ctx context.Context, in ResearchTopicInput) (json.RawMessage, error) {
	return c.Call(ctx, "research_topic", in)
}

// ResendTeamInviteInput is the input to resend_team_invite.
type ResendTeamInviteInput struct {
	InviteID string `json:"inviteId"`
}

// ResendTeamInvite calls resend_team_invite: Resend an invite. Sends an invite
// again with a new link and a fresh 7 days; the old link stops working. By the
// inviteId list_team_invites gives. At most 5 sends per invite and one a
// minute. Answers 'they never got the email'.
func (c *Client) ResendTeamInvite(ctx context.Context, in ResendTeamInviteInput) (json.RawMessage, error) {
	return c.Call(ctx, "resend_team_invite", in)
}

// RestoreClientInput is the input to restore_client.
type RestoreClientInput struct {
	ClientID string `json:"clientId"`
}

// RestoreClient calls restore_client: Restore an archived client. Brings an
// archived client back: it is listed, tracked, crawled and reported on again,
// with its history as it was, and counts against the plan once more (refused
// when the plan's client limit is reached). Answers 'we're working with them
// again'. Safe to repeat: restored is false when the client was not archived.
func (c *Client) RestoreClient(ctx context.Context, in RestoreClientInput) (json.RawMessage, error) {
	return c.Call(ctx, "restore_client", in)
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

// RevokeLogSourceInput is the input to revoke_log_source.
type RevokeLogSourceInput struct {
	ClientID string `json:"clientId"`
	SourceID string `json:"sourceId"`
}

// RevokeLogSource calls revoke_log_source: Disconnect a log feed. Stops a log
// feed: its ingest URL refuses deliveries from now on, and an S3 source stops
// being pulled and its keys are deleted. Stored requests stay. Answers
// 'disconnect that log drain'. Free.
func (c *Client) RevokeLogSource(ctx context.Context, in RevokeLogSourceInput) (json.RawMessage, error) {
	return c.Call(ctx, "revoke_log_source", in)
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

// RevokeReportLinkInput is the input to revoke_report_link.
type RevokeReportLinkInput struct {
	LinkID string `json:"linkId"`
}

// RevokeReportLink calls revoke_report_link: Turn off one report link. Stops
// one link made with create_report_link (linkId from list_report_links or from
// create_report_link's answer) at once, leaving the report's other links
// working. Answers 'I sent the report link to the wrong person'. Revoking a
// link twice keeps the first time. Free.
func (c *Client) RevokeReportLink(ctx context.Context, in RevokeReportLinkInput) (json.RawMessage, error) {
	return c.Call(ctx, "revoke_report_link", in)
}

// RevokeReportLinksInput is the input to revoke_report_links.
type RevokeReportLinksInput struct {
	ClientID string `json:"clientId"`
	ReportID string `json:"reportId"`
}

// RevokeReportLinks calls revoke_report_links: Turn off a report's links.
// Stops every link made so far with create_report_link for one saved report,
// at once. Answers 'turn off all the links to this report'. To stop one link
// and keep the rest, revoke_report_link. Links made afterwards work as normal.
// Free.
func (c *Client) RevokeReportLinks(ctx context.Context, in RevokeReportLinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "revoke_report_links", in)
}

// RevokeTeamInviteInput is the input to revoke_team_invite.
type RevokeTeamInviteInput struct {
	InviteID string `json:"inviteId"`
}

// RevokeTeamInvite calls revoke_team_invite: Withdraw an invite. Withdraws an
// invite so its link no longer works. Answers 'cancel the invite to sam@'.
// Safe to repeat: revoked is false when it was already used or withdrawn.
func (c *Client) RevokeTeamInvite(ctx context.Context, in RevokeTeamInviteInput) (json.RawMessage, error) {
	return c.Call(ctx, "revoke_team_invite", in)
}

// RewriteSectionInput is the input to rewrite_section.
type RewriteSectionInput struct {
	ClientID      string   `json:"clientId"`
	Keyword       string   `json:"keyword"`
	Section       string   `json:"section"` // The section, Markdown or text.
	Mode          string   `json:"mode,omitempty"`
	MissingTerms  []string `json:"missingTerms,omitempty"`
	Questions     []string `json:"questions,omitempty"`
	Draft         string   `json:"draft,omitempty"` // The whole page as Markdown, scored to find what is missing.
	URL           string   `json:"url,omitempty"`
	Instructions  string   `json:"instructions,omitempty"`
	LocationCode  *int64   `json:"locationCode,omitempty"`  // Google location code (2840 United States, 2826 United Kingdom). Default: the client's project.
	LanguageCode  string   `json:"languageCode,omitempty"`  // Language code (en, es). Default: the client's project.
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// RewriteSection calls rewrite_section: Rewrite a section to cover what is
// missing. Rewrites (mode rewrite) or extends (mode extend) one section of a
// page so it covers the terms, entities and questions the content score finds
// missing, keeping the section's own voice and the client's. Give missingTerms
// and questions yourself, or the whole draft (draft as Markdown, or url) to
// have them taken from a fresh score_content run. Returns the new Markdown and
// which asked-for terms it now uses and still misses, as we count them.
// Answers 'work these terms into this section'. A paid Claude call; scoring a
// draft also costs what score_content does.
func (c *Client) RewriteSection(ctx context.Context, in RewriteSectionInput) (json.RawMessage, error) {
	return c.Call(ctx, "rewrite_section", in)
}

// RunAIVisibilityCheckInput is the input to run_ai_visibility_check.
type RunAIVisibilityCheckInput struct {
	ClientID      string `json:"clientId"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// RunAuditPerformanceSampleInput is the input to run_audit_performance_sample.
type RunAuditPerformanceSampleInput struct {
	ClientID      string `json:"clientId"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// RunAuditPerformanceSample calls run_audit_performance_sample: Measure the
// audit's speed again. Takes the latest complete audit's speed sample again,
// in the background: a Lighthouse run (mobile) of the homepage and one page of
// each of the largest URL templates (up to 5 pages), and the origin's Chrome
// UX Report reading for phone and desktop — the same sample an audit takes
// when it finishes, replacing that audit's figures. For an audit that finished
// before a PageSpeed Insights key was connected, or whose runs failed. Takes a
// few minutes; read the result with get_audit_performance. Uses the agency's
// PageSpeed Insights and Chrome UX Report connections. Answers 'measure the
// site's speed now'.
func (c *Client) RunAuditPerformanceSample(ctx context.Context, in RunAuditPerformanceSampleInput) (json.RawMessage, error) {
	return c.Call(ctx, "run_audit_performance_sample", in)
}

// RunDataSyncNowInput is the input to run_data_sync_now.
type RunDataSyncNowInput struct {
	ClientID string `json:"clientId"`
	SyncID   string `json:"syncId"`
}

// RunDataSyncNow calls run_data_sync_now: Run a data sync now. Runs one sync
// immediately, outside its schedule, and says how many rows were written.
// Answers 'refresh the sheet now'. Free.
func (c *Client) RunDataSyncNow(ctx context.Context, in RunDataSyncNowInput) (json.RawMessage, error) {
	return c.Call(ctx, "run_data_sync_now", in)
}

// RunLocalGridScanInput is the input to run_local_grid_scan.
type RunLocalGridScanInput struct {
	ClientID      string   `json:"clientId"`
	LocationID    string   `json:"locationId"`
	Term          string   `json:"term"`
	Size          *float64 `json:"size,omitempty"` // Grid width in points: 3, 5, 7, 9 or 13.
	SpacingKm     *float64 `json:"spacingKm,omitempty"`
	Device        string   `json:"device,omitempty"`
	MoreTerms     []string `json:"moreTerms,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// RunLocalGridScan calls run_local_grid_scan: Run a map-pack grid scan.
// Searches the term at every point of a size × size grid around the location
// (spacingKm apart, 1.5 by default — 0.5 for a dense city centre, 5 for a
// rural area) and records where the business appears in the map pack. Sizes 3,
// 5, 7, 9 (suburban or a wide service area) and 13 (rural or a large metro);
// device mobile asks as a phone, which sees a different pack; moreTerms (up to
// two) scans other terms on the same grid — the service, the brand with the
// town, a long-tail query — each its own scan. Every point of every term is a
// paid lookup, so a 7 × 7 grid is 49 lookups and 13 × 13 is 169: quote it with
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
	UpdatePrompts   []map[string]any `json:"updatePrompts,omitempty"`
	RetirePromptIDs []string         `json:"retirePromptIds,omitempty"`
	AddBrands       []map[string]any `json:"addBrands,omitempty"`
	DryRun          *bool            `json:"dryRun,omitempty"` // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
}

// SaveAIVisibilitySetup calls save_ai_visibility_setup: Change what AI
// visibility tracks. Adds questions (each optionally with a two-letter market,
// a topic, and branded true/false to override the automatic reading), regroups
// or reclassifies questions by id (updatePrompts: topic null to ungroup,
// branded null to go back to automatic), retires questions by id (past answers
// keep their question), and adds brands to look for (name, domain, aliases;
// isClient for the client's own — there should be exactly one). Questions
// already tracked are skipped. dryRun: true changes nothing and lists what
// would: the questions it would add, retire and regroup, those it would skip,
// and the brands it would add (with why one would be refused). Write questions
// the way a buyer would ask an assistant, without naming any brand;
// suggest_ai_prompts drafts them. Answers 'track these questions'. Free; the
// checks that follow are what cost.
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

// SaveDisavowFileInput is the input to save_disavow_file.
type SaveDisavowFileInput struct {
	ClientID string           `json:"clientId"`
	Entries  []map[string]any `json:"entries"`
}

// SaveDisavowFile calls save_disavow_file: Build, keep and compare a disavow
// file. Builds a Google disavow file for the client from domains and URLs (as
// build_disavow_file does), keeps it in the client's disavow history, and says
// what changed against the previous saved file: domains and URLs added and
// dropped. Nothing is submitted — Google takes the file by upload in Search
// Console, which a person does after reviewing every line. Answers 'make this
// month's disavow file and show what changed'. Free.
func (c *Client) SaveDisavowFile(ctx context.Context, in SaveDisavowFileInput) (json.RawMessage, error) {
	return c.Call(ctx, "save_disavow_file", in)
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
// DataForSEO (the API login and API password, or leave login out and pass the
// Base64 login:password DataForSEO's panel shows as secret; the check also
// confirms the account is verified), SerpApi, a proxy pool (serpproxy: one URL
// per line), Google Ads (six values in googleAds), an Anthropic, OpenAI,
// Gemini or Perplexity key, and the rest get_vendor_connections names — and
// checks it works straight away with one free call. Answers 'connect our
// DataForSEO account'. A failed check gives errorCode
// (vendor_credentials_rejected, vendor_account_unverified,
// vendor_balance_empty, vendor_product_not_enabled) and next, the step that
// fixes it. Replaces any key already saved for that vendor. The response says
// whether the check passed; it never contains the key.
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

// ScheduleWeeklyDigestInput is the input to schedule_weekly_digest.
type ScheduleWeeklyDigestInput struct {
	ClientID string `json:"clientId"`
	Provider string `json:"provider"`
	Weekday  *int64 `json:"weekday,omitempty"`
	RunHour  *int64 `json:"runHour,omitempty"`
	Paused   *bool  `json:"paused,omitempty"`
}

// ScheduleWeeklyDigest calls schedule_weekly_digest: Post the weekly digest to
// a chat. Posts the weekly digest (see preview_weekly_digest) to the client's
// connected Slack, Discord or Telegram every week on a weekday (0 Sunday – 6
// Saturday, default Monday) at a UTC hour. One schedule per chat connection;
// calling again changes it. Answers 'send Acme's weekly digest to their Slack
// on Mondays'.
func (c *Client) ScheduleWeeklyDigest(ctx context.Context, in ScheduleWeeklyDigestInput) (json.RawMessage, error) {
	return c.Call(ctx, "schedule_weekly_digest", in)
}

// ScoreAmazonListingInput is the input to score_amazon_listing.
type ScoreAmazonListingInput struct {
	ClientID      string `json:"clientId"`
	Asin          string `json:"asin"`
	Keyword       string `json:"keyword"`
	LocationCode  *int64 `json:"locationCode,omitempty"`
	LanguageCode  string `json:"languageCode,omitempty"`
	Competitors   *int64 `json:"competitors,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// ScoreAmazonListing calls score_amazon_listing: Score an Amazon listing
// against the top results. Reads an Amazon product (by ASIN) and the listings
// ranking for a keyword, and scores the listing: title length and whether it
// carries the keyword, the bullet points (count and length), images, the terms
// the top titles share that this one lacks, and its rating, ratings count and
// price against the top listings' median. Each check pass, warn or fail with
// the reason. Answers 'how does our Amazon listing compare' and 'what should
// our title say'. Two reads from DataForSEO's Amazon data. A paid lookup on
// the client's account (spend says how much); quote it with estimate_cost.
func (c *Client) ScoreAmazonListing(ctx context.Context, in ScoreAmazonListingInput) (json.RawMessage, error) {
	return c.Call(ctx, "score_amazon_listing", in)
}

// ScoreAppMetadataInput is the input to score_app_metadata.
type ScoreAppMetadataInput struct {
	ClientID      string   `json:"clientId"`
	Store         string   `json:"store"`
	AppID         string   `json:"appId"` // App Store: the numeric id (id123456789 or 123456789). Google Play: the package name.
	Keywords      []string `json:"keywords,omitempty"`
	LocationCode  *int64   `json:"locationCode,omitempty"`
	LanguageCode  string   `json:"languageCode,omitempty"`
	TaskID        string   `json:"taskId,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// ScoreAppMetadata calls score_app_metadata: Score an app's store listing.
// Scores an App Store or Google Play listing for search: the title's length
// and whether it carries the target keywords, the subtitle (App Store) or
// short description (Play), the description's length and keyword use,
// screenshots, rating, ratings count and how recently it was updated — each
// pass, warn or fail with the reason — and, per keyword, where the listing
// uses it. App Store listings are read from Apple's free lookup API; a Play
// listing is read through DataForSEO's App Data, which queues the read: when
// it is not back within half a minute the answer is pending with a taskId to
// pass again. Answers 'how good is our app listing' and 'is our keyword in the
// title'. App Store: free. Google Play: A paid lookup on the client's account
// (spend says how much); quote it with estimate_cost.
func (c *Client) ScoreAppMetadata(ctx context.Context, in ScoreAppMetadataInput) (json.RawMessage, error) {
	return c.Call(ctx, "score_app_metadata", in)
}

// ScoreContentInput is the input to score_content.
type ScoreContentInput struct {
	ClientID      string `json:"clientId"`
	Keyword       string `json:"keyword"`
	URL           string `json:"url,omitempty"`
	Html          string `json:"html,omitempty"`
	Markdown      string `json:"markdown,omitempty"`
	LocationCode  *int64 `json:"locationCode,omitempty"`  // Google location code (2840 United States, 2826 United Kingdom, 2724 Spain). Default: the client's project.
	LanguageCode  string `json:"languageCode,omitempty"`  // Language code (en, es). Default: the client's project.
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// ScoreContent calls score_content: Score a draft before publishing. Scores a
// draft (a URL, HTML or Markdown) 0–100 against the pages that rank for its
// keyword: term coverage, outline coverage, questions answered, length and
// structure — with the missing terms (and how often to use them), missing
// sections, unanswered questions and a fix list, most important first —
// factConflicts, sentences stating a price, limit or date that contradicts the
// client's registered facts; uncitedClaims, the percentages, money and large
// numbers, 'according to X', 'by 2030' and '3x' figures with no link, footnote
// or source line within 200 characters (uncitedRatio, a warning above 0.4; a
// heuristic); invisibleChars, zero-width, bidi and tag characters by
// codepoint; draftQuality, which only detects and never rewrites: 26 filler
// and 46 stock AI phrases scored per 1,000 tokens, information density, bigram
// repetition, flags (filler 50+, AI phrasing 40+, density under 0.20,
// repetition 30+, under 300 tokens) and a weighted overall; and placement, the
// keyword in the first 100 words, the URL slug and the front of the title,
// heading stuffing (more than two subheadings repeating it), average sentence
// length against 15–20 words, internal links per 1,000 words against 3–5, and
// a table of contents on pages of 1,500 words or more. Answers 'is this ready
// to publish' and 'what should I add'. Run it after each revision until it is
// at par with the results page. Same cost as get_serp_content_profile; a
// results page bought for the brief is reused free for a day.
func (c *Client) ScoreContent(ctx context.Context, in ScoreContentInput) (json.RawMessage, error) {
	return c.Call(ctx, "score_content", in)
}

// ScoreEeatInput is the input to score_eeat.
type ScoreEeatInput struct {
	ClientID      string   `json:"clientId"`
	URLs          []string `json:"urls,omitempty"`
	TopPages      *int64   `json:"topPages,omitempty"`
	Judge         *bool    `json:"judge,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// ScoreEeat calls score_eeat: Score pages for E-E-A-T. Scores pages of the
// client's site (urls, up to 20, or the topPages by Search Console clicks) for
// Experience (20), Expertise (25), Authoritativeness (25) and Trust (30), each
// from sub-criteria graded Strong, Moderate, Weak or None with the evidence
// and the fix: first-hand method and results, figures, original images, a
// byline and author page, Person markup with jobTitle and sameAs, credentials,
// citations to outside and primary sources and the share of uncited claims,
// visible dates, about, contact, privacy and terms pages, HTTPS, reviews, a
// returns policy on shops, specialised vocabulary, awards and accreditations,
// 'featured in' media coverage, professional affiliations, a history of
// publishing on the topic (other crawled pages on it), visible corrections and
// update history, no deceptive practices (undisclosed affiliate links,
// clickbait headlines, false urgency), secure payments on shops, and for the
// site referring domains and brand mentions from our index and a Wikidata
// entity. Flags pages weak on who, how and why (atRisk), YMYL pages (health,
// finance, legal, news, civic, groups) with no credentials (Expertise capped
// at Weak), generic AI-style pages (one heading template, no author, stock
// phrasing), stale pages on fast-moving topics, and reading ease outside 60–70
// (information only). Bands at 90/70/50/30, each page and the sample with the
// next steps for its band. A heuristic. judge: true regrades Experience with
// one metered Claude call (quote it with estimate_cost). Answers 'does our
// content show E-E-A-T' and 'which pages look untrustworthy'. Free without
// judge.
func (c *Client) ScoreEeat(ctx context.Context, in ScoreEeatInput) (json.RawMessage, error) {
	return c.Call(ctx, "score_eeat", in)
}

// SearchLocalBusinessesInput is the input to search_local_businesses.
type SearchLocalBusinessesInput struct {
	ClientID      string   `json:"clientId"`
	Lat           float64  `json:"lat"`
	Lng           float64  `json:"lng"`
	RadiusKm      *float64 `json:"radiusKm,omitempty"`
	Categories    []string `json:"categories,omitempty"`
	Name          string   `json:"name,omitempty"`
	MinRating     *float64 `json:"minRating,omitempty"`
	Claimed       *bool    `json:"claimed,omitempty"`
	Limit         *int64   `json:"limit,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// SearchLocalBusinesses calls search_local_businesses: Businesses in an area.
// Every Google Maps business within a radius of a point, by category and/or
// name, with rating, reviews, claimed status, photos, site and ids — up to
// 1,000. Answers 'who are all the plumbers within 10 km' and 'which rivals
// have unclaimed listings'. Category names are Google's (dentist, plumber,
// pizza_restaurant); get_local_categories lists those in an area. A paid
// lookup on the client's account unless anyone read the same thing recently;
// quote it with estimate_cost.
func (c *Client) SearchLocalBusinesses(ctx context.Context, in SearchLocalBusinessesInput) (json.RawMessage, error) {
	return c.Call(ctx, "search_local_businesses", in)
}

// SearchLocationsInput is the input to search_locations.
type SearchLocationsInput struct {
	ClientID string `json:"clientId"`
	Query    string `json:"query"`
	Country  string `json:"country,omitempty"` // ISO country code: us, gb, de…
	Type     string `json:"type,omitempty"`    // Only this kind: City, County, State, Region, DMA Region, Postal Code…
	Limit    *int64 `json:"limit,omitempty"`
}

// SearchLocations calls search_locations: Find a location to track from. Looks
// up search locations in a country by name — cities, counties, regions, DMAs,
// postal areas — with the location code create_rank_tracker and the research
// tools take. Answers 'track us from Manchester, not all of the UK'. Free.
func (c *Client) SearchLocations(ctx context.Context, in SearchLocationsInput) (json.RawMessage, error) {
	return c.Call(ctx, "search_locations", in)
}

// SearchWebContentInput is the input to search_web_content.
type SearchWebContentInput struct {
	ClientID      string   `json:"clientId"`
	Keyword       string   `json:"keyword"`
	Exact         *bool    `json:"exact,omitempty"`   // Match the phrase as written.
	InTitle       *bool    `json:"inTitle,omitempty"` // Only pages with it in the title.
	PageTypes     []string `json:"pageTypes,omitempty"`
	Sentiment     string   `json:"sentiment,omitempty"`
	Since         string   `json:"since,omitempty"` // Only pages published on or after this date.
	MinDomainRank *int64   `json:"minDomainRank,omitempty"`
	ExcludeDomain string   `json:"excludeDomain,omitempty"` // Leave this domain out, usually the client's own.
	OnePerDomain  *bool    `json:"onePerDomain,omitempty"`
	Sort          string   `json:"sort,omitempty"`
	Summary       *bool    `json:"summary,omitempty"` // Also read the summary of the whole set (a second lookup).
	Limit         *int64   `json:"limit,omitempty"`
	Offset        *int64   `json:"offset,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// SearchWebContent calls search_web_content: Pages on the web that mention a
// phrase. Content Explorer over DataForSEO's content index: pages that mention
// a word or phrase, each with its domain, title, snippet, publish date, domain
// rank (0–1000), sentiment (positive, negative, neutral) and page type,
// filtered by page type, sentiment, date and domain rank, and a summary of the
// whole set — sentiment split, top domains, page types. Answers 'who writes
// about this', 'who mentions our brand without linking', 'what is said about
// this topic, and how warmly' and 'which blogs cover this niche'. A paid
// lookup on the client's account unless anyone read the same thing recently
// (spend says which); quote it with estimate_cost.
func (c *Client) SearchWebContent(ctx context.Context, in SearchWebContentInput) (json.RawMessage, error) {
	return c.Call(ctx, "search_web_content", in)
}

// SendOutreachInput is the input to send_outreach.
type SendOutreachInput struct {
	ClientID   string           `json:"clientId"`
	MailboxID  string           `json:"mailboxId"`
	SequenceID string           `json:"sequenceId,omitempty"`
	Prospects  []map[string]any `json:"prospects"`
	DryRun     *bool            `json:"dryRun,omitempty"` // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
}

// SendOutreach calls send_outreach: Send outreach emails from a connected
// mailbox. Sends the first email to each prospect from the chosen mailbox now,
// and schedules the sequence's follow-ups (when a sequenceId is given), which
// the hourly sweep sends unless a reply arrives first. Each prospect takes its
// own subject and body (from draft_outreach_emails, say) or the sequence's
// first step; the address is to, or the contact saved on the prospect. A
// prospect with no address, or already contacted, is skipped and listed. Sent
// prospects move to contacted. This sends real email from a person's mailbox:
// confirm the list and the wording with the user first — dryRun: true sends
// nothing and returns wouldSend, every email (address, subject, body) and
// follow-up (with when it would go) exactly as it would be sent. Answers 'send
// these outreach emails'. No vendor charge.
func (c *Client) SendOutreach(ctx context.Context, in SendOutreachInput) (json.RawMessage, error) {
	return c.Call(ctx, "send_outreach", in)
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

// SendTestEventInput is the input to send_test_event.
type SendTestEventInput struct {
	WebhookID string `json:"webhookId"`
}

// SendTestEvent calls send_test_event: Send a test event. Queues an
// alert.fired test event to one webhook, delivered within a minute, so a
// receiver's signature check can be tried end to end. Answers 'test my
// webhook'. Free.
func (c *Client) SendTestEvent(ctx context.Context, in SendTestEventInput) (json.RawMessage, error) {
	return c.Call(ctx, "send_test_event", in)
}

// SendWeeklyDigestNowInput is the input to send_weekly_digest_now.
type SendWeeklyDigestNowInput struct {
	ClientID string `json:"clientId"`
	Provider string `json:"provider"`
}

// SendWeeklyDigestNow calls send_weekly_digest_now: Post the digest now.
// Builds this week's digest and posts it to the client's connected Slack,
// Discord or Telegram straight away, outside any schedule. Answers 'send the
// digest to the channel now'.
func (c *Client) SendWeeklyDigestNow(ctx context.Context, in SendWeeklyDigestNowInput) (json.RawMessage, error) {
	return c.Call(ctx, "send_weekly_digest_now", in)
}

// SerpOverviewInput is the input to serp_overview.
type SerpOverviewInput struct {
	ClientID      string `json:"clientId"`
	Term          string `json:"term"`
	LocationCode  *int64 `json:"locationCode,omitempty"`
	LanguageCode  string `json:"languageCode,omitempty"`
	Refresh       *bool  `json:"refresh,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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
	ClientID  string   `json:"clientId"`
	Cadence   string   `json:"cadence,omitempty"`
	Engines   []string `json:"engines,omitempty"`
	Sentiment *bool    `json:"sentiment,omitempty"`
}

// SetAIVisibilitySchedule calls set_ai_visibility_schedule: Set how often AI
// visibility is checked. Sets a client's cadence (off, daily, weekly or
// monthly; weekly unless changed) which assistants it is checked on (engines;
// null or empty for every one the agency has configured), and whether each
// answer naming a tracked brand is also judged for sentiment (off by default;
// one metered Claude call per such answer, its ceiling in
// estimate_ai_visibility_cost as sentimentIfOnPerRunMaxMicros). Every
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

// SetBudgetInput is the input to set_budget.
type SetBudgetInput struct {
	ClientID   string   `json:"clientId"`
	MonthlyUsd *float64 `json:"monthlyUsd,omitempty"`
	AlertAt    *float64 `json:"alertAt,omitempty"`
}

// SetBudget calls set_budget: Set a client's monthly budget and alert. Sets a
// client's monthly spending budget in USD (null removes it; paid calls past it
// are refused) and the share of it, 0–1, at which a budget.threshold event and
// alert fire once that month (null turns the alert off). Leave either out to
// keep it. A per-call ceiling belongs to an API key: create_api_key's
// maxCostUsdPerCall. Answers 'cap this client at $50 a month and warn me at
// 80%'. Free.
func (c *Client) SetBudget(ctx context.Context, in SetBudgetInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_budget", in)
}

// SetClientFactsInput is the input to set_client_facts.
type SetClientFactsInput struct {
	ClientID string           `json:"clientId"`
	Facts    []map[string]any `json:"facts"`
	DryRun   *bool            `json:"dryRun,omitempty"` // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
}

// SetClientFacts calls set_client_facts: Record the client's canonical facts.
// Writes the client's facts — prices, plan limits, countries served, founding
// date, contact details — each with a key ("starter plan price"), its value
// ("$29/month"), a category and optionally the terms a sentence must contain
// to be about it; replaces a fact with the same key. dryRun: true saves
// nothing and lists wouldSet, each fact as it would be saved with the value it
// would replace. check_ai_brand_accuracy checks assistants against them first,
// and score_content and get_freshness_queue flag pages whose text contradicts
// one. Answers 'remember that our price is now $29'. Free.
func (c *Client) SetClientFacts(ctx context.Context, in SetClientFactsInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_client_facts", in)
}

// SetCrawlAuthInput is the input to set_crawl_auth.
type SetCrawlAuthInput struct {
	ClientID  string           `json:"clientId"`
	BasicAuth map[string]any   `json:"basicAuth,omitempty"`
	Headers   []map[string]any `json:"headers,omitempty"`
	Cookies   string           `json:"cookies,omitempty"`
	Clear     *bool            `json:"clear,omitempty"`
}

// SetCrawlAuth calls set_crawl_auth: Credentials for crawling a staging site.
// Lets the crawler into a staging or password-protected site: HTTP basic auth
// (username and password), custom request headers (e.g. a bypass token for a
// firewall, up to 10) and cookies ('name=value; other=value'). Replaces
// whatever was set; clear: true removes it all. Stored encrypted, sent only to
// the client's own domain (never to a redirect elsewhere or a third-party
// script), and never shown again — get_crawl_settings lists only the username
// and header and cookie names. Answers 'audit our staging site' and 'crawl
// behind the login'. Free.
func (c *Client) SetCrawlAuth(ctx context.Context, in SetCrawlAuthInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_crawl_auth", in)
}

// SetCrawlScheduleInput is the input to set_crawl_schedule.
type SetCrawlScheduleInput struct {
	ClientID string `json:"clientId"`
	Cadence  string `json:"cadence"`
	Day      *int64 `json:"day,omitempty"`
	Hour     *int64 `json:"hour,omitempty"`
}

// SetCrawlSchedule calls set_crawl_schedule: Set when the site is crawled.
// Sets how often the client's site is crawled without anyone asking: off,
// daily, weekly or monthly; weekly on Monday at 02:00 UTC unless changed. day
// is the weekday for weekly (0 Sunday to 6 Saturday; default 1, Monday) or the
// day of the month for monthly (1–28; default 1); hour is the UTC hour
// (default 2). Scheduled crawls use the client's crawl settings and wait their
// turn in the crawl queue like any other; one is never started while another
// crawl of the site is running, and a crawl started by hand since the last
// slot counts as that slot's. get_crawl_settings shows the schedule and when
// it next runs. Answers 'crawl this site every night'. Free.
func (c *Client) SetCrawlSchedule(ctx context.Context, in SetCrawlScheduleInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_crawl_schedule", in)
}

// SetCrawlSettingsInput is the input to set_crawl_settings.
type SetCrawlSettingsInput struct {
	ClientID          string           `json:"clientId"`
	StripParams       []string         `json:"stripParams,omitempty"`
	Extractions       []map[string]any `json:"extractions,omitempty"`
	Render            string           `json:"render,omitempty"`
	RenderMaxPages    *int64           `json:"renderMaxPages,omitempty"`
	Mode              string           `json:"mode,omitempty"`
	ListURLs          []string         `json:"listUrls,omitempty"`
	IncludePatterns   []string         `json:"includePatterns,omitempty"`
	ExcludePatterns   []string         `json:"excludePatterns,omitempty"`
	UserAgent         map[string]any   `json:"userAgent,omitempty"`
	RequestsPerSecond *float64         `json:"requestsPerSecond,omitempty"`
	Concurrency       *int64           `json:"concurrency,omitempty"`
	MaxDepth          *int64           `json:"maxDepth,omitempty"`
	MaxPages          *int64           `json:"maxPages,omitempty"`
}

// SetCrawlSettings calls set_crawl_settings: Change how the site is crawled.
// Changes the client's crawl settings for the next crawl; fields left out are
// kept. mode: spider (follow links from the home page and sitemap), list
// (crawl only listUrls, up to 5,000 addresses on the client's domain) or
// sitemap (crawl only the sitemap's URLs). includePatterns / excludePatterns:
// wildcards on the path and query ('/blog/*', '*?sort=*', '/*.pdf') or full
// addresses — with includes set, only matching URLs are crawled (the home page
// is still read to find them). maxPages (up to 20,000) and maxDepth (clicks
// from the start, up to 50). userAgent: preset seolyticize (default),
// googlebot-smartphone, googlebot-desktop, bingbot, chrome, or custom with
// your own string; robots.txt is obeyed as that crawler. requestsPerSecond
// (0.2–10, default 4) and concurrency (1–5); the crawler still backs off on
// 429s, 5xx and Retry-After. For staging sites' passwords, headers and cookies
// use set_crawl_auth. stripParams removes query parameters from every
// discovered URL (names, or prefixes like utm_*), so sorted and filtered
// copies of a page are crawled once. extractions reads values off every page
// (up to 10): css (a selector, optionally an attribute), xpath
// (//div[@class='author'], //meta[@property='og:type']/@content,
// (//a)[1]/@href — child and descendant steps, attribute and text tests,
// positions), regex (first group), jsonld ('Article.datePublished') or meta (a
// name or property) — see get_extraction_results. render runs JavaScript: off,
// shells (pages that look unrendered) or all, capped at renderMaxPages.
// Answers 'crawl this site this way'. Free.
func (c *Client) SetCrawlSettings(ctx context.Context, in SetCrawlSettingsInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_crawl_settings", in)
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

// SetKeywordOwnersInput is the input to set_keyword_owners.
type SetKeywordOwnersInput struct {
	ClientID string           `json:"clientId"`
	Owners   []map[string]any `json:"owners"`
	Source   string           `json:"source,omitempty"`
	DryRun   *bool            `json:"dryRun,omitempty"` // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
}

// SetKeywordOwners calls set_keyword_owners: Say which page owns a keyword.
// Registers the page that owns each keyword (or cluster: the keyword plus its
// cluster members) for the client, replacing an earlier claim on the same
// keyword. build_content_brief and draft_article then warn when a piece would
// compete with the owner, and get_cannibalization says which page should win.
// Accept propose_keyword_owners' proposals with source search_console. dryRun:
// true saves nothing and lists wouldSet, each claim as it would be saved with
// the claim it would replace. Answers 'make /pricing the page for "crm
// pricing"'. Free.
func (c *Client) SetKeywordOwners(ctx context.Context, in SetKeywordOwnersInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_keyword_owners", in)
}

// SetLinkProspectStatusInput is the input to set_link_prospect_status.
type SetLinkProspectStatusInput struct {
	ClientID     string `json:"clientId"`
	ProspectID   string `json:"prospectId"`
	Status       string `json:"status"`
	LinkURL      string `json:"linkUrl,omitempty"`
	ContactEmail string `json:"contactEmail,omitempty"`
}

// SetLinkProspectStatus calls set_link_prospect_status: Move a link prospect
// to another stage. Moves one prospect on the client's outreach list to
// identified, contacted, replied, won or rejected, and saves its contact
// email. For won, give linkUrl — the page the link sits on (a page prospect
// defaults to its own page): a link check is queued to confirm the link is
// live, and the weekly link check keeps verifying it (list_link_prospects
// shows the result). Take the id from list_link_prospects. Answers 'we heard
// back from this site' and 'we got the link'. Free.
func (c *Client) SetLinkProspectStatus(ctx context.Context, in SetLinkProspectStatusInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_link_prospect_status", in)
}

// SetPathOwnersInput is the input to set_path_owners.
type SetPathOwnersInput struct {
	ClientID string           `json:"clientId"`
	Owners   []map[string]any `json:"owners,omitempty"`
	Mode     string           `json:"mode,omitempty"`   // replace: owners is the whole list. merge: owners are added or updated, remove is dropped, the rest kept.
	Remove   []string         `json:"remove,omitempty"` // Prefixes to drop, with mode merge.
	DryRun   *bool            `json:"dryRun,omitempty"` // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
}

// SetPathOwners calls set_path_owners: Say who owns which part of the site.
// For a site built on several platforms, names the owner (and a contact) of
// each path prefix: /docs, /blog and /community on one vendor, the shop on
// another. mode replace (the default) makes owners the client's whole list,
// and an empty list clears it; mode merge adds or updates the prefixes given,
// drops those in remove, and keeps every other. The answer lists the resulting
// owners and the changes: added, updated (with what it was), removed and kept.
// dryRun: true saves nothing and answers the same, to check a change first. A
// URL belongs to the longest prefix it falls under, on segment boundaries
// (/docs covers /docs/x, not /docsearch); / is the catch-all. Then
// get_audit_summary and list_issue_pages split issues by owner,
// create_tasks_from_audit puts the owner on each task, and the digest groups
// its fixes by owner. Answers 'the docs team owns /docs, send those fixes to
// them'. Free.
func (c *Client) SetPathOwners(ctx context.Context, in SetPathOwnersInput) (json.RawMessage, error) {
	return c.Call(ctx, "set_path_owners", in)
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

// SimulateSerpSnippetInput is the input to simulate_serp_snippet.
type SimulateSerpSnippetInput struct {
	ClientID    string `json:"clientId,omitempty"`
	URL         string `json:"url,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// SimulateSerpSnippet calls simulate_serp_snippet: How a result will look.
// Shows a page's title and meta description as Google's results would,
// measured in pixels rather than characters: what is shown and where it is cut
// with an ellipsis on desktop and on mobile, the breadcrumb URL, and warnings
// (cut, too short to survive Google's rewriting, missing — with what Google
// falls back to). Give url to read the live page, and/or title and description
// to try new wording (given text overrides the page's). Widths are
// approximations of Google's layout. Answers 'will this title be truncated'
// and 'what will this look like in Google'. Free.
func (c *Client) SimulateSerpSnippet(ctx context.Context, in SimulateSerpSnippetInput) (json.RawMessage, error) {
	return c.Call(ctx, "simulate_serp_snippet", in)
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
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// SubmitBingSitemapInput is the input to submit_bing_sitemap.
type SubmitBingSitemapInput struct {
	ClientID   string `json:"clientId"`
	SitemapURL string `json:"sitemapUrl"`
}

// SubmitBingSitemap calls submit_bing_sitemap: Submit a sitemap to Bing.
// Submits a sitemap or feed (XML sitemap, RSS 2.0, Atom or a text list) on the
// client's site to Bing Webmaster Tools. Answers 'tell Bing about this
// sitemap'. list_sitemaps then shows it with Bing's reading. Refuses addresses
// not on the client's site.
func (c *Client) SubmitBingSitemap(ctx context.Context, in SubmitBingSitemapInput) (json.RawMessage, error) {
	return c.Call(ctx, "submit_bing_sitemap", in)
}

// SubmitBingURLsInput is the input to submit_bing_urls.
type SubmitBingURLsInput struct {
	ClientID      string   `json:"clientId"`
	URLs          []string `json:"urls"`
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// SubmitBingURLs calls submit_bing_urls: Submit pages to Bing. Submits
// addresses on the client's site to Bing through Webmaster Tools' URL
// submission (SubmitUrlBatch), after reading the site's daily quota: addresses
// beyond what is left today are returned as not sent rather than refused by
// Bing one by one. Answers 'get Bing to index these pages now'. Only addresses
// on the client's own site are sent. submit_indexnow reaches Bing too, without
// the quota, once the client's IndexNow key is published. dryRun: true reads
// the quota but sends nothing, returning wouldSubmit, refused, and the quote.
func (c *Client) SubmitBingURLs(ctx context.Context, in SubmitBingURLsInput) (json.RawMessage, error) {
	return c.Call(ctx, "submit_bing_urls", in)
}

// SubmitIndexnowInput is the input to submit_indexnow.
type SubmitIndexnowInput struct {
	ClientID      string   `json:"clientId"`
	URLs          []string `json:"urls"`
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// SubmitIndexnow calls submit_indexnow: Tell search engines a page changed.
// Submits changed addresses to IndexNow, which Bing, Yandex, Seznam and Naver
// share. Google does not take part. Only accepts addresses on a site the
// agency manages — anything else is refused and named. Answers 'we have just
// republished these pages, tell the engines'. Refuses entirely until the
// client's key file has been published and checked. dryRun: true sends nothing
// and returns wouldSubmit (empty when blocked says the real call would not
// send), refused, and the quote.
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
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
}

// SubmitURLs calls submit_urls: Submit URLs for crawling. Adds addresses to
// the crawl frontier for a client. Only accepts addresses on a site the agency
// manages — anything else is refused and named. dryRun: true adds nothing and
// returns wouldSubmit (the addresses as they would be queued), refused, and
// the quote. Answers 'crawl these pages I found'.
func (c *Client) SubmitURLs(ctx context.Context, in SubmitURLsInput) (json.RawMessage, error) {
	return c.Call(ctx, "submit_urls", in)
}

// SubmitYandexRecrawlInput is the input to submit_yandex_recrawl.
type SubmitYandexRecrawlInput struct {
	ClientID      string   `json:"clientId"`
	URLs          []string `json:"urls"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// SuggestAIVisibilitySetupInput is the input to suggest_ai_visibility_setup.
type SuggestAIVisibilitySetupInput struct {
	ClientID       string   `json:"clientId"`
	SeedKeywords   []string `json:"seedKeywords,omitempty"`
	MaxSeeds       *int64   `json:"maxSeeds,omitempty"`
	PromptsPerSeed *int64   `json:"promptsPerSeed,omitempty"`
	MaxCompetitors *int64   `json:"maxCompetitors,omitempty"`
	Platform       string   `json:"platform,omitempty"` // Which AI answer index the questions come from.
	Apply          *bool    `json:"apply,omitempty"`
	MaxCostMicros  *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun         *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// SuggestAIVisibilitySetup calls suggest_ai_visibility_setup: Set up AI
// visibility from the client's own data. Drafts a client's AI visibility
// setup: its own brand (from the client's name, site and brand terms),
// competitors (from the brief and rank tracker, then DataForSEO's organic
// competitors to fill up to maxCompetitors), and starter questions grouped by
// topic — real questions people ask AI assistants, from DataForSEO's AI answer
// index, around the seed keywords (given, or the tracked keywords with the
// most AI search volume), each with its AI search volume and fan-out searches,
// branded or neutral. Nothing is saved unless apply is true, which adds them
// all (topics included) as save_ai_visibility_setup would, skipping anything
// already tracked. Answers 'set up AI tracking for this client'. A paid lookup
// on the client's account unless anyone asked the same thing recently (served
// says which); quote it with estimate_cost. (up to one competitors read, one
// AI volume read and one question read per seed)
func (c *Client) SuggestAIVisibilitySetup(ctx context.Context, in SuggestAIVisibilitySetupInput) (json.RawMessage, error) {
	return c.Call(ctx, "suggest_ai_visibility_setup", in)
}

// SuggestInternalLinksInput is the input to suggest_internal_links.
type SuggestInternalLinksInput struct {
	ClientID  string           `json:"clientId"`
	CrawlID   string           `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	TargetURL string           `json:"targetUrl,omitempty"`
	Phrases   []string         `json:"phrases,omitempty"`
	Limit     *int64           `json:"limit,omitempty"`
	Format    string           `json:"format,omitempty"`
	Sources   []map[string]any `json:"sources,omitempty"` // The linking pages' own source, for diffs against what the CMS stores.
}

// SuggestInternalLinks calls suggest_internal_links: Suggest internal links.
// Pages that already mention a topic in their text but do not link to the page
// about it, with the exact sentence and the words to use as the anchor,
// strongest linking pages first. Give targetUrl (and phrases, default: its
// title and H1 and the queries it ranks for); without a target, it picks the
// client's pages in striking distance in Search Console (positions 4–20) —
// where an internal link moves rankings most. format 'patch' also returns, per
// page to edit, a unified diff of its HTML or Markdown source (from sources,
// else the page as served) with the links inserted, and for every link the
// CMS-ready edit: the words to link, the target and the sentence with the link
// in it. Answers 'where should we add internal links' and 'give me the edits
// to make'. Free.
func (c *Client) SuggestInternalLinks(ctx context.Context, in SuggestInternalLinksInput) (json.RawMessage, error) {
	return c.Call(ctx, "suggest_internal_links", in)
}

// SuggestMetaTagsInput is the input to suggest_meta_tags.
type SuggestMetaTagsInput struct {
	ClientID      string   `json:"clientId"`
	CrawlID       string   `json:"crawlId,omitempty"`
	URLs          []string `json:"urls,omitempty"`
	Problems      []string `json:"problems,omitempty"`
	Limit         *int64   `json:"limit,omitempty"`
	Format        string   `json:"format,omitempty"`
	MaxCostMicros *int64   `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool    `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// SuggestMetaTags calls suggest_meta_tags: Titles and descriptions for a whole
// crawl. Finds the crawled pages whose title or meta description is missing,
// duplicated, too long or too short (measured in pixels as Google draws them:
// a 600px title, a 920px description — not characters), and writes a new title
// and description for each, keeping the Search Console queries the page
// already earns clicks for. Every suggestion is pixel-checked again. Returns
// rows, and with format csv a spreadsheet of url, problems, current and
// suggested values and their widths. Answers 'fix the titles across the site'.
// Reads the latest complete crawl unless crawlId is given; urls narrows it. A
// paid Claude call per 25 pages.
func (c *Client) SuggestMetaTags(ctx context.Context, in SuggestMetaTagsInput) (json.RawMessage, error) {
	return c.Call(ctx, "suggest_meta_tags", in)
}

// SyncLocalReviewsInput is the input to sync_local_reviews.
type SyncLocalReviewsInput struct {
	ClientID      string `json:"clientId"`
	LocationID    string `json:"locationId"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
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

// SyncOutreachInput is the input to sync_outreach.
type SyncOutreachInput struct {
	ClientID string `json:"clientId"`
}

// SyncOutreach calls sync_outreach: Check outreach threads for replies, and
// send due follow-ups. Does now what the hourly sweep does for this client:
// reads every live outreach thread for a reply (moving a prospect who answered
// to replied and calling off its follow-ups), then sends the follow-ups that
// are due. Answers 'did anyone reply to our outreach'. No vendor charge; it
// sends due follow-ups, so it needs client:manage.
func (c *Client) SyncOutreach(ctx context.Context, in SyncOutreachInput) (json.RawMessage, error) {
	return c.Call(ctx, "sync_outreach", in)
}

// SyncSearchConsoleInput is the input to sync_search_console.
type SyncSearchConsoleInput struct {
	ClientID string `json:"clientId"`
}

// SyncSearchConsole calls sync_search_console: Collect Search Console data
// now. Collects the latest days from Search Console now instead of waiting for
// the nightly sync — and the sixteen-month history too if it is still owed.
// Runs in the background for a minute or two; read the data with the Search
// Console tools afterwards. Returns the previous sync's time and note: a note
// saying Google reported no search data means the property had no impressions
// on those days (a new or barely visible site), not a broken connection.
// Answers 'pull our Search Console data now'. Free.
func (c *Client) SyncSearchConsole(ctx context.Context, in SyncSearchConsoleInput) (json.RawMessage, error) {
	return c.Call(ctx, "sync_search_console", in)
}

// SyncTaskStatusesInput is the input to sync_task_statuses.
type SyncTaskStatusesInput struct {
	ClientID string `json:"clientId"`
}

// SyncTaskStatuses calls sync_task_statuses: Sync task status with the tracker
// now. Reads every mirrored issue's state now and brings each pair into step:
// an issue moved in the tracker moves its task, a task moved here moves its
// issue (GitHub has no in-progress, so an open issue stays open). The hourly
// sync does this on its own. Answers 'did they close the Jira tickets'. Free.
func (c *Client) SyncTaskStatuses(ctx context.Context, in SyncTaskStatusesInput) (json.RawMessage, error) {
	return c.Call(ctx, "sync_task_statuses", in)
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

// TagListKeywordsInput is the input to tag_list_keywords.
type TagListKeywordsInput struct {
	ListID     string   `json:"listId"`
	KeywordIDs []string `json:"keywordIds"`
	Tags       []string `json:"tags"`
	Remove     *bool    `json:"remove,omitempty"`
}

// TagListKeywords calls tag_list_keywords: Tag keywords on a list. Puts tags
// on keywords in a list, or takes them off with remove: true — by keywordId
// from query_keyword_list, or keywordIds: 'all' for every keyword on the list.
// Answers 'tag these as pillar' and 'take q3 off the whole list'. Tags are
// case-insensitive labels up to 40 characters; changed counts the keywords
// that gained or lost one. For a tracker's keywords use tag_keywords. Free.
func (c *Client) TagListKeywords(ctx context.Context, in TagListKeywordsInput) (json.RawMessage, error) {
	return c.Call(ctx, "tag_list_keywords", in)
}

// TestRobotsInput is the input to test_robots.
type TestRobotsInput struct {
	ClientID string   `json:"clientId"`
	URLs     []string `json:"urls"`
	Bots     []string `json:"bots,omitempty"`
}

// TestRobots calls test_robots: Can this bot fetch this page. Reads the
// client's live robots.txt and says, for each address and each crawler,
// whether it may fetch it and which line decides — for every search engine and
// AI crawler by default (Googlebot, Bingbot, Applebot, DuckDuckBot, YandexBot,
// Baiduspider, GPTBot, OAI-SearchBot, ChatGPT-User, ClaudeBot,
// Claude-SearchBot, Claude-User, PerplexityBot, Perplexity-User,
// Google-Extended, Applebot-Extended, Meta-ExternalAgent, Bytespider, CCBot,
// Amazonbot, DuckAssistBot, MistralAI-User, AhrefsBot, SemrushBot), or the
// tokens in bots. Same parser the crawler obeys. Answers 'can GPTBot read our
// pricing page', 'are we blocking Google-Extended' and 'which rule blocks
// Googlebot here'. Free.
func (c *Client) TestRobots(ctx context.Context, in TestRobotsInput) (json.RawMessage, error) {
	return c.Call(ctx, "test_robots", in)
}

// UnacceptIssueInput is the input to unaccept_issue.
type UnacceptIssueInput struct {
	ClientID string   `json:"clientId"`
	IDs      []string `json:"ids"`
}

// UnacceptIssue calls unaccept_issue: Stop accepting an issue. Removes
// acceptances by id (from list_accepted_issues), so their findings count again
// in the scorecard, audit summary, digest, tasks and alerts. Unknown ids are
// ignored. Answers 'start flagging those pages again'. Free.
func (c *Client) UnacceptIssue(ctx context.Context, in UnacceptIssueInput) (json.RawMessage, error) {
	return c.Call(ctx, "unaccept_issue", in)
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
	PaidSource      *bool    `json:"paidSource,omitempty"`
}

// UpdateAlertRule calls update_alert_rule: Update an alert rule. Changes a
// rule's threshold, cooldown or delivery (channel, webhookUrl — null removes
// it — emailRecipients, telegramChatId), or pauses and resumes it, or turns
// its paidSource on or off. Answers 'this alert is too noisy', 'send this to
// email instead' or 'mute this for now'. Pausing keeps the rule's history.
// paidSource (competitor-new-keywords and web-mentions only; off by default;
// needs spend permission) also buys the signal from DataForSEO on every sweep:
// for competitor-new-keywords, Labs ranked keywords each tracker competitor
// (up to 5) newly ranks for at the tracker's market, not only tracked
// keywords; for web-mentions, pages DataForSEO's content index first saw
// naming the brand, off the client's own domains. The sweep is hourly but a
// purchase is reused from the archive while fresh, so it buys at most once a
// week per competitor (once a day for mentions); it is capped by the client's
// monthly budget, and a failed or over-budget read falls back to the free
// signal. Quote it first with estimate_cost.
func (c *Client) UpdateAlertRule(ctx context.Context, in UpdateAlertRuleInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_alert_rule", in)
}

// UpdateBusinessProfileInput is the input to update_business_profile.
type UpdateBusinessProfileInput struct {
	ClientID   string         `json:"clientId"`
	LocationID string         `json:"locationId,omitempty"` // The client's location, from list_local_locations. Default: its one connected location.
	Changes    map[string]any `json:"changes"`
	DryRun     *bool          `json:"dryRun,omitempty"` // True to do nothing and get back exactly what the call would do: the rows it would write, the addresses it would submit, the emails it would send.
}

// UpdateBusinessProfile calls update_business_profile: Change a Google
// Business Profile listing. Changes the client's own Google Business Profile
// listing through its connected Business Profile (the connection review
// replies are posted with): regularHours (the whole week, 24-hour HH:MM,
// closeDay past midnight), specialHours (holidays and one-off dates, closed or
// open–close; needs regular hours on the listing), description (up to 750
// characters), phone (primary and up to two additional), website, categories
// (Google ids such as gcid:dentist; primary and up to nine additional) and
// simple yes/no attributes by Google id. Only the fields given are changed
// (locations.patch with an updateMask; attributes through updateAttributes);
// phone and category lists Google replaces whole keep the half not given from
// the listing. dryRun: true changes nothing and returns the exact patch and
// updateMask, the listing's fields before and after, and Google's own check of
// the change (validateOnly) — show it to the person first. Refused with
// missing_scope when the connection was not granted permission to manage the
// listing, and not_connected when the location has none; next says how to
// connect it. Google may hold a change for review before it shows. Answers
// 'update our opening hours' and 'change the business description'. Free.
func (c *Client) UpdateBusinessProfile(ctx context.Context, in UpdateBusinessProfileInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_business_profile", in)
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
	BrandTerms  []string         `json:"brandTerms,omitempty"` // Replaces the brand terms: what people type when already looking for this business (its name, products, misspellings). Search Console splits branded from non-bra
}

// UpdateClientBrief calls update_client_brief: Update the client brief.
// Changes the client brief that every teammate and agent reads. Answers
// 'remember this about the client'. Set fields (overview, goal, positioning,
// audience, voice; null or empty clears one, fields left out are untouched);
// add entries (kind competitor, key_page, section or research, each with a
// title and optional body and url); edit entries by id; remove entries by id
// (unknown ids are ignored); replace the brand terms (brandTerms) that split
// branded from non-branded search. Applied in that order, so removing an old
// competitor and adding its replacement in one call works. Record a research
// entry when you finish a piece of work whose conclusions the next one should
// start from; the 50 newest research notes are kept. Returns the brief as it
// now stands. Free.
func (c *Client) UpdateClientBrief(ctx context.Context, in UpdateClientBriefInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_client_brief", in)
}

// UpdateContentPlanItemInput is the input to update_content_plan_item.
type UpdateContentPlanItemInput struct {
	ClientID  string `json:"clientId"`
	ID        string `json:"id"`
	Month     string `json:"month,omitempty"`
	PublishOn string `json:"publishOn,omitempty"`
	Status    string `json:"status,omitempty"`
	TargetURL string `json:"targetUrl,omitempty"`
	Notes     string `json:"notes,omitempty"`
	Keyword   string `json:"keyword,omitempty"`
}

// UpdateContentPlanItem calls update_content_plan_item: Change a planned
// piece. Moves a planned piece to another month or date, changes its status
// (planned, briefed, drafting, published, dropped), target URL, keyword or
// notes. Answers 'mark that article published' and 'push this to May'. Free.
func (c *Client) UpdateContentPlanItem(ctx context.Context, in UpdateContentPlanItemInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_content_plan_item", in)
}

// UpdateDataSyncInput is the input to update_data_sync.
type UpdateDataSyncInput struct {
	ClientID   string         `json:"clientId"`
	SyncID     string         `json:"syncId"`
	Name       string         `json:"name,omitempty"`
	Query      map[string]any `json:"query,omitempty"` // The same query query_data takes, without clientId: see describe_datasets.
	Mode       string         `json:"mode,omitempty"`
	WindowDays *int64         `json:"windowDays,omitempty"`
	Cadence    string         `json:"cadence,omitempty"`
	RunHour    *int64         `json:"runHour,omitempty"`
	Paused     *bool          `json:"paused,omitempty"`
}

// UpdateDataSync calls update_data_sync: Change a data sync. Changes a sync's
// name, query, mode, window, cadence or hour, or pauses or resumes it; only
// what is passed changes. Answers 'make the sheet weekly' and 'pause the
// BigQuery export'. Free.
func (c *Client) UpdateDataSync(ctx context.Context, in UpdateDataSyncInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_data_sync", in)
}

// UpdateLeadWidgetInput is the input to update_lead_widget.
type UpdateLeadWidgetInput struct {
	WidgetID       string   `json:"widgetId"`
	Name           string   `json:"name,omitempty"`
	Heading        string   `json:"heading,omitempty"`
	AccentColor    string   `json:"accentColor,omitempty"`
	AllowedOrigins []string `json:"allowedOrigins,omitempty"` // Sites allowed to show the form, as origins (https://agency.example). Empty allows any site, for testing.
	NotifyEmails   []string `json:"notifyEmails,omitempty"`   // Who is emailed when a lead arrives. Empty keeps leads without telling anyone.
	Enabled        *bool    `json:"enabled,omitempty"`
}

// UpdateLeadWidget calls update_lead_widget: Change a lead-gen widget. Renames
// a widget, changes its heading, colour, allowed sites or who is notified, or
// switches it off and on (enabled). A switched-off form stops taking requests
// at once; its leads are kept. Answers 'turn off the form on our old site'.
// Free.
func (c *Client) UpdateLeadWidget(ctx context.Context, in UpdateLeadWidgetInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_lead_widget", in)
}

// UpdatePortfolioInput is the input to update_portfolio.
type UpdatePortfolioInput struct {
	PortfolioID string           `json:"portfolioId"`
	Name        string           `json:"name,omitempty"`
	Description string           `json:"description,omitempty"`
	Add         []map[string]any `json:"add,omitempty"`
	Remove      []map[string]any `json:"remove,omitempty"`
}

// UpdatePortfolio calls update_portfolio: Change a portfolio. Renames or
// re-describes a portfolio, and adds or removes members — a whole client, or
// one site with its client. Renaming needs access to every client in it.
// Adding a member already there is a no-op. Answers 'add this client to the
// retail portfolio'.
func (c *Client) UpdatePortfolio(ctx context.Context, in UpdatePortfolioInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_portfolio", in)
}

// UpdateRankTrackerInput is the input to update_rank_tracker.
type UpdateRankTrackerInput struct {
	ClientID  string `json:"clientId"`
	TrackerID string `json:"trackerId"`
	Name      string `json:"name,omitempty"`
	Device    string `json:"device,omitempty"`
	Frequency string `json:"frequency,omitempty"`
	RunHour   *int64 `json:"runHour,omitempty"` // The UTC hour scheduled runs happen at (default 3).
	RunDay    *int64 `json:"runDay,omitempty"`  // Weekly: the weekday, 0 Sunday – 6 Saturday. Monthly: the day of the month, 1–28.
	Depth     *int64 `json:"depth,omitempty"`   // Results read per keyword (default 100). Deeper costs more per check.
	Match     string `json:"match,omitempty"`   // domain: the exact host; subdomains: it and every subdomain; prefix: only addresses under urlPrefix (a folder or one page).
	URLPrefix string `json:"urlPrefix,omitempty"`
}

// UpdateRankTracker calls update_rank_tracker: Update a rank tracker. Renames
// a tracker or changes its device (desktop, mobile or both), cadence (daily,
// weekly, monthly, or manual — only when refresh_rankings asks), the hour and
// day it runs, how deep it reads (10–200 results), or what counts as the
// client (the exact host, every subdomain, or only a folder or page via
// urlPrefix). Answers 'check these weekly instead', 'track only our /blog/
// pages'. Weekly costs a seventh as much, monthly a thirtieth.
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

// UpdateTaskInput is the input to update_task.
type UpdateTaskInput struct {
	ClientID       string         `json:"clientId"`
	TaskID         string         `json:"taskId"`
	Title          string         `json:"title,omitempty"`
	Notes          string         `json:"notes,omitempty"`
	Status         string         `json:"status,omitempty"`
	AssigneeUserID string         `json:"assigneeUserId,omitempty"` // A user id from list_team_members (the member's userId). They must be able to open this client.
	DueDate        string         `json:"dueDate,omitempty"`        // YYYY-MM-DD.
	Link           map[string]any `json:"link,omitempty"`           // What the task is about: an audit issue (value: its ruleId from list_crawl_issues), a page (an absolute URL) or a keyword.
}

// UpdateTask calls update_task: Change a task. Changes a task's title, notes,
// status, assignee, due date or link; only what is passed changes, and null
// clears. Moving to done or wont_do records when; moving back to todo clears
// it. Answers 'mark the title fixes done' and 'give this to Sam for next
// week'.
func (c *Client) UpdateTask(ctx context.Context, in UpdateTaskInput) (json.RawMessage, error) {
	return c.Call(ctx, "update_task", in)
}

// ValidateProductSchemaInput is the input to validate_product_schema.
type ValidateProductSchemaInput struct {
	ClientID string   `json:"clientId"`
	URLs     []string `json:"urls,omitempty"`
	Limit    *int64   `json:"limit,omitempty"`
	Eu       *bool    `json:"eu,omitempty"` // Sold in the EU: require the energy efficiency class EPREL asks for.
	Render   string   `json:"render,omitempty"`
}

// ValidateProductSchema calls validate_product_schema: Check Product and Offer
// structured data. Reads product pages and checks their Product and Offer
// JSON-LD against what Google's merchant listings need — errors: name, image,
// offers as Offer (not AggregateOffer) with a price above 0 and an ISO 4217
// priceCurrency (or a priceSpecification), a valid availability; warnings:
// brand (its name not empty or N/A), a GTIN or MPN or SKU, description,
// offers.url, itemCondition, priceValidUntil when it has passed, validFrom
// without an end date, shipping details with shippingDestination and
// deliveryTime, a return policy with applicableCountry and
// returnPolicyCategory, a rating with ratingValue and reviewCount,
// hasAdultConsideration values, a ProductGroup's productGroupID, variesBy and
// hasVariant, and with eu: true the energy efficiency class; notes:
// returnPolicyCountry, member programs, variants — says whether each page is
// eligible, and gives a completenessScore (required fields 50, aggregateRating
// 65, an identifier 75, shippingDetails 85, a return policy 90, three or more
// reviews 100). With no urls, the latest crawl's pages that carry Product
// markup (up to limit). render 'auto' renders a page in a browser when its
// HTML carries no Product markup, for stores that inject their JSON-LD with
// JavaScript. Answers 'is our product schema right' and 'why don't our
// products show prices in Google'. Free.
func (c *Client) ValidateProductSchema(ctx context.Context, in ValidateProductSchemaInput) (json.RawMessage, error) {
	return c.Call(ctx, "validate_product_schema", in)
}

// ValidateSitemapInput is the input to validate_sitemap.
type ValidateSitemapInput struct {
	ClientID   string `json:"clientId"`
	CrawlID    string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	CheckFiles *bool  `json:"checkFiles,omitempty"`
}

// ValidateSitemap calls validate_sitemap: Is the sitemap telling the truth.
// Checks the sitemaps a crawl read against the pages it found: entries that do
// not answer 200, are noindexed or are canonicalised elsewhere; indexable
// pages no sitemap lists (the sitemap's coverage of the crawl) and listed
// pages the crawl never reached (the crawl's coverage of the sitemap); lastmod
// accuracy against the previous crawl (content changed but lastmod did not,
// lastmod moved but nothing changed, dates in the future, every entry stamped
// with the build date); image and video sitemap coverage; and the entries
// added, removed and re-dated since the previous crawl. With checkFiles
// (default on) it also reads each sitemap file live, up to 20 — robots.txt's
// Sitemap: lines, the files the crawl read and what an index names — and
// judges each on its own: over 50,000 URLs or 50 MB uncompressed (critical),
// XML errors with their line number, not declared in robots.txt, http://
// entries, <priority>/<changefreq> (ignored by Google), over 1,000 image:image
// per URL and the deprecated image tags, video entries missing thumbnail_loc,
// title, description or a content/player address, and news sitemaps over 1,000
// entries, older than two days or missing publication name, language, date or
// title. Answers 'is our sitemap clean', 'does Google trust our lastmod' and
// 'which pages are missing from the sitemap'. Free.
func (c *Client) ValidateSitemap(ctx context.Context, in ValidateSitemapInput) (json.RawMessage, error) {
	return c.Call(ctx, "validate_sitemap", in)
}

// ValidateStructuredDataInput is the input to validate_structured_data.
type ValidateStructuredDataInput struct {
	ClientID string `json:"clientId"`
	URL      string `json:"url,omitempty"`
	CrawlID  string `json:"crawlId,omitempty"` // A crawl id from list_crawls, or "latest" for the newest complete one.
	Limit    *int64 `json:"limit,omitempty"`
}

// ValidateStructuredData calls validate_structured_data: Structured data
// against Google's requirements. Checks JSON-LD against what Google's search
// gallery requires and recommends for each rich result type — Article,
// Breadcrumb, Course list (Course + ItemList), Carousel (ItemList), Dataset,
// Discussion forum, Event, Image metadata, Job posting, Local business, Movie,
// Organization, Product (snippets and merchant listings), Profile page, Q&A,
// Recipe, Review and rating, Software app, Vacation rental, Video, Site name —
// plus FAQ (retired for every site on 7 May 2026), HowTo, the 2025–2026
// retirements (special announcement, fact check, vehicle listing, estimated
// salary, practice problem), Service and SoftwareSourceCode, which earn no
// rich result any more (or never did). Values Google cannot read (non-ISO
// dates and durations, prices with currency signs, relative URLs that are not
// URLs) are errors. Give url to read one of the client's pages live, or a
// crawlId (default latest) for the whole site by type: pages, how many are
// valid, and the properties most often missing. Answers 'is our product markup
// eligible for rich results' and 'what structured data is broken'. Free.
func (c *Client) ValidateStructuredData(ctx context.Context, in ValidateStructuredDataInput) (json.RawMessage, error) {
	return c.Call(ctx, "validate_structured_data", in)
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

// WaitForJobInput is the input to wait_for_job.
type WaitForJobInput struct {
	ClientID       string `json:"clientId"`
	JobID          string `json:"jobId"`
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`
}

// WaitForJob calls wait_for_job: Wait for a slow job. get_job that waits:
// holds the call open until the job finishes or timeoutSeconds pass (at most
// 60), then answers with its status, so an agent makes one call instead of
// polling. Call again while it says running. Answers 'wait for the crawl'.
// Free.
func (c *Client) WaitForJob(ctx context.Context, in WaitForJobInput) (json.RawMessage, error) {
	return c.Call(ctx, "wait_for_job", in)
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

// YoutubeKeywordResearchInput is the input to youtube_keyword_research.
type YoutubeKeywordResearchInput struct {
	ClientID      string `json:"clientId"`
	Seed          string `json:"seed"`
	Measure       *int64 `json:"measure,omitempty"`
	LocationCode  *int64 `json:"locationCode,omitempty"`
	LanguageCode  string `json:"languageCode,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// YoutubeKeywordResearch calls youtube_keyword_research: YouTube search
// phrases around a topic. What people type into YouTube around a seed, in
// YouTube's own order of popularity (its suggestions; YouTube publishes no
// search volume, so none is given), and for the first measure phrases how
// crowded the results are: the top videos' median views, how many are from the
// last year, how many channels share them and how many are Shorts. Answers
// 'what should our next video target' and 'which YouTube phrases are
// winnable'. Suggestions are free; each phrase measured is one YouTube results
// page: A paid lookup on the client's account unless anyone read the same
// thing today (spend says which); quote it with estimate_cost.
func (c *Client) YoutubeKeywordResearch(ctx context.Context, in YoutubeKeywordResearchInput) (json.RawMessage, error) {
	return c.Call(ctx, "youtube_keyword_research", in)
}

// YoutubeVideoBriefInput is the input to youtube_video_brief.
type YoutubeVideoBriefInput struct {
	ClientID      string `json:"clientId"`
	Keyword       string `json:"keyword"`
	Videos        *int64 `json:"videos,omitempty"`
	LocationCode  *int64 `json:"locationCode,omitempty"`
	LanguageCode  string `json:"languageCode,omitempty"`
	MaxCostMicros *int64 `json:"maxCostMicros,omitempty"` // Optional spending limit for this call, in micro-USD (1,000,000 = $1). A purchase that would take the call past it is refused before anything is bought. Pass the
	DryRun        *bool  `json:"dryRun,omitempty"`        // True to buy nothing and get the quote instead: what the call would cost at most and what it would buy.
}

// YoutubeVideoBrief calls youtube_video_brief: A title, description and
// chapter brief from the top videos. Reads the videos ranking on YouTube for a
// phrase (the results page, then each video's details) and writes the brief
// for a video to beat them: how their titles are built (length, the phrase
// first, numbers, brackets, the year) with examples, how long their
// descriptions run and how many use chapters and hashtags, the chapter labels
// they share in their usual order, the tags they use most, and how long the
// videos are. Answers 'how should we title and structure this video'. One
// results page plus one read per video: A paid lookup on the client's account
// unless anyone read the same thing today (spend says which); quote it with
// estimate_cost.
func (c *Client) YoutubeVideoBrief(ctx context.Context, in YoutubeVideoBriefInput) (json.RawMessage, error) {
	return c.Call(ctx, "youtube_video_brief", in)
}
