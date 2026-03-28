package ascii

import "testing"

// TestRenderWithDifferentBanners checks if the Render function
// works correctly with different banner files.
func TestRenderWithDifferentBanners(t *testing.T) {

	// Define test cases with different banner file paths
	tests := []struct {
		name       string
		bannerPath string
	}{
		{"Standard Banner", "../../banners/standard.txt"},
		{"Shadow Banner", "../../banners/shadow.txt"},
		{"Thinkertoy Banner", "../../banners/thinkertoy.txt"},
	}

	// Loop through each test case
	for _, tt := range tests {

		// Load the banner file
		banner, err := LoadBanner(tt.bannerPath)
		if err != nil {
			// Stop the test if the banner cannot be loaded
			t.Fatalf("%s: failed to load banner: %v", tt.name, err)
		}

		// Render the text "Hi" using the loaded banner
		output := Render("Hi", banner)

		// Check that the render output is not empty
		if output == "" {
			t.Errorf("%s: Render returned empty string", tt.name)
		}
	}
}