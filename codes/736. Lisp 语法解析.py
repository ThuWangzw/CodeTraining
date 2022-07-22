'''
Exp = Num | Var | ( Let ) | ( Add ) | ( Mult )
Let = let + (V + Exp)* + Exp
Add = add + Exp + Exp
Mult = mult + Exp + Exp
'''
import enum
class TokenType(enum.Enum):
    EXP = 1
    NUM = 2
    VAR = 3
    LEFTB = 4
    RIGHTB = 5
    LET = 6
    ADD = 7
    MULT = 8
    END = 9


class Tokenizer:
    def __init__(self, s) -> None:
        self.s = s
        self.tokens = []
        self.parse()

    def parse(self):
        while True:
            token, tokenType = self._next()
            self.tokens.append((token, tokenType))
            if tokenType is TokenType.END:
                return

    def next(self):
        token, tokenType = self.tokens[0]
        self.tokens = self.tokens[1:]
        return token, tokenType
    
    def look(self, i=0):
        return self.tokens[i]

    def _next(self):
        for i in range(len(self.s)):
            if self.s[i] != ' ':
                self.s = self.s[i:]
                break
        if len(self.s)==0:
            return "", TokenType.END
        if self.s[0] == '(':
            top = self.s[0]
            self.s = self.s[1:]
            return str(top), TokenType.LEFTB
        if self.s[0] == ')':
            top = self.s[0]
            self.s = self.s[1:]
            return str(top), TokenType.RIGHTB
        idx = len(self.s)
        for i in range(len(self.s)):
            if self.s[i] == ' ' or self.s[i] == '(' or self.s[i] == ')':
                idx = i
                break
        res = self.s[:idx]
        self.s = self.s[idx:]
        if res == 'let':
            return res, TokenType.LET
        elif res == 'add':
            return res, TokenType.ADD
        elif res == 'mult':
            return res, TokenType.MULT
        elif res[0] >= 'a' and res[0] <= 'z':
            return res, TokenType.VAR
        else:
            return int(res), TokenType.NUM

class Parser:
    def __init__(self, s) -> None:
        self.s = s
        self.tokenizer = Tokenizer(s)
        self.vars = {}

    def Exp(self) -> int:
        token, tokenType = self.tokenizer.look()
        if tokenType == TokenType.NUM:
            val = self.Num()
            return val
        elif tokenType == TokenType.VAR:
            val = self.Var()
            return self.vars[val][len(self.vars[val])-1]
        elif tokenType == TokenType.LEFTB:
            self.tokenizer.next()
            val = 0
            token, tokenType = self.tokenizer.look()
            if tokenType == TokenType.LET:
                val = self.Let()
            elif tokenType == TokenType.ADD:
                val = self.Add()
            elif tokenType == TokenType.MULT:
                val = self.Mult()
            self.tokenizer.next()
            return val


    def Num(self) -> int:
        token, tokenType = self.tokenizer.next()
        return token

    def Var(self) -> str:
        token, tokenType = self.tokenizer.next()
        return token

    def Let(self) -> int:
        token, tokenType = self.tokenizer.next()
        varstack = []
        while True:
            token, tokenType = self.tokenizer.look()
            _, tokeType2 = self.tokenizer.look(1)
            if tokenType != TokenType.VAR or tokeType2 == TokenType.RIGHTB:
                val = self.Exp()
                while len(varstack)>0:
                    var = varstack[0]
                    varstack = varstack[1:]
                    self.vars[var].pop()
                return val
            var = self.Var()
            val = self.Exp()
            if var not in self.vars:
                self.vars[var] = []
            self.vars[var].append(val)
            varstack.append(var)

    def Add(self) -> int:
        self.tokenizer.next()
        a = self.Exp()
        b = self.Exp()
        return a+b

    def Mult(self) -> int:
        self.tokenizer.next()
        a = self.Exp()
        b = self.Exp()
        return a*b

class Solution:
    def evaluate(self, expression: str) -> int:
        parser = Parser(expression)
        return parser.Exp()