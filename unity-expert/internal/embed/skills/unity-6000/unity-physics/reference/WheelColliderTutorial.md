# Create a car with Wheel colliders

This tutorial takes you through the process of creating a basic functioning four-wheeled car with the PhysX [Wheel collider](class-WheelCollider.md).

The steps include:

1. [Prepare the project](#prepare-project)
2. [Prepare the car model](#prepare-model)
3. [Add the Wheel colliders](#add-wheel-colliders), and set their position and radius
4. [Add scripts to control car behavior](#add-scripts), including car movement and wheel rotation
5. [Test the car](#test-car)

## Prepare the project

To follow these instructions, you need:

* A new empty scene
* The demo car model (download: [car.fbx](../uploads/Main/car.fbx)), imported into the open project
* A ground surface for the car to travel across.

To create a ground surface:

1. Create a plane **GameObject** and rename it `Ground`. This plane is the ground that the car drives on.
2. Set its position to (0,0,0), and its scale to (100,1,100).

## Prepare the car model

First, place the Car model in the **Scene**:

1. In the **Project window**, select the imported `Car` asset.
2. Extract the car model’s textures so that Unity can display them in the Scene. To do this, navigate to the [Import Settings](class-FBXImporter.md) in the **Inspector** window and do the following:
   1. Select **Materials**
   2. Select **Textures** > **Extract Textures**. In the file explorer that opens, save them in the `Assets` folder (the default location).
   3. Select **Apply**.
3. Place the Car asset into the Scene, on or just above the Ground plane.

Take a look at the Car GameObject’s hierarchy in the Hierarchy window. There is a root GameObject called `Car`, and child GameObjects for the car body model and each wheel model.

Configure the car body for **collision**:

1. On the `Car` root GameObject, add a [Rigidbody](class-Rigidbody.md) component.
2. Set the Rigidbody’s **Mass** to `1500``. This value defines the weight of the car in kilograms. It is a suitable weight for the Wheel collider’s default suspension settings. For more information, see [Wheel collider suspension: Mass and suspension values](wheel-colliders-suspension.md#mass-suspension-values).
3. On the `Car Body` GameObject, add a [Mesh collider](class-MeshCollider.md) component.
4. On the new **Mesh** **collider** component, enable **Convex**.

## Add the Wheel colliders

To add Wheel colliders to the wheel models, you need to create four new separate GameObjects at the same position as the wheel models (but not as child GameObjects of the wheels).

A fast way to set this up at approximately the same position is to duplicate the wheel GameObjects and then configure the new GameObject:

1. Select all four wheel GameObjects.
2. Duplicate the GameObjects (Ctrl/Cmd+D).
3. Select all four duplicates and do the following:
   1. Remove the **Mesh Renderer** and **Mesh Filter** components
   2. Add a Wheel collider component
4. Rename the new Wheel collider GameObjects. By default, Unity gives duplicate GameObjects the same name and adds a number in parentheses; for example, `Wheel Back Left (1)`. For clarity, add the word “Collider” to the GameObject name instead; for example, `Wheel Back Left collider`.

The Car GameObject hierarchy should now look like this:

![The root Car GameObject and all child GameObjects, including the Car Body, four wheel model GameObjects, and four Wheel collider GameObjects.](../uploads/Main/wheel-collider-tutorial-hierarchy.png)

The root Car GameObject and all child GameObjects, including the Car Body, four wheel model GameObjects, and four Wheel collider GameObjects.

Next, you need to adjust the Wheel colliders’ position and size to match the wheel models.

When you select a Wheel collider, the **Scene view** displays a [gizmo](GizmosMenu.md) which provides a visualization of the Wheel collider settings (see [Wheel collider visualization](wheel-colliders-introduction.md#wheel-collider-visualization)). You can use the gizmo to check the position and size of the Wheel collider against the position and size of the wheel model.

To see the wheel orientation and the gizmo more clearly, set the Scene’s [Draw mode](overlays-reference-draw-modes.md) to **Wireframe** and the [Scene orientation](SceneViewNavigation.md) to **Isometric**.

### Set the Wheel collider position

When you first add the Wheel colliders with the workflow described in [Add Wheel colliders](#add-wheel-colliders), they are too low (in the Scene view, the Wheel collider circle appears below the wheel model mesh). This is because the **Suspension Distance** starts from these GameObjects’ positions, and extends downwards by the distance specified in the **Suspension Distance** setting. The Scene view visualization displays the **Suspension Distance** as an orange line down the Y axis of the Wheel collider’s gizmo.

The green circle outline displays the wheel at the halfway point of the suspension distance extents, and should be considered the wheel’s normal position, when the car is not squashed down or raised up on its suspension. Therefore the green outline of each Wheel collider needs to be centered over its corresponding wheel mesh.

![The Wheel collider gizmo indicates the position of the Wheel collider in relation to the wheel model. In this image, the Wheel collider is too large and too low.](../uploads/Main/wheel-collider-tutorial-gizmo-before.png)

The Wheel collider gizmo indicates the position of the Wheel collider in relation to the wheel model. In this image, the Wheel collider is too large and too low.

To correct this, you need to move the WheelCollider GameObjects up (on the Y axis) by half the value of the Wheel collider’s **Suspension Distance**. In this example project, the **Suspension Distance** is 0.3 (the default value), so you need to move the Wheel collider GameObjects up by 0.15 units.

Unity allows you to enter simple mathematical calculations into numeric fields. You can use this to add to the Y axis value.

1. Select all four Wheel collider GameObjects.
2. In the Inspector, navigate to the Transform’s **Position** property.
3. In the **Y** axis value, place the cursor at the end of the value.
   Add `+0.15` to the end of the value (for example, if the value is 0.5, the value should now read `0.5+0.15`).
4. Press the Return key.

Unity applies +0.15 to the previous value, which moves the Wheel collider GameObjects up the Y axis by 0.15 units.

The Wheel collider gizmo should now be perfectly centered on the wheel meshes:

![The Wheel collider gizmo indicates the position of the Wheel collider in relation to the wheel model. In this image, the Wheel collider is centered correctly in the wheel model.](../uploads/Main/wheel-collider-tutorial-gizmo-correct-transform.png)

The Wheel collider gizmo indicates the position of the Wheel collider in relation to the wheel model. In this image, the Wheel collider is centered correctly in the wheel model.

### Set the Wheel collider radius

When you first add the Wheel colliders with the workflow described in [Add Wheel colliders](#add-wheel-colliders), they are too big (in the Scene view, the Wheel collider gizmo is larger than the wheel model mesh).

To correct this accurately, you need to know the exact radius of the wheel model. This information should be available from your 3D modeling software, or from the technical artist who authored the model.

In this example project, the radius of the wheel model is 0.44.

1. Select all four Wheel collider GameObjects
2. In the Inspector, navigate to the Wheel collider component’s **Radius** property
3. Set the **Radius** to `0.44`.

If the exact radius of a wheel model is unknown or unavailable, you can use the Wheel collider gizmo to approximately match its radius to the model. Alternatively, you can use a [Sphere collider](class-SphereCollider.md) to get the radius, because the Sphere collider automatically resizes to encompass the mesh of its associated model.

To get the radius with a Sphere collider:

1. Select any wheel model GameObject.
2. Add a Sphere collider.
3. Take note of the Sphere collider’s **Radius** property value.
4. Remove the Sphere collider.
5. Select all four Wheel collider GameObjects.
6. In the Inspector, navigate to the Wheel collider component’s **Radius** property.
7. Set the Wheel collider **Radius** to the Sphere collider **Radius** value you noted.

The Wheel colliders should now exactly match the position and size of the wheel models.

![Two Wheel collider gizmos indicate the position of the Wheel colliders in relation to the wheel models. In this image, the Wheel colliders are centered correctly in the wheel model, and their radius correctly matches the radius of the wheel model.](../uploads/Main/wheel-collider-tutorial-gizmo-correct-radius.png)

Two Wheel collider gizmos indicate the position of the Wheel colliders in relation to the wheel models. In this image, the Wheel colliders are centered correctly in the wheel model, and their radius correctly matches the radius of the wheel model.

## Add scripts to control car behavior

To control the car, you need to add **scripts** to the project that do the following:

* **Control car movement**: Send steering, torque, and brake values to the Wheel colliders based on user input.
* **Control wheel movement**: Move and rotate the wheel meshes according to each Wheel collider’s state.

In this example, we use two scripts to do this: `CarControl.cs` and `WheelControl.cs`.

### Add control for car movement

Create a C# file named `CarControl.cs`, and paste in the code below:

```
using UnityEngine;
using UnityEngine.InputSystem;

public class CarControl : MonoBehaviour
{
    [Header("Car Properties")]
    public float motorTorque = 2000f;
    public float brakeTorque = 2000f;
    public float maxSpeed = 20f;
    public float steeringRange = 30f;
    public float steeringRangeAtMaxSpeed = 10f;
    public float centreOfGravityOffset = -1f;

    private WheelControl[] wheels;
    private Rigidbody rigidBody;
    
    [Header("Input Actions")]
    public InputActionReference moveAction;

    private void OnEnable()
    {
        moveAction.action.Enable();
    }

    private void OnDisable()
    {
        moveAction.action.Disable();
    }

    // Start is called before the first frame update
    void Start()
    {
        rigidBody = GetComponent<Rigidbody>();

        // Adjust center of mass to improve stability and prevent rolling
        Vector3 centerOfMass = rigidBody.centerOfMass;
        centerOfMass.y += centreOfGravityOffset;
        rigidBody.centerOfMass = centerOfMass;

        // Get all wheel components attached to the car
        wheels = GetComponentsInChildren<WheelControl>();
    }

    // FixedUpdate is called at a fixed time interval 
    void FixedUpdate()
    {
        // Get player input for acceleration and steering
        float vInput = moveAction.action.ReadValue<Vector2>().y; // Forward/backward input
        float hInput = moveAction.action.ReadValue<Vector2>().x; // Steering input

        // Calculate current speed along the car's forward axis
        float forwardSpeed = Vector3.Dot(transform.forward, rigidBody.linearVelocity);
        float speedFactor = Mathf.InverseLerp(0, maxSpeed, Mathf.Abs(forwardSpeed)); // Normalized speed factor

        // Reduce motor torque and steering at high speeds for better handling
        float currentMotorTorque = Mathf.Lerp(motorTorque, 0, speedFactor);
        float currentSteerRange = Mathf.Lerp(steeringRange, steeringRangeAtMaxSpeed, speedFactor);

        // Determine if the player is accelerating or trying to reverse
        bool isAccelerating = Mathf.Sign(vInput) == Mathf.Sign(forwardSpeed);

        foreach (var wheel in wheels)
        {
            // Apply steering to wheels that support steering
            if (wheel.steerable)
            {
                wheel.WheelCollider.steerAngle = hInput * currentSteerRange;
            }

            if (isAccelerating)
            {
                // Apply torque to motorized wheels
                if (wheel.motorized)
                {
                    wheel.WheelCollider.motorTorque = vInput * currentMotorTorque;
                }
                // Release brakes when accelerating
                wheel.WheelCollider.brakeTorque = 0f;
            }
            else
            {
                // Apply brakes when reversing direction
                wheel.WheelCollider.motorTorque = 0f;
                wheel.WheelCollider.brakeTorque = Mathf.Abs(vInput) * brakeTorque;
            }
        }
    }
}
```

### Add control for car movement with the Input System package

If you are using the [Input System package](com.unity.inputsystem.md), use the following guidance instead. First, you need to create and configure Input Actions in Unity’s Input System package:

1. In Unity, go to **Assets** > **Create** > **Input Actions** and name it `CarInputActions`.
2. Double-click to open it in the Input Actions Editor.
3. Create a new Action Map called `Car` (or any name you prefer).

Inside the `Car` action map, add the following actions:

1. Add a new Action called “Movement”. Set **Action Type** to **Value**, and **Control Type** to **Vector 2**.
2. Right-click on **Movement** and select **Add Up/Down/Left/Right Composite**.
3. Assign a keyboard input to each binding. For a WASD configuration, set **Up** to **W**, **Down** to **S**, **Left** to **A**, and **Right** to **D**.

Finally, save your configuration and generate a C# class from the actions asset:

1. In the Input Actions Editor, select **Save Asset**.
2. Select the `CarInputActions` asset.
3. In the Inspector window, select **Create C# Class**.

Create a C# file named `CarControl.cs`, and paste in the code below:

```
using UnityEngine;

public class CarControl : MonoBehaviour
{
    [Header("Car Properties")]
    public float motorTorque = 2000f;
    public float brakeTorque = 2000f;
    public float maxSpeed = 20f;
    public float steeringRange = 30f;
    public float steeringRangeAtMaxSpeed = 10f;
    public float centreOfGravityOffset = -1f;

    private WheelControl[] wheels;
    private Rigidbody rigidBody;

    private CarInputActions carControls; // Reference to the new input system

    void Awake()
    {
        carControls = new CarInputActions(); // Initialize Input Actions
    }
    void OnEnable()
    {
        carControls.Enable();
    }

    void OnDisable()
    {
        carControls.Disable();
    }
    
    // Start is called before the first frame update
    void Start()
    {
        rigidBody = GetComponent<Rigidbody>();

        // Adjust center of mass to improve stability and prevent rolling
        Vector3 centerOfMass = rigidBody.centerOfMass;
        centerOfMass.y += centreOfGravityOffset;
        rigidBody.centerOfMass = centerOfMass;

        // Get all wheel components attached to the car
        wheels = GetComponentsInChildren<WheelControl>();
    }
            
    // FixedUpdate is called at a fixed time interval
    void FixedUpdate()
    {
        // Read the Vector2 input from the new Input System
        Vector2 inputVector = carControls.Car.Movement.ReadValue<Vector2>();

        // Get player input for acceleration and steering
        float vInput = inputVector.y; // Forward/backward input
        float hInput = inputVector.x; // Steering input
        
        // Calculate current speed along the car's forward axis
        float forwardSpeed = Vector3.Dot(transform.forward, rigidBody.linearVelocity);
        float speedFactor = Mathf.InverseLerp(0, maxSpeed, Mathf.Abs(forwardSpeed)); // Normalized speed factor

        // Reduce motor torque and steering at high speeds for better handling
        float currentMotorTorque = Mathf.Lerp(motorTorque, 0, speedFactor);
        float currentSteerRange = Mathf.Lerp(steeringRange, steeringRangeAtMaxSpeed, speedFactor);

        // Determine if the player is accelerating or trying to reverse
        bool isAccelerating = Mathf.Sign(vInput) == Mathf.Sign(forwardSpeed);

        foreach (var wheel in wheels)
        {
            // Apply steering to wheels that support steering
            if (wheel.steerable)
            {
                wheel.WheelCollider.steerAngle = hInput * currentSteerRange;
            }

            if (isAccelerating)
            {
                // Apply torque to motorized wheels
                if (wheel.motorized)
                {
                    wheel.WheelCollider.motorTorque = vInput * currentMotorTorque;
                }
                // Release brakes when accelerating
                wheel.WheelCollider.brakeTorque = 0f;
            }
            else
            {
                // Apply brakes when reversing direction
                wheel.WheelCollider.motorTorque = 0f;
                wheel.WheelCollider.brakeTorque = Mathf.Abs(vInput) * brakeTorque;
            }
        }
    }
}
```

Add this `CarControl.cs` script to the `Car` root GameObject.

The `CarControl.cs` script handles car behavior such as acceleration, torque, and braking, based on user input. See the code comments for details.

Some elements of the `CarControl.cs` script refer to the `WheelControl.cs` script created in the next section.

### Add control for wheel movement

Create a C# file named `WheelControl.cs`, and paste in the code below:

```
using UnityEngine;

public class WheelControl : MonoBehaviour
{
    public Transform wheelModel;

    [HideInInspector] public WheelCollider WheelCollider;

    // Create properties for the CarControl script
    // (You should enable/disable these via the 
    // Editor Inspector window)
    public bool steerable;
    public bool motorized;

    Vector3 position;
    Quaternion rotation;

    // Start is called before the first frame update
    private void Start()
    {
        WheelCollider = GetComponent<WheelCollider>();
    }

    // Update is called once per frame
    void Update()
    {
        // Get the Wheel collider's world pose values and
        // use them to set the wheel model's position and rotation
        WheelCollider.GetWorldPose(out position, out rotation);
        wheelModel.transform.position = position;
        wheelModel.transform.rotation = rotation;
    }
}
```

Add this script to each Wheel collider GameObject.

The `WheelControl.cs` script uses [`WheelCollider.GetWorldPose`](../ScriptReference/WheelCollider.GetWorldPose.md) to get the Wheel collider’s position in the scene. The script then assigns that position information to a specified wheel model GameObject. See the code comments for details.

Each instance of the `WheelControl.cs` script must have a reference to its corresponding wheel model GameObject.

To assign the correct wheel model GameObject to each Wheel collider:

1. Select a Wheel collider GameObject
2. In the Inspector window, navigate to the **Wheel Control** component created by the script
3. Set the **Wheel Model** to the corresponding wheel model GameObject (for example, on the “Wheel Rear Left collider” GameObject, assign the “Wheel Rear Left” wheel model GameObject):
4. Repeat for each Wheel collider.

You also need to select which wheels receive motor input and steering input from the CarControl script. To simulate a four-wheel drive car via the Wheel Control properties:

1. On all four Wheel collider GameObjects, enable **Motorized**.
2. On the two front Wheel collider GameObjects, enable **Steerable**.

## Test the car

To test the car, enter Play mode and use the arrow or WASD keys to move and steer. Note that input controls only work when the Game view has focus.

To get a better view of the car moving around the Scene:

1. Arrange the Editor windows so that both the Scene view and Game view are visible at the same time.
2. Lock the Scene view’s focus to the car. To do this, select the `Car` root GameObject in the Scene view, then press Shift + F.

Now that you have a basic setup, you can try changing different settings to observe how they affect the movement of the car. You can also follow these instructions with different car models and observe the similarities and differences in their setup.