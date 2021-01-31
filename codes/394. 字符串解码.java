/*
    比较明显要用栈来解，搞清楚计算方式就能解出来
 */
class Solution {
    static class TokenParse{
        String token;
        String s;
        TokenParse(String token, String s) {
            this.token = token;
            this.s = s;
        }
        TokenParse() {
            token = null;
            s = null;
        }
    }

    TokenParse readToken(String s) {
        if(s.length()==0) return null;
        String token = null;
        if(s.startsWith("[")) {
            token = "[";
            return new TokenParse(token, s.substring(1));
        }
        else if (s.startsWith("]")) {
            token = "]";
            return new TokenParse(token, s.substring(1));
        }
        else if(Character.isAlphabetic(s.charAt(0))) {
            for(int i=0; i<s.length(); i++) {
                if(!Character.isAlphabetic(s.charAt(i))) {
                    token = s.substring(0, i);
                    return new TokenParse(token, s.substring(i));
                }
            }
        }
        else {
            for(int i=0; i<s.length(); i++) {
                if(!Character.isDigit(s.charAt(i))) {
                    token = s.substring(0, i);
                    return new TokenParse(token, s.substring(i));
                }
            }
        }
        return new TokenParse(s, "");
    }

    public String decodeString(String s) {
        ArrayList<String> data = new ArrayList<>();
        while (s.length() > 0) {
            TokenParse tmp = readToken(s);
            String token = tmp.token;
            s = tmp.s;
            if(token.startsWith("[")) {
                continue;
            }
            else if (token.startsWith("]")) {
                String right = data.remove(data.size()-1);
                String left = data.remove(data.size()-1);
                while (Character.isAlphabetic(left.charAt(0))) {
                    right = left + right;
                    left = data.remove(data.size()-1);
                }
                right = right.repeat(Integer.valueOf(left));
                data.add(right);
            }
            else {
                data.add(token);
            }
        }
        String result = new String("");
        for (String ss : data) {
            result += ss;
        }
        return result;
    }
}