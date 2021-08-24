const path = require('path');
const dotenv = require('dotenv-webpack');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

const srcPath = 'template/';

module.exports = {
    // モード値を production に設定すると最適化された状態で、
    // development に設定するとソースマップ有効でJSファイルが出力される
    mode: 'production', // "production" | "development" | "none"

    // メインとなるJavaScriptファイル（エントリーポイント）
    entry: {
        index: './template/ts/index.ts',
        dashboard: './template/ts/dashboard.tsx',
        local_dashboard: './template/ts/local-dashboard.tsx',
        develop: './template/ts/develop.tsx',
    },

    output: {
        path: path.join(__dirname, srcPath, "js"),
        filename: "[name].js"
    },

    module: {
        rules: [{
            // 拡張子 .ts の場合
            test: /\.(ts|tsx)$/,
            // TypeScript をコンパイルする
            use: 'ts-loader'
        },
        {
            test: /\.scss$/i,
            use: [
                {
                    loader: MiniCssExtractPlugin.loader,
                },
                {
                    loader: 'css-loader',
                    options: {
                        url: false,
                    }
                },
                {
                    loader: 'sass-loader',
                    options: {
                        sassOptions: {
                            outputStyle: 'expanded',
                        },
                    },
                }
            ]
        },
        ]
    },
    // import 文で .ts ファイルを解決するため
    resolve: {
        modules: [
            "node_modules", // node_modules 内も対象とする
        ],
        extensions: [
            '.ts',
            '.js',// node_modulesのライブラリ読み込みに必要
            '.tsx',
        ]
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: `../css/[name].css`,
            ignoreOrder: true,
        }),
        new dotenv()
    ]
};