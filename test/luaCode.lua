

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



print("Hello World!")
print("sum: " , testval ,  sum(testval))
print("Prime: ", testval , isPrime(testval))