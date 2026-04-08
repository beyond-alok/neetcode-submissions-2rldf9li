class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        operators = {"+":0,"-":1,"*":2,"/":3}
        sum = None

        for i,v in enumerate(tokens):

            if v in operators:
                if operators[v] == 0:
                    b = stack.pop()
                    a = stack.pop()
                    stack.append(a + b)

                elif operators[v] == 1:
                    b = stack.pop()
                    a = stack.pop()
                    stack.append(a - b)


                elif operators[v] == 2:
                    b = stack.pop()
                    a = stack.pop()
                    stack.append(a * b)

                else:
                    b = stack.pop()
                    a = stack.pop()
                    stack.append(int(a / b))
                
            else :
                v = int(v)
                stack.append(v)
        return stack[0]