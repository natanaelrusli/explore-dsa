## Car Toll Cost Maker

**Description**

Given three different types of total cost. These for the details:

- Small Car: 1000
- Medium Car: 2000
- Big Car: 3000

As a user, you can input your car type and tap your e-money. The system will show you the cost, you car type and remaining balance of your e-money.

### Flow Chart

Start → Input car type → Input e-money balance → Check if car type is valid → If not valid, display "Invalid car type entered" → If valid, check if balance is sufficient → If not sufficient, display "Insufficient balance" → If sufficient, calculate remaining balance → Display cost, car type, and remaining balance → End

### Pseudo Code

```jsx
FUNCTION calculateTollCost(carType, balance)
    SWITCH carType
        CASE "Small":
            cost = 1000
        CASE "Medium":
            cost = 2000
        CASE "Big":
            cost = 3000
        DEFAULT:
            RETURN "Invalid car type entered."
    END SWITCH

    IF balance < cost THEN
        RETURN "Insufficient balance."
    ELSE
        remainingBalance = balance - cost
        RETURN "Cost: " + cost + " Car type: " + carType + " Remaining balance: " + remainingBalance
    END IF
END FUNCTION

BEGIN MAIN
    INPUT carType
    INPUT balance
    PRINT calculateTollCost(carType, balance)
END MAIN
```
