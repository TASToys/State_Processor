

local testval = 153

function sum (input)
    local sum = 0
    for i = 0, input, 1 do
        sum = sum + i
    end
    return sum
end

function isPrime (input)
    if( input < 2 )then
        return false
    end
    
    for i = 2, input/2, 1 do
        if((input % i) == 0)then
            return false
        end
    end
    
    return true
    
end


local function fib(n)
    if n < 2 then return n end
    return fib(n - 2) + fib(n - 1)
end

--console.log(fib(35));




print("Hello World!")
print("sum: " , testval ,  sum(testval))
print("Prime: ", testval , isPrime(testval))

print("Fib 20: ", fib(20))
