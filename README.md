# 🚀 LEET-YOUR-CODE

A powerful CLI tool that generates clean, well-commented Go solutions for LeetCode problems using AI. Get instant, production-ready code with detailed explanations for any LeetCode problem!

## ✨ Features

- 🤖 **AI-Powered Solutions**: Uses OpenRouter API to generate optimal Go solutions
- 📝 **Detailed Comments**: Each line of code comes with clear explanations
- 📁 **Organized Output**: Creates separate directories and files for each problem
- ⚡ **Fast & Efficient**: Get solutions in seconds
- 🎯 **CLI-First**: Simple command-line interface for developers

## 🛠️ Prerequisites

Before you begin, ensure you have:

- **Go 1.19+** installed on your system
- **OpenRouter API Key** (free tier available)
- **Git** for cloning the repository

## 📦 Installation

### 1. Clone the Repository
```bash
git clone https://github.com/Morizz00/LEET-YOUR-CODE.git
cd LEET-YOUR-CODE
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Set Up Environment Variables
Create a `.env` file in the project root:
```bash
touch .env
```

Add your OpenRouter API key to the `.env` file:
```env
OPENROUTER_API_KEY=your_api_key_here
```

**🔑 Getting Your API Key:**
1. Visit [OpenRouter.ai](https://openrouter.ai/)
2. Sign up for a free account
3. Navigate to API Keys section
4. Generate a new API key
5. Copy and paste it into your `.env` file

## 🚀 Usage

### Starting the Tool
From the project root directory, run:
```bash
go run cmd/main.go
```

You should see:
```
API KEY LOADED: sk-or-v1-****
Agent ready to roll! Enter Command With 'dsa'
-->
```

### Basic Commands

#### Generate Solution
```bash
dsa <problem-name>
```

**Examples:**
```bash
# Classic problems
dsa Two Sum
dsa Valid Parentheses
dsa Merge Two Sorted Lists

# Complex problems
dsa Best Time to Buy and Sell Stock
dsa Longest Substring Without Repeating Characters
dsa Container With Most Water

# Problems with numbers
dsa 3Sum
dsa 4Sum
dsa Add Two Numbers
```

#### Exit the Tool
- Press `Ctrl+C` or close the terminal
- Type any non-dsa command to get prompted again

### 📁 Output Structure

When you run `dsa Two Sum`, the tool creates:
```
📁 two_sum/
  └── 📄 two_sum.go
```

The generated file contains:
- Complete Go solution
- Detailed line-by-line comments
- Time and space complexity explanations
- Example usage

### 💡 Example Session

```bash
$ go run cmd/main.go
API KEY LOADED: sk-or-v1-****
Agent ready to roll! Enter Command With 'dsa'
-->dsa Two Sum
Processing...
Problem File created for: Two Sum

-->dsa Valid Parentheses  
Processing...
Problem File created for: Valid Parentheses

-->dsa Binary Search
Processing...
Problem File created for: Binary Search
```

After this session, you'll have:
```
📁 two_sum/
  └── 📄 two_sum.go
📁 valid_parentheses/
  └── 📄 valid_parentheses.go  
📁 binary_search/
  └── 📄 binary_search.go
```

## 🎯 Supported Problem Types

The tool can generate solutions for various LeetCode categories:

- **🔢 Array & String Problems**: Two Sum, 3Sum, Valid Anagram
- **🔗 Linked Lists**: Merge Lists, Reverse Linked List, Cycle Detection
- **🌳 Trees & Graphs**: Binary Tree Traversal, Path Sum, DFS/BFS
- **📚 Stack & Queue**: Valid Parentheses, Min Stack, Sliding Window
- **🔍 Search & Sort**: Binary Search, Quick Sort, Merge Sort
- **💡 Dynamic Programming**: Fibonacci, Climbing Stairs, Coin Change
- **🧮 Math & Logic**: Palindrome Number, Roman Numerals

## ⚠️ Important Notes

### Input Format
- Always start commands with `dsa`
- Use exact LeetCode problem names for best results
- Spaces in problem names are supported

### Error Handling
- **Invalid commands**: Tool will prompt "Invalidated, start over lil bro"
- **API errors**: Check your internet connection and API key
- **No content**: Try rephrasing the problem name

### Rate Limits
- Free OpenRouter tier has usage limits
- Consider upgrading for heavy usage
- Tool includes automatic error handling for rate limits

## 🔧 Troubleshooting

### Common Issues

**"Error loading .env file"**
```bash
# Make sure you're running from project root
cd path/to/LEET-YOUR-CODE
go run cmd/main.go
```

**"API key not found"**
- Check if `.env` file exists in project root
- Verify API key is correctly formatted
- Ensure no extra spaces or quotes around the key

**"No content returned from OpenRouter"**
- Check internet connection
- Verify API key is valid and has credits
- Try a simpler problem name

**"Error fetching"**
- The AI model might not understand the problem name
- Try using the exact LeetCode problem title
- Check if the problem exists on LeetCode

### Getting Help
- Check the problem name spelling
- Use official LeetCode problem titles
- Try variations like "Two Sum" vs "2Sum"

## 🛡️ Development

### Project Structure
```
LEET-YOUR-CODE/
├── cmd/
│   └── main.go          # Main application entry point
├── scrapper/
│   └── agent.go         # OpenRouter API client
├── problem/
│   └── problem.go       # File creation utilities
├── .env                 # Environment variables
├── go.mod              # Go modules
├── go.sum              # Dependency checksums
└── README.md           # This file
```

### Building from Source
```bash
# Build binary
go build -o leet-helper cmd/main.go

# Run binary
./leet-helper
```

### Contributing
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🌟 Support

If you find this tool helpful:
- ⭐ Star the repository
- 🐛 Report bugs in Issues
- 💡 Suggest features
- 🤝 Contribute code

## 📞 Contact

- **GitHub**: [@Morizz00](https://github.com/Morizz00)
- **Issues**: [Report here](https://github.com/Morizz00/LEET-YOUR-CODE/issues)

---

**Happy Coding! 🎉**

*Turn any LeetCode problem into clean, commented Go code in seconds!*
