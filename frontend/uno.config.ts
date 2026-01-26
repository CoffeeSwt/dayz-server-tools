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
    rules: [
        [
            'draggable',
            {
                '--wails-draggable': 'drag;',
            },
        ],
        [
            'aspect-ratio-1k',
            {
                'aspect-ratio': '16 / 9',
            },
        ]
    ],
    shortcuts: [
        {
            'size-full': 'w-full h-full',
            'flex-center': 'flex items-center justify-center',
            'absolute-center': 'absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2',
        },
    ],
    theme: {
        colors: {
            'light-white-1': '#6c6c70',
            'light-white-2': '#f7f9fc',
            'light-white-3': '#f0f3f6',
            'light-black-1': '#283248',
            'light-black-2': '#e4e8ec',
            'dark-black-1': '#13131a',
            'dark-black-2': '#1a1a21',
            'dark-gray-1': '#a9a9ab',
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
        //@ts-ignore
        ...Array.from({ length: 128 }, (_, i) => `p-${i + 1}`),
        //@ts-ignore
        ...Array.from({ length: 128 }, (_, i) => `w-${i + 1}`),
        'i-material-symbols:light-mode',
        'i-material-symbols:dark-mode',
    ],
})