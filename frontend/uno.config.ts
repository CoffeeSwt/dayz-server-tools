import {
    defineConfig,
    presetAttributify,
    presetIcons,
    presetTypography,
    presetWebFonts,
    presetWind3,
    transformerDirectives,
    transformerVariantGroup
} from 'unocss'

export default defineConfig({
    shortcuts: [
        // ...
    ],
    theme: {
        colors: {
            // ...
        }
    },
    presets: [
        presetWind3(),
        presetAttributify(),
        presetIcons(),
        presetTypography(),
        presetWebFonts({
            fonts: {
                // ...
            },
        }),
    ],
    transformers: [
        transformerDirectives(),
        transformerVariantGroup(),
    ],
    safelist: [
        ...Array.from({ length: 128 }, (_, i) => `p-${i + 1}`),
        ...Array.from({ length: 128 }, (_, i) => `w-${i + 1}`),
    ],
})