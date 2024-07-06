const path = require('path')

module.exports = {
    entry: ['./client/index.ts'],
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
        path: path.resolve(__dirname, 'client/static'),
        filename: 'bundle.js'
    },
}