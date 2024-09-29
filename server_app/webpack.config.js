// webpack.config.js
const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

module.exports = {
    entry: {
        main: './js/index.js',
        styles: './css/styles.css',
    },
    output: {
        filename: 'js/[name].bundle.js',
        path: path.resolve(__dirname, 'www'),
        clean: true,
    },
    module: {
        rules: [
            {
                test: /\.css$/i,
                use: [
                    MiniCssExtractPlugin.loader,
                    // 'style-loader',
                    'css-loader',
                    'postcss-loader',
                ],
            },
            {
                test: /\.m?js$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['@babel/preset-env'],
                    },
                },
            }
        ],
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: './css/[name].css',
        }),
        new CopyWebpackPlugin({
            patterns: [
                {
                    from: '**/*.html',
                    to: '[path][name][ext]',
                    context: 'layout/',
                },
            ],
        }),
        new CopyWebpackPlugin({
            patterns: [
                {
                    from: '**/*.html',
                    to: 'pages/[path][name][ext]',
                    context: 'pages/',
                },
            ],
        })
    ],
    mode: 'development',
    devtool: 'source-map',
};
