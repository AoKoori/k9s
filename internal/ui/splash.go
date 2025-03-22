1| // SPDX-License-Identifier: Apache-2.0
2| // Copyright Authors of K9s
3| 
4| package ui
5| 
6| import (
7| 	"fmt"
8| 	"io"
9| 	"log/slog"
10| 	"net/http"
11| 	"strings"
12| 
13| 	"github.com/derailed/k9s/internal/config"
14| 	"github.com/derailed/tview"
15| )
16| 
17| // LogoSmall K9s small log.
18| var LogoSmall []string
19| 
20| // Removed the unnecessary blank line
21| 
22| // LogoBig K9s big logo for splash page.
23| var LogoBig = []string{
24| 	` ____  __ ________        _______  ____     ___ `,
25| 	`|    |/  /   __   \______/   ___ \|    |   |   |`,
26| 	`|       /\____    /  ___/    \  \/|    |   |   |`,
27| 	`|    \   \  /    /\___  \     \___|    |___|   |`,
28| 	`|____|\__ \/____//____  /\______  /_______ \___|`,
29| 	`         \/           \/        \/        \/    `,
30| }
31| 
32| // Splash represents a splash screen.
33| type Splash struct {
34| 	*tview.Flex
35| }
36| 
37| // NewSplash instantiates a new splash screen with product and company info.
38| func NewSplash(styles *config.Styles, version string) *Splash {
39| 	s := Splash{Flex: tview.NewFlex()}
40| 	s.SetBackgroundColor(styles.BgColor())
41| 
42| 	logo := tview.NewTextView()
43| 	logo.SetDynamicColors(true)
44| 	logo.SetTextAlign(tview.AlignCenter)
45| 	s.layoutLogo(logo, styles)
46| 
47| 	vers := tview.NewTextView()
48| 	vers.SetDynamicColors(true)
49| 	vers.SetTextAlign(tview.AlignCenter)
50| 	s.layoutRev(vers, version, styles)
51| 
52| 	s.SetDirection(tview.FlexRow)
53| 	s.AddItem(logo, 10, 1, false)
54| 	s.AddItem(vers, 1, 1, false)
55| 
56| 	return &s
57| }
58| 
59| func (s *Splash) layoutLogo(t *tview.TextView, styles *config.Styles) {
60| 	logo := strings.Join(LogoBig, fmt.Sprintf("\n[%s::b]", styles.Body().LogoColor))
61| 	fmt.Fprintf(t, "%s[%s::b]%s\n",
62| 		strings.Repeat("\n", 2),
63| 		styles.Body().LogoColor,
64| 		logo)
65| }
66| 
67| func (s *Splash) layoutRev(t *tview.TextView, rev string, styles *config.Styles) {
68| 	fmt.Fprintf(t, "[%s::b]Revision [red::b]%s", styles.Body().FgColor, rev)
69| }
70| 
71| // function to get the logo []string from the LogoUrl
72| // by making a request to the LogoUrl
73| func GetLogo(logoUrl string) {
74| 	slog.Debug("Fetching logo from URL", "url", logoUrl)
75|     defaultLogo := []string{
76|         ` ____  __ ________       `,
77|         `|    |/  /   __   \______`,
78|         `|       /\____    /  ___/`,
79|         `|    \   \  /    /\___  \`,
80|         `|____|\__ \/____//____  /`,
81|         `         \/           \/ `,
82|     }
83| 
84|     if logoUrl == "" {
85|         LogoSmall = defaultLogo
86|         return
87|     }
88| 
89|     resp, err := http.Get(logoUrl)
90|     if err != nil {
91|         slog.Error("Error fetching logo from URL", "url", logoUrl, "error", err)
92|         LogoSmall = defaultLogo
93|         return
94|     }
95|     defer func() {
96|         if resp != nil {
97|             resp.Body.Close()
98|         }
99|     }()
100| 
101|     if resp.StatusCode != http.StatusOK {
102|         slog.Error("Non-OK HTTP status", "status", resp.Status)
103|         LogoSmall = defaultLogo
104|         return
105|     }
106| 
107|     body, err := io.ReadAll(resp.Body)
108|     if err != nil {
109|         slog.Error("Error reading response body", "error", err)
110|         LogoSmall = defaultLogo
111|         return
112|     }
113| 
114|     logo := strings.Split(string(body), "\n")
115|     // last line is empty, remove it
116|     if len(logo) > 0 && logo[len(logo)-1] == "" {
117|         logo = logo[:len(logo)-1]
118|     }
119|     slog.Debug("Successfully fetched logo from URL", "url", logoUrl)
120|     LogoSmall = logo
121| }
122| 
