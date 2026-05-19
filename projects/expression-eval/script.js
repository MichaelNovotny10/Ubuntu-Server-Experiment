const expressionInput = document.getElementById('expression-field');
const evaluateBtn = document.getElementById('evaluate-btn');

function runEvaluation() {
    const expression = expressionInput.value;
    if (!expression.trim()) {
        alert("Please enter an expression first!")
        return;
    }
    try {
        const result = math.evaluate(expression);
        alert("Result: " + result);
    } catch (error) {
        alert("Invalid math Expression!")
    }
}

evaluateBtn.addEventListener('click', runEvaluation)

expressionInput.addEventListener('keydown', function(event) {
    if (event.key == 'Enter') {
        runEvaluation();
    }
});
