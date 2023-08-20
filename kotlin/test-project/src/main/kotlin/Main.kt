import syntax.MyData

fun main(args: Array<String>) {
    val myData = MyData(10);
    myData.test1();
    myData.name = "Miyako";

    ifExpression(true)
    ifExpression(false)
}

fun ifExpression(condition: Boolean) {
    val value: String = if (condition) "Hi, true" else "Hey, false"
    println(value)
}
