from pygments import lex
from pygments.lexers import GoLexer
from pygments.token import Token
import os

def all_operators_parsing(code):

    tokens = lex(code, GoLexer())
    count = 0

    operator_values = {
        "+", "-", "*", "/", "%", "&", "|", "^", "<<", ">>", "&^",
        "==", "!=", "<", "<=", ">", ">=", "&&", "||", "!", "=",
        "+=", "-=", "*=", "/=", "%=", "<<=", ">>=", "&=", "|=", "^=", "&^=",
        "<-", "if", "for", "range", "goto", "break", "continue", 
        "fallthrough", "return", "defer", "go", "case"
    }
    
    else_flag = False

    for token_type, value in tokens: 
        if value == "else":
            else_flag = True
            continue
        if else_flag and (value != " " or token_type != Token.Text.Whitespace):
            else_flag = False
            if value == "if":
                count += 1
                continue
        
        if token_type in Token.Operator or value in operator_values:
            count += 1
    
    return count

# file_path = os.path.abspath("Go/main.go")

# with open(file_path, "r", encoding="utf-8") as f:
#     code = f.read()

# print(all_operators_parsing(code))
