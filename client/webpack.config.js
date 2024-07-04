const path = require('path')

module.exports = {
    entry: ['./index.ts'],
    module: {
        rules: [
            {
                test: /\.ts$/,
                use: 'ts-loader',
            },
            {
                test: /\.css$/,
                use: [ 'style-loader', 'css-loader' ],
            },
        ]
    },
    output: {
        path: path.resolve(__dirname, 'static'),
        filename: 'bundle.js'
    },
}