// webpack.config.js
const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

module.exports = {
    entry: {
        main: './js/index.js',
        styles: './css/main.css',
    },
    output: {
        filename: 'js/[name].bundle.js',
        path: path.resolve(__dirname, 'www'),
        clean: true, // Nettoie le répertoire de sortie avant chaque build
    },
    module: {
        rules: [
            {
                test: /\.css$/i,
                use: [
                    // Extrait le CSS dans des fichiers séparés
                    MiniCssExtractPlugin.loader,
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
            filename: 'css/[name].css',
        }),
        // new CopyWebpackPlugin({
        //     patterns: [
        //         {
        //             from: '**/*.html', // Source des fichiers HTML
        //             to: '[path][name][ext]', // Destination dans 'www'
        //             context: 'templates/', // Contexte pour conserver la structure des dossiers
        //         },
        //     ],
        // })
    ],
    mode: 'development', // Changez en 'production' pour des builds optimisés
    devtool: 'source-map', // Utile pour le debugging
};
