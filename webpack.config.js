const path = require("path");
const HtmlWebpackPlugin = require('html-webpack-plugin');
// const CopyWebpackPlugin = require('copy-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const OptimizeCssAssetsPlugin = require('optimize-css-assets-webpack-plugin');
const TerserPlugin = require('terser-webpack-plugin');

module.exports = {
  mode: 'production',
  entry: {
    vendor: './app/data/assets/vendor.webpack.js',
    main: './app/data/assets/main.webpack.js',
  },
  output: {
    publicPath: '/assets/dist/',
		path: path.resolve(__dirname, "./app/data/assets/dist/"),
    filename: '[name].[chunkhash].js'
  },
  optimization: {
    minimizer: [
      new OptimizeCssAssetsPlugin(),
      new TerserPlugin()
    ]
  },
  module: {
    rules: [{
      test: /\.scss$/,
      use: [MiniCssExtractPlugin.loader, 'css-loader', 'sass-loader']
    },{
      test: /.(png|jpg|jpeg|gif|svg|woff|woff2|ttf|eot)$/,
      loader: 'url-loader?limit=10000',
    }]
	},
	plugins: [
    new HtmlWebpackPlugin({
			inject: "head",
			template: './app/views/template/base.html',
      filename: "../../base.html",
      minify: {

      }
		}),
    new MiniCssExtractPlugin({filename: "[name].[contentHash].css"}),
    new CleanWebpackPlugin()
	]
};
