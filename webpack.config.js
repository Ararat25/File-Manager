const path = require('path')

module.exports = {
    resolve: {
        extensions: ['.js', '.ts']
    },
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