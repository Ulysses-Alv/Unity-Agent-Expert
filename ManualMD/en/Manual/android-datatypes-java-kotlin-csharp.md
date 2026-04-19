# Supported data types for Java/Kotlin and C# code

When accessing Java/Kotlin code from a C# script in Unity, you can pass certain data types. The following table outlines the data types you can send between Java/Kotlin code and C# script, and the ones that you cannot.

| **Supported data types** | **Unsupported data types** |
| --- | --- |
| * byte, int, short, long * float, double * bool * char * string * Arrays of any of these supported types * Objects/structs that are serialized/deserialized with JSON | * Objects that aren’t serialized/deserialized with JSON * Language specific collections, such as [List](https://developer.android.com/reference/java/util/List), [Set](https://developer.android.com/reference/java/util/Set), [Map](https://developer.android.com/reference/java/util/Map) * [Runnable](https://developer.android.com/reference/java/lang/Runnable) * Unity-specific data types, such as ([Vector](../ScriptReference/Vector3.md), [GameObject](../ScriptReference/GameObject.md), [UI](UIToolkits.md), [Transform](../ScriptReference/Transform.md)) |

## Code examples to transfer data between Java and C# code

The following code examples demonstrate how you can send all the supported data types between Java and C# code. You can save these Java and C# code files anywhere in your Unity project.

### Example 1: Transfer basic data types

The following C# script uses `AndroidJavaClass` to call the Java functions to transfer variable data.

```
using UnityEngine;

public class TransferTypes : MonoBehaviour
{
    void Start()
    {
        using (AndroidJavaClass pluginClass = new AndroidJavaClass("com.unity3d.player.JavaTransferTypesExample"))
        {
            //Receive data from Java functions
            Debug.Log($"Received from Java Byte: " + pluginClass.CallStatic<sbyte>("getByte"));
            Debug.Log($"Received from Java Short: " + pluginClass.CallStatic<short>("getShort"));
            Debug.Log($"Received from Java Int: " + pluginClass.CallStatic<int>("getInt"));
            Debug.Log($"Received from Java Long: " + pluginClass.CallStatic<long>("getLong"));
            Debug.Log($"Received from Java Float: " + pluginClass.CallStatic<float>("getFloat"));
            Debug.Log($"Received from Java Double: " + pluginClass.CallStatic<double>("getDouble"));
            Debug.Log($"Received from Java Boolean: " + pluginClass.CallStatic<bool>("getBoolean"));
            Debug.Log($"Received from Java Char: " + pluginClass.CallStatic<char>("getChar"));
            Debug.Log($"Received from Java String: " + pluginClass.CallStatic<string>("getString"));

            //Send data to Java functions
            pluginClass.CallStatic("setByte", (sbyte)9);
            pluginClass.CallStatic("setShort", (short)87);
            pluginClass.CallStatic("setInt", 6);
            pluginClass.CallStatic("setLong", (long)5);
            pluginClass.CallStatic("setFloat", 4.32f);
            pluginClass.CallStatic("setDouble", 9.987);
            pluginClass.CallStatic("setBoolean", false);
            pluginClass.CallStatic("setChar", 'b');
            pluginClass.CallStatic("setString", "This string is sent from C#");
        }
    }
}
```

The following Java code provides functions that the C# script calls to transfer data.

```
package com.unity3d.player;

public class JavaTransferTypesExample {

    // Java functions to send data
    public static byte getByte() {
        return 12;
    }
    public static short getShort() {
        return 34;
    }
    public static int getInt() {
        return 5;
    }
    public static long getLong() {
        return 6;
    }
    public static float getFloat() {
        return 7.89f;
    }
    public static double getDouble() {
        return 1.123;
    }
    public static boolean getBoolean() {
        return true;
    }
    public static char getChar() {
        return 'a';
    }
    public static String getString() {
        return "This string is sent from Java";
    }

    //Java functions to receive data
    public static void setByte(byte value) {
        android.util.Log.i("Example", "Received from C# Byte: " + value);
    }
    public static void setShort(short value) {
        android.util.Log.i("Example", "Received from C# Short: " + value);
    }
    public static void setInt(int value) {
        android.util.Log.i("Example", "Received from C# Int: " + value);
    }
    public static void setLong(long value) {
        android.util.Log.i("Example", "Received from C# Long: " + value);
    }
    public static void setFloat(float value) {
        android.util.Log.i("Example", "Received from C# Float: " + value);
    }
    public static void setDouble(double value) {
        android.util.Log.i("Example", "Received from C# Double: " + value);
    }
    public static void setBoolean(boolean value) {
        android.util.Log.i("Example", "Received from C# Boolean: " + value);
    }
    public static void setChar(char value) {
        android.util.Log.i("Example", "Received from C# Char: " + value);
    }
    public static void setString(String value) {
        android.util.Log.i("Example", "Received from C# String: " + value);
    }
}
```

### Example 2: Transfer complex data types

The following C# script calls the Java functions to transfer complex data types, such as class/object.

```
using UnityEngine;
using System;

public class PassComplexStructure : MonoBehaviour
{
    [Serializable]
    public class DataTypes
    {
        public int number;
        public float floatNumber;
        public bool boolValue;
        public string phrase;
    }
    
    void Start()
    {
        var example = new DataTypes
        {
            number = 123,
            floatNumber = 4.56f,
            boolValue = true,
            phrase = "Testing passing an object from C# - successful"
        };

        using var pluginClass = new AndroidJavaClass("com.unity3d.player.JavaPassObject");

        var jsonString = JsonUtility.ToJson(example);
        pluginClass.CallStatic("setDataSet", jsonString);

        var jsonFromJava = pluginClass.CallStatic<string>("getDataSet");
        var receivedData = JsonUtility.FromJson<DataTypes>(jsonFromJava);

        Debug.Log($@"C# Deserialized from Java JSON: Number: {receivedData.number}, Float: {receivedData.floatNumber}, Bool: {receivedData.boolValue}, String: {receivedData.phrase}");
    }
}
```

The following Java code provides functions that the C# script calls.

```
package com.unity3d.player;

import org.json.JSONException;
import org.json.JSONObject;

public class JavaPassObject {
   
    public static class DataTypes
    {
        public int number;
        public float floatNumber;
        public boolean boolValue;
        public String phrase;

        public DataTypes(int number, float floatNumber, boolean boolValue, String phrase) {
            this.number = number;
            this.floatNumber = floatNumber;
            this.boolValue = boolValue;
            this.phrase = phrase;
        }

        @Override
        public String toString() {
            return "Java Deserialized from C# JSON:" +
               " Number: " + number +
               ", Float: " + floatNumber +
               ", Bool: " + boolValue +
               ", String: " + phrase;
            }
    }
    
    public static void setDataSet(String value)
    {
        try {
            var jsonObject = new JSONObject(value);
           
            int number = (int) jsonObject.getInt("number");
            float floatNumber = (float) jsonObject.getDouble("floatNumber");
            boolean boolValue = (boolean) jsonObject.getBoolean("boolValue");
            String phrase = (String) jsonObject.getString("phrase");

            var receivedData = new DataTypes(number, floatNumber, boolValue, phrase);
            android.util.Log.i("Test", receivedData.toString());

        } catch (JSONException e) {
            System.err.println("Error parsing JSON: " + e.getMessage());
        }
    }

    public static String getDataSet() throws JSONException
    {
        var example = new DataTypes(987, 6.54f, false, "Testing passing an object from Java - successful");
       
        var jsonObject = new JSONObject();
        jsonObject.put("number", example.number);
        jsonObject.put("floatNumber", example.floatNumber);
        jsonObject.put("boolValue", example.boolValue);
        jsonObject.put("phrase", example.phrase);

        return jsonObject.toString();
    }
}
```

## Additional resources

* [JSON Serialization](json-serialization.md)