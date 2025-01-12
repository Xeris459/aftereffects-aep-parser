package aep

import (
	"testing"

	aep "github.com/xeris459/aftereffects-aep-parser/src"
)

func TestLayerMetadata(t *testing.T) {
	project, err := aep.Open("data/Layer-01.aep")
	if err != nil {
		t.Fatal(err)
	}

	comp01 := project.RootFolder.FolderContents[0]
	expect(t, comp01.CompositionLayers[0].CollapseTransformEnabled)
	expect(t, comp01.CompositionLayers[1].EffectsEnabled)
	expect(t, comp01.CompositionLayers[2].MotionBlurEnabled)
	expect(t, comp01.CompositionLayers[4].ShyEnabled)
	expect(t, comp01.CompositionLayers[5].AdjustmentLayerEnabled)
	expect(t, comp01.CompositionLayers[6].ThreeDEnabled)
	expect(t, comp01.CompositionLayers[7].SoloEnabled)
	expect(t, comp01.CompositionLayers[8].GuideEnabled)
	expect(t, comp01.CompositionLayers[9].FrameBlendMode, aep.LayerFrameBlendModePixelMotion)
	expect(t, comp01.CompositionLayers[10].FrameBlendMode, aep.LayerFrameBlendModeFrameMix)
	expect(t, comp01.CompositionLayers[11].Quality, aep.LayerQualityWireframe)
	expect(t, comp01.CompositionLayers[12].Quality, aep.LayerQualityDraft)
	expect(t, comp01.CompositionLayers[13].Quality, aep.LayerQualityBest)
	expect(t, comp01.CompositionLayers[14].SamplingMode, aep.LayerSamplingModeBilinear)
	expect(t, comp01.CompositionLayers[15].SamplingMode, aep.LayerSamplingModeBicubic)
	expect(t, comp01.CompositionLayers[16].VideoEnabled)
	expect(t, comp01.CompositionLayers[16].AudioEnabled)
}
